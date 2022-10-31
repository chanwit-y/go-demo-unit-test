package product

import (
	"context"
	"go-demo-unit-test/domain/entity"
	"go-demo-unit-test/query"

	"github.com/fatih/structs"
	"gorm.io/gorm"
)

var ctx = context.Background()

type (
	Repository interface {
		Find() ([]*entity.Product, error)
		Create(data entity.Product) (entity.Product, error)
		Update(data entity.Product) error
		Delete(id uint) error
	}
	repository struct{}
)

func NewRepository(db *gorm.DB) repository {
	query.SetDefault(db)
	return repository{}
}

func (r repository) Find() ([]*entity.Product, error) {
	product := query.Product
	return product.WithContext(ctx).Find()
}

func (r repository) Create(data entity.Product) (entity.Product, error) {
	product := query.Product
	err := product.WithContext(ctx).Create(&data)

	return data, err
}

func (r repository) Update(data entity.Product) error {
	product := query.Product
	_, err := product.WithContext(ctx).Where(product.ID.Eq(data.ID)).Updates(structs.Map(data))

	return err
}

func (r repository) Delete(id uint) error {
	product := query.Product
	_, err := product.WithContext(ctx).Where(product.ID.Eq(id)).Delete()
	return err
}
