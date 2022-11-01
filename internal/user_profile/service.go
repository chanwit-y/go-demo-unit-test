package user_profile

import (
	"encoding/json"
	"fmt"
	"go-demo-unit-test/domain/model"
	"go-demo-unit-test/pkg/env"
	"go-demo-unit-test/pkg/util"
)

type (
	Service interface {
		GetByEmail(email string) (model.Respone[model.UserProfile], error)
	}
	service struct {
		http util.HttpClient
	}
)

func NewService(http util.HttpClient) Service {
	return service{http}
}

func (s service) GetByEmail(email string) (model.Respone[model.UserProfile], error) {
	var userProfile model.Respone[model.UserProfile]
	res, err := s.http.Get(fmt.Sprintf("%v/employees/search/email-kpi?email=%v", env.Env().UserProfileURL, email))
	if err != nil {
		return userProfile, err
	}
	json.Unmarshal(res, &userProfile)
	return userProfile, nil
}
