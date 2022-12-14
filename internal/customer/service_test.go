//go:generate mockery --outpkg mock --all --output ./mock
package customer

import (
	"go-demo-unit-test/domain/entity"
	"go-demo-unit-test/domain/model"
	"reflect"
	"testing"
)

func TestNewService(t *testing.T) {
	type args struct {
		repository Repository
	}
	tests := []struct {
		name string
		args args
		want service
	}{
		// TODO: Add test cases.
		{
			name: "NewService",
			args: args{NewMockRepository(t)},
			want: service{NewMockRepository(t)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewService(tt.args.repository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_Find(t *testing.T) {
	type fields struct {
		repository Repository
	}
	tests := []struct {
		name    string
		fields  fields
		want    model.Respone[[]*entity.Customer]
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "Test Find",
			fields: fields{repository: NewMockData(NewMockRepository(t), "Find", MockCustomers)},
			want:   MockWantCusomerFind,
		},
		{
			name:   "Test Find not found",
			fields: fields{repository: NewMockData(NewMockRepository(t), "Find", MockCustomersNoData)},
			want:   MockWantCusomerNoData,
		},
		{
			name: "Test Find repository error",
			fields: fields{
				repository: NewMockError(
					NewMockRepository(t),
					"Find")},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := service{tt.fields.repository}
			got, err := s.Find()
			if (err != nil) != tt.wantErr {
				t.Errorf("service.Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.Find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_Create(t *testing.T) {
	type fields struct {
		repository Repository
	}
	type args struct {
		data entity.Customer
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.Respone[entity.Customer]
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Create customer",
			args: args{data: *MockCustomers[0]},
			fields: fields{
				repository: NewMockDataParam(
					NewMockRepository(t),
					"Create",
					*MockCustomers[0],
					*MockCustomers[0])},
			want: model.Respone[entity.Customer]{
				Data: *MockCustomers[0],
			},
		},
		{
			name: "Create customer error param code",
			args: args{data: MockCustomerParamNoCode},
			want: model.Respone[entity.Customer]{
				Data:    MockCustomerParamNoCode,
				Message: "customer code is not empty",
				Error:   true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := service{
				repository: tt.fields.repository,
			}
			got, err := s.Create(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_Update(t *testing.T) {
	type fields struct {
		repository Repository
	}
	type args struct {
		data entity.Customer
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		want    model.Respone[entity.Customer]
	}{
		// TODO: Add test cases.
		{
			name:   "Update customer data success",
			args:   args{MockCustomer},
			fields: fields{repository: NewMockParamNoData(NewMockRepository(t), "Update", MockCustomer)},
			want: model.Respone[entity.Customer]{
				Data:    MockCustomer,
				Message: "Update customer is success",
				Error:   false,
			},
		},
		{
			name: "Update db error",
			args: args{MockCustomer},
			fields: fields{repository: NewMockParamError(
				NewMockRepository(t),
				"Update",
				MockCustomer,
				"error")},
			want: model.Respone[entity.Customer]{
				Data:    MockCustomer,
				Message: "error",
				Error:   true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := service{
				repository: tt.fields.repository,
			}
			got, err := s.Update(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_Delete(t *testing.T) {
	type fields struct {
		repository Repository
	}
	type args struct {
		id uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		want    model.Respone[entity.Customer]
	}{
		// TODO: Add test cases.
		{
			name:   "Delete customer",
			args:   args{id: uint(1)},
			fields: fields{NewMockParamNoData(NewMockRepository(t), "Delete", uint(1))},
			want: model.Respone[entity.Customer]{
				Message: "Delete customer is success",
			},
		},
		{
			name:   "Delete error",
			args:   args{id: uint(1)},
			fields: fields{NewMockParamError(NewMockRepository(t), "Delete", uint(1), "error")},
			want: model.Respone[entity.Customer]{
				Message: "error",
				Error:   true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := service{
				repository: tt.fields.repository,
			}
			got, err := s.Delete(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}
