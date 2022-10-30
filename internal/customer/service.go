//go:generate mockery --outpkg mock --all --output ./mock
package customer

import (
	"go-demo-unit-test/domain/entity"
	"go-demo-unit-test/domain/model"
)

type (
	Service interface {
		Find() (model.Respone[[]*entity.Customer], error)
		Create(data entity.Customer) (model.Respone[entity.Customer], error)
		Update(data entity.Customer) (model.Respone[entity.Customer], error)
		Delete(id uint) (model.Respone[entity.Customer], error)
	}
	service struct {
		repository Repository
	}
)

func NewService(
	repository Repository) service {
	return service{repository}
}

func (s service) Find() (model.Respone[[]*entity.Customer], error) {
	res := model.Respone[[]*entity.Customer]{}
	data, err := s.repository.Find()

	if err != nil {
		return res, err
	}

	res.Data = data
	return res, nil
}

func (s service) Create(data entity.Customer) (model.Respone[entity.Customer], error) {
	res := model.Respone[entity.Customer]{}
	if data.Code == "" {
		msg := "customer code is not empty"
		res.Data = data
		res.Message = msg
		res.Error = true
		return res, nil
	}
	data, err := s.repository.Create(data)
	res.Data = data

	return res, err
}

func (s service) Update(data entity.Customer) (model.Respone[entity.Customer], error) {
	res := model.Respone[entity.Customer]{
		Data:    data,
		Message: "Update customer is success",
	}
	err := s.repository.Update(data)

	if err != nil {
		res.Message = err.Error()
		res.Error = true
		return res, nil
	}

	return res, err
}

func (s service) Delete(id uint) (model.Respone[entity.Customer], error) {
	res := model.Respone[entity.Customer]{
		Message: "Delete customer is success",
	}
	err := s.repository.Delete(id)

	if err != nil {
		res.Message = err.Error()
		res.Error = true
		return res, nil
	}

	return res, err

}
