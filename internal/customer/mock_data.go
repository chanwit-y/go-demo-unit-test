package customer

import (
	"go-demo-unit-test/domain/entity"
	"go-demo-unit-test/domain/model"
)

var (
	MockCustomers = []*entity.Customer{
		{
			Code:  "0001",
			Name:  "cus-1",
			Tel:   "xxxxxxx",
			Email: "test@mail.com",
		},
	}

	MockWantCusomerFind = model.Respone[[]*entity.Customer]{
		Data:    MockCustomers,
		Message: "",
		Errors:  nil,
	}
)
