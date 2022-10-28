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
	mock := NewMockRepository(t)
	tests := []struct {
		name    string
		fields  fields
		want    model.Respone[[]*entity.Customer]
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "Test Find have data",
			fields: fields{repository: NewMock(mock, "Find", MockCustomers)},
			want:   MockWantCusomerFind,
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
		want    entity.Customer
		wantErr bool
	}{
		// TODO: Add test cases.
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
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := service{
				repository: tt.fields.repository,
			}
			if err := s.Update(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("service.Update() error = %v, wantErr %v", err, tt.wantErr)
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
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := service{
				repository: tt.fields.repository,
			}
			if err := s.Delete(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("service.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
