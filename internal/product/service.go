//go:generate mockery --outpkg mock --all --output ./mock
package product

import (
	"go-demo-unit-test/domain/entity"
	"go-demo-unit-test/domain/model"
)

type (
	Service interface {
		Find() (model.Respone[[]*entity.Product], error)
		Create(data entity.Product) (model.Respone[entity.Product], error)
		Update(data entity.Product) (model.Respone[entity.Product], error)
		Delete(id uint) (model.Respone[entity.Product], error)
	}
	service struct {
		repository Repository
	}
)

func NewService(
	repository Repository) service {
	return service{repository}
}

func (s service) Find() (model.Respone[[]*entity.Product], error) {
	res := model.Respone[[]*entity.Product]{}
	data, err := s.repository.Find()

	if err != nil {
		return res, err
	}

	res.Data = data
	return res, nil
}

func (s service) Create(data entity.Product) (model.Respone[entity.Product], error) {
	res := model.Respone[entity.Product]{}

	if data.Code == "" {
		res.Error = true
		res.Message = "code is not empty"
		return res, nil
	}

	data, err := s.repository.Create(data)
	res.Data = data

	return res, err
}

func (s service) Update(data entity.Product) (model.Respone[entity.Product], error) {
	res := model.Respone[entity.Product]{
		Data:    data,
		Message: "Update product is success",
	}
	err := s.repository.Update(data)

	if err != nil {
		res.Message = err.Error()
		res.Error = true
		return res, nil
	}

	return res, err
}

func (s service) Delete(id uint) (model.Respone[entity.Product], error) {
	res := model.Respone[entity.Product]{
		Message: "Delete product is success",
	}
	err := s.repository.Delete(id)

	if err != nil {
		res.Message = err.Error()
		res.Error = true
		return res, nil
	}

	return res, err

}
