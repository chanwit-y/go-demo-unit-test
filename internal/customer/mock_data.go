package customer

import (
	"go-demo-unit-test/domain/entity"
	"go-demo-unit-test/domain/model"

	"gorm.io/gorm"
)

var (
	MockCustomer = entity.Customer{
		Model: gorm.Model{
			ID: 1,
		},
		Code:  "0001",
		Name:  "cus-1",
		Tel:   "xxxxxxx",
		Email: "test@mail.com",
	}
	//Begin Test Find
	MockCustomers = []*entity.Customer{
		&MockCustomer,
		// {
		// 	Code:  "0001",
		// 	Name:  "cus-1",
		// 	Tel:   "xxxxxxx",
		// 	Email: "test@mail.com",
		// },
	}

	MockWantCusomerFind = model.Respone[[]*entity.Customer]{
		Data:    MockCustomers,
		Message: "",
		Error:   false,
	}
	//End Test Find

	MockCustomersNoData   = []*entity.Customer{}
	MockWantCusomerNoData = model.Respone[[]*entity.Customer]{
		Data:    MockCustomersNoData,
		Message: "",
		Error:   false,
	}

	MockCustomerParamNoCode = entity.Customer{
		Code:  "",
		Name:  "cus-1",
		Tel:   "xxxxxxx",
		Email: "test@mail.com",
	}
)
