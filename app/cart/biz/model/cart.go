package model

import (
	"context"
	"errors"
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserId    uint32 `gorm:"type int(11); not null; index:idx_user_id"` // 索引可以帮助优化查询
	ProductId uint32 `gorm:"type int(11); not null"`
	Qty       uint32 `gorm:"type int(11); not null"`
}

func TableName() string {
	return "cart"
}

func AddItem(ctx context.Context, db *gorm.DB, c *Cart) error {
	// 为了避免和包名cart冲突，用c作为变量
	var row Cart
	// 数据库找一下有没有报错
	err := db.WithContext(ctx).
		Model(&Cart{}).
		Where(&Cart{UserId: c.UserId, ProductId: c.ProductId}).
		First(&row).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) { // 不是找不到类型的错误
		return err // 如果没创建返回的就是找不到这条记录
	}
	if row.ID > 0 { // 找到这行了，给它加数量
		return db.WithContext(ctx).
			Model(&Cart{}).
			Where(&Cart{UserId: c.UserId}).
			UpdateColumn("qty", gorm.Expr("qty + ?", c.Qty)).Error
	}
	// 不存在这行
	return db.Where(ctx).Create(c).Error
}

func EmptyCart(ctx context.Context, db *gorm.DB, userId uint32) error {
	if userId == 0 {
		return errors.New("user id is required")
	}
	return db.WithContext(ctx).Delete(&Cart{}, "user_id = ?", userId).Error
}

func GetCartByUserId(ctx context.Context, db *gorm.DB, userId uint32) ([]*Cart, error) {
	var rows []*Cart
	err := db.WithContext(ctx).Model(&Cart{}).Where(&Cart{UserId: userId}).Find(&rows).Error
	return rows, err
}
