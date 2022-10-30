package customer

import (
	"go-demo-unit-test/domain/entity"
	"go-demo-unit-test/domain/model"
)

var (
	//Begin Test Find
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
	//End Test Find

	MockCustomersNoData   = []*entity.Customer{}
	MockWantCusomerNoData = model.Respone[[]*entity.Customer]{
		Data:    MockCustomersNoData,
		Message: "",
		Errors:  nil,
	}

	MockCustomerParamNoCode = entity.Customer{
		Code:  "",
		Name:  "cus-1",
		Tel:   "xxxxxxx",
		Email: "test@mail.com",
	}
)
