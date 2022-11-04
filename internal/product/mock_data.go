package product

import (
	"go-demo-unit-test/domain/entity"

	"gorm.io/gorm"
)

var (
	MockProductNoCode = entity.Product{
		Model: gorm.Model{ID: 1},
		Code:  "",
		Name:  "Product 1",
		Price: 1000,
	}

	MockProduct = entity.Product{
		Model: gorm.Model{ID: 1},
		Code:  "0001",
		Name:  "Product 1",
		Price: 1000,
	}

	MockProducts = []*entity.Product{&MockProduct}
)
