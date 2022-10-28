package entity

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Code   string
	Name   string
	Price  float64
	Orders []*Order `gorm:"many2many:order_details;"`
}
