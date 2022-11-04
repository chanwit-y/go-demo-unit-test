//go:generate mockery --outpkg mock --all --output ./mock
package product

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
			name: "New service",
			args: args{NewMockRepository(t)},
			want: NewService(NewMockRepository(t)),
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
		want    model.Respone[[]*entity.Product]
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "Find success",
			fields: fields{NewMockData(NewMockRepository(t), "Find", MockProducts)},
			want: model.Respone[[]*entity.Product]{
				Data: MockProducts,
			},
		},
		{
			name:    "Find repo error",
			fields:  fields{NewMockError(NewMockRepository(t), "Find")},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := service{
				repository: tt.fields.repository,
			}
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
		data entity.Product
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.Respone[entity.Product]
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "Create success",
			fields: fields{NewMockDataParam(NewMockRepository(t), "Create", MockProduct, MockProduct)},
			args:   args{data: MockProduct},
			want: model.Respone[entity.Product]{
				Data: MockProduct,
			},
		},
		{
			name: "Create error code is not empty",
			args: args{data: MockProductNoCode},
			want: model.Respone[entity.Product]{
				Message: "code is not empty",
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
		data entity.Product
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.Respone[entity.Product]
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := service{
				repository: tt.fields.repository,
			}
			got, err := s.Update(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.Update() = %v, want %v", got, tt.want)
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
		want    model.Respone[entity.Product]
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := service{
				repository: tt.fields.repository,
			}
			got, err := s.Delete(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}
