package product

import (
	"go-demo-unit-test/domain/entity"
	"go-demo-unit-test/domain/model"

	"gorm.io/gorm"
)

var (
	MockProductNoCode = entity.Product{
		Model: gorm.Model{
			ID: 1,
		},
		Code:  "",
		Name:  "Product 01",
		Price: 100.0,
	}

	MockProduct = entity.Product{
		Model: gorm.Model{
			ID: 1,
		},
		Code:  "0001",
		Name:  "Product 01",
		Price: 100.0,
	}

	MockProducts = []*entity.Product{
		&MockProduct,
	}

	MockWantProductFind = model.Respone[[]*entity.Product]{
		Data: MockProducts,
	}
)
