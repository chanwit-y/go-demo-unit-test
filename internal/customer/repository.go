package customer

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
		Find() ([]*entity.Customer, error)
		Create(data entity.Customer) (entity.Customer, error)
		Update(data entity.Customer) error
		Delete(id uint) error
	}
	repository struct{}
)

func NewRepository(db *gorm.DB) repository {
	query.SetDefault(db)
	return repository{}
}

func (r repository) Find() ([]*entity.Customer, error) {
	customer := query.Customer
	return customer.WithContext(ctx).Find()
}

func (r repository) Create(data entity.Customer) (entity.Customer, error) {
	customer := query.Customer
	err := customer.WithContext(ctx).Create(&data)

	return data, err
}

func (r repository) Update(data entity.Customer) error {
	customer := query.Customer
	_, err := customer.WithContext(ctx).Where(customer.ID.Eq(data.ID)).Updates(structs.Map(data))

	return err
}

func (r repository) Delete(id uint) error {
	customer := query.Customer
	_, err := customer.WithContext(ctx).Where(customer.ID.Eq(id)).Delete()
	return err
}
