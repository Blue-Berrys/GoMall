package model

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"time"
)

type Product struct {
	Base
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Picture     string  `json:"picture"`
	Price       float32 `json:"price"`

	Categories []Category `json:"categories" gorm:"many2many:product_category"` // 商品和分类的多对多关系,连接product和category
}

func (p Product) TableName() string {
	return "product"
}

// 封装查询逻辑，便于复用和扩展
type ProductQuery struct {
	ctx context.Context
	db  *gorm.DB
}

func (p ProductQuery) GetById(productId int) (product Product, err error) {
	err = p.db.WithContext(p.ctx).Model(&Product{}).Where("id = ?", productId).First(&product).Error
	// db.WithContext 将一个 context.Context 与数据库操作绑定，允许在数据库操作期间传递上下文信息。
	// 支持超时和取消, 适用于需要严格控制响应时间的场景（如高并发系统）
	// 便于跟踪请求链路

	return product, err
}

func (p ProductQuery) SearchProducts(q string) (product []*Product, err error) {
	err = p.db.WithContext(p.ctx).Model(&Product{}).Find(&product, "name like ? or description like ?",
		"%"+q+"%", "%"+q+"%",
	).Error
	// 会像这样 SELECT * FROM product.templ WHERE name LIKE '%phone%' OR description LIKE '%phone%';
	return product, err
}

func NewProductQuery(ctx context.Context, db *gorm.DB) *ProductQuery {
	return &ProductQuery{
		ctx: ctx,
		db:  db,
	}
}

type CachedProductQuery struct {
	productQuery ProductQuery
	cacheClient  *redis.Client
	prefix       string
}

func (c CachedProductQuery) GetById(productId int) (product Product, err error) {
	lockKey := fmt.Sprintf("%s_%s_%d_lock", c.prefix, "product_by_id", productId)
	cachedKey := fmt.Sprintf("%s_%s_%d", c.prefix, "product_by_id", productId)
	cachedResult := c.cacheClient.Get(c.productQuery.ctx, cachedKey)
	err = func() error { // 如果redis获取到了
		if err := cachedResult.Err(); err != nil {
			return err
		}
		cachedResultByte, err := cachedResult.Bytes() // 缓存结果转换为字节数组
		if err != nil {
			return err
		}
		err = json.Unmarshal(cachedResultByte, &product) // 将字节数组反序列化为 product 对象
		if err != nil {
			return err
		}
		return nil
	}()
	if err == nil {
		return product, nil
	}
	// 2. 防止缓存击穿，不在缓存中的数据突然高流量访问，分布式锁保证同一时间只有一个请求访问数据库
	maxRetries := 3
	retryDelay := 100 * time.Millisecond
	for i := 0; i < maxRetries; i++ {
		if c.cacheClient.SetNX(c.productQuery.ctx, lockKey, "1", 5*time.Second).Val() {
			// SetNX 命令尝试获取分布式锁。如果获取锁失败，说明有其他请求正在从数据库中读取数据，当前请求等待一段时间后重试。
			// 如果获取成功了，别的就不能用这个了
			defer c.cacheClient.Del(c.productQuery.ctx, lockKey) // defer 确保在函数返回时释放锁。
			break
		}
		time.Sleep(retryDelay) // 如果获取锁失败，等待一段时间后重试
	}

	product, err = c.productQuery.GetById(productId) // 从数据库获取
	if err != nil {                                  // 数据库中没有的时候
		// 1.防止缓存穿透，缓存和存储层都不存在时，redis缓存一个空对象设置一个较短的过期时间防止占用过多的内存空间，或者用一个布隆过滤器
		product := Product{}
		encoded, err := json.Marshal(&product)
		if err != nil {
			return product, err
		}
		_ = c.cacheClient.Set(c.productQuery.ctx, cachedKey, encoded, 5*time.Minute).Err()
		return product, nil
	}
	// 从数据库获取成功，放到缓存里
	encoded, err := json.Marshal(product) // 结果序列化为 JSON 格式
	if err != nil {
		return product, nil
	}
	_ = c.cacheClient.Set(c.productQuery.ctx, cachedKey, encoded, time.Hour)
	return product, nil
}

func (c CachedProductQuery) SearchProducts(q string) (products []*Product, err error) {
	return c.productQuery.SearchProducts(q) // 假设缓存命中率很低，直接用mysql找
}

func NewCachedProductQuery(ctx context.Context, db *gorm.DB, cacheClient *redis.Client) *CachedProductQuery {
	return &CachedProductQuery{
		productQuery: *NewProductQuery(ctx, db),
		cacheClient:  cacheClient,
		prefix:       "shop",
	}
}

// 给ProductQuery传读的db，给ProductMutation传写的db，实现读写分离
type ProductMutation struct {
	ctx context.Context
	db  *gorm.DB
}
