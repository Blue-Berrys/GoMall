package model

import (
	"context"
	"gorm.io/gorm"
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
