package model

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	ProductId    uint32  `gorm:"type:int(11)"`
	OrderIdRefer string  `gorm:"type:varchar(100);index"` //与订单主表关联
	Quantity     uint32  `gorm:"int(11)"`
	Cost         float32 `gorm:"type:decimal(10,2)"`
}

func (OrderItem) TableName() string {
	return "order_item"
}
