package model

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type PaymentLog struct {
	gorm.Model
	UserId        uint32    `json:"user_id"`
	OrderId       string    `json:"order_id"`
	TransactionId string    `json:"transaction_id"`
	Amount        float32   `json:"amount"`
	PayAt         time.Time `json:"pay_at"`
}

func (p *PaymentLog) TableName() string {
	return "payment_log"
}

func CreatePaymentLog(ctx context.Context, db *gorm.DB, payment *PaymentLog) error {
	return db.WithContext(ctx).Model(&PaymentLog{}).Create(payment).Error
}
