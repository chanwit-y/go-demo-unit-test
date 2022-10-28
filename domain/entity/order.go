package entity

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	OrderNo      string
	OrderDate    time.Time
	CustomerID   uint
	OrderDetails []*OrderDetail `gorm:"many2many:order_details;"`
}
