package entity

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Code   string
	Name   string
	Tel    string
	Email  string
	Orders []Order
}
