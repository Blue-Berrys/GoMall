package kafka

import (
	"context"
	"fmt"
	"github.com/Blue-Berrys/GoMall/app/checkout/conf"
	"github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/order"
	"github.com/segmentio/kafka-go"
	"sync"
	"time"
)

var (
	writer         *kafka.Writer
	readerOrder    *kafka.Reader
	readerPayment  *kafka.Reader
	topic          string = "order_topic"
	ctx            context.Context
	groupOrderId   string = "order_group"
	groupPaymentId string = "payment_group"
	once           sync.Once
	initErr        error
)

func InitWrite() error {
	writer = &kafka.Writer{
		Addr:                   kafka.TCP(conf.GetConf().Kafka.Address),
		Topic:                  topic,
		Balancer:               &kafka.Hash{},
		WriteTimeout:           1 * time.Second,
		RequiredAcks:           kafka.RequireAll,
		AllowAutoTopicCreation: true,
	}
	testMsg := kafka.Message{
		Key:   []byte("test_key"),
		Value: []byte("test_value"),
	}
	err := writer.WriteMessages(ctx, testMsg)
	if err != nil {
		return fmt.Errorf("write messages: %w", err)
	}
	return nil
}
func InitReadOrder() error {
	readerOrder = kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{conf.GetConf().Kafka.Address},
		Topic:          topic,
		GroupID:        groupOrderId,
		StartOffset:    kafka.FirstOffset,      // 从最早的消息开始消费
		CommitInterval: 500 * time.Millisecond, // 提交偏移量的间隔
	})
	_, err := readerOrder.ReadMessage(ctx)
	if err != nil {
		return fmt.Errorf("read message: %w", err)
	}
	return nil
}
func InitReadPayment() error {
	readerPayment = kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{conf.GetConf().Kafka.Address},
		Topic:          topic,
		GroupID:        groupPaymentId,
		StartOffset:    kafka.FirstOffset,      // 从最早的消息开始消费
		CommitInterval: 500 * time.Millisecond, // 提交偏移量的间隔
	})
	_, err := readerPayment.ReadMessage(ctx)
	if err != nil {
		return fmt.Errorf("read message: %w", err)
	}
	return nil
}

type Address struct {
	StreetAddress string `protobuf:"bytes,1,opt,name=street_address,json=streetAddress,proto3" json:"street_address,omitempty"`
	City          string `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	State         string `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
	Country       string `protobuf:"bytes,4,opt,name=country,proto3" json:"country,omitempty"`
	ZipCode       string `protobuf:"bytes,5,opt,name=zip_code,json=zipCode,proto3" json:"zip_code,omitempty"`
}

type CreditCardInfo struct {
	CreditCardNumber          string `protobuf:"bytes,1,opt,name=credit_card_number,json=creditCardNumber,proto3" json:"credit_card_number,omitempty"`
	CreditCardCvv             string `protobuf:"bytes,2,opt,name=credit_card_cvv,json=creditCardCvv,proto3" json:"credit_card_cvv,omitempty"`
	CreditCardExpirationYear  int32  `protobuf:"varint,3,opt,name=credit_card_expiration_year,json=creditCardExpirationYear,proto3" json:"credit_card_expiration_year,omitempty"`
	CreditCardExpirationMonth int32  `protobuf:"varint,4,opt,name=credit_card_expiration_month,json=creditCardExpirationMonth,proto3" json:"credit_card_expiration_month,omitempty"`
}

type OrderMessage struct {
	OrderId    string             `json:"order_id"`
	UserId     uint32             `json:"user_id"`
	Email      string             `json:"email"`
	Address    Address            `json:"address"`
	Items      []*order.OrderItem `json:"items"`
	CreditCard *CreditCardInfo    `json:"credit_card,omitempty"`
	TotalPrice float32            `json:"total_price"`
}

func Init() {
	once.Do(func() {
		ctx = context.Background()
		if err := InitWrite(); err != nil {
			initErr = fmt.Errorf("failed to initialize Kafka writer: %v", err)
		}
		if err := InitReadOrder(); err != nil {
			initErr = fmt.Errorf("failed to initialize Kafka reader: %v", err)
		}
		if err := InitReadPayment(); err != nil {
			initErr = fmt.Errorf("failed to initialize Kafka reader: %v", err)
		}
	})
	fmt.Println(initErr)
}
