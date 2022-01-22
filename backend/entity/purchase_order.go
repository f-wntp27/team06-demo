package entity

import (
	"time"

	"gorm.io/gorm"
)

type PaymentMethod struct {
	gorm.Model
	MethodName string

	PurchaseOrders []PurchaseOrder `gorm:"foreignKey:PaymentMethodID"`
}

type PurchaseOrder struct {
	gorm.Model

	PaymentMethodID *uint
	PaymentMethod   PaymentMethod `gorm:"references:ID"`

	OrderTime       time.Time
	DeliveryAddress string
}
