package model

type Master struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Status bool   `json:"status"`
}

type Manager struct {
	UserId            string `json:"userId"`
	PersonId          string `json:"personId"`
	EmpId             string `json:"empId"`
	FirstNameEn       string `json:"firstNameEn"`
	LastNameEN        string `json:"lastNameEN"`
	FirstName         string `json:"firstName"`
	LastName          string `json:"lastName"`
	PreferredNameEn   string `json:"preferredNameEn"`
	UserName          string `json:"userName"`
	Email             string `json:"email"`
	HiredDate         string `json:"hiredDate"`
	ResignDate        string `json:"resignDate"`
	PositionCode      string `json:"positionCode"`
	JobTitle          string `json:"jobTitle"`
	ManagerID         string `json:"managerID"`
	CompanyCode       string `json:"companyCode"`
	BusinessUnitCode  string `json:"businessUnitCode"`
	FunctionCode      string `json:"functionCode"`
	SubFunction1Code  string `json:"subFunction1Code"`
	SubFunction2Code  string `json:"subFunction2Code"`
	SubFunction3Code  string `json:"subFunction3Code"`
	SubFunction4Code  string `json:"subFunction4Code"`
	SubFunction5Code  string `json:"subFunction5Code"`
	SubFunction6Code  string `json:"subFunction6Code"`
	SubFunction7Code  string `json:"subFunction7Code"`
	WorkingLocation   string `json:"workingLocation"`
	Country           string `json:"country"`
	OfficePhone       string `json:"officePhone"`
	IpPhone           string `json:"ipPhone"`
	Status            bool   `json:"status"`
	CostCenter        string `json:"costCenter"`
	JobLevel          string `json:"jobLevel"`
	MultipleJob       string `json:"multipleJob"`
	JobLevelDoa       string `json:"jobLevelDoa"`
	JobLevelDoaNew    string `json:"jobLevelDoaNew"`
	UserType          string `json:"userType"`
	JobStatus         string `json:"jobStatus"`
	UserPrincipalName string `json:"userPrincipalName"`
}

type UserProfile struct {
	Gender            string  `json:"gender"`
	DateOfBirth       string  `json:"dateOfBirth"`
	NationalID        string  `json:"nationalID"`
	UserId            string  `json:"userId"`
	PersonId          string  `json:"personId"`
	EmpId             string  `json:"empId"`
	FirstNameEn       string  `json:"firstNameEn"`
	LastNameEN        string  `json:"lastNameEN"`
	FirstName         string  `json:"firstName"`
	LastName          string  `json:"lastName"`
	PreferredNameEn   string  `json:"preferredNameEn"`
	UserName          string  `json:"userName"`
	Email             string  `json:"email"`
	HiredDate         string  `json:"hiredDate"`
	ResignDate        string  `json:"resignDate"`
	PositionCode      string  `json:"positionCode"`
	JobTitle          string  `json:"jobTitle"`
	ManagerID         string  `json:"managerID"`
	CompanyCode       string  `json:"companyCode"`
	BusinessUnitCode  string  `json:"businessUnitCode"`
	FunctionCode      string  `json:"functionCode"`
	SubFunction1Code  string  `json:"subFunction1Code"`
	SubFunction2Code  string  `json:"subFunction2Code"`
	SubFunction3Code  string  `json:"subFunction3Code"`
	SubFunction4Code  string  `json:"subFunction4Code"`
	SubFunction5Code  string  `json:"subFunction5Code"`
	SubFunction6Code  string  `json:"subFunction6Code"`
	SubFunction7Code  string  `json:"subFunction7Code"`
	WorkingLocation   string  `json:"workingLocation"`
	Country           string  `json:"country"`
	OfficePhone       string  `json:"officePhone"`
	IpPhone           string  `json:"ipPhone"`
	Status            bool    `json:"status"`
	CostCenter        string  `json:"costCenter"`
	JobLevel          string  `json:"jobLevel"`
	MultipleJob       string  `json:"multipleJob"`
	JobLevelDoa       string  `json:"jobLevelDoa"`
	JobLevelDoaNew    string  `json:"jobLevelDoaNew"`
	UserType          string  `json:"userType"`
	JobStatus         string  `json:"jobStatus"`
	UserPrincipalName string  `json:"userPrincipalName"`
	JobLevelModel     Master  `json:"jobLevelModel"`
	CompanyModel      Master  `json:"companyModel"`
	BusinessUnitModel Master  `json:"businessUnitModel"`
	FunctionModel     Master  `json:"functionModel"`
	SubFunction1Model Master  `json:"subFunction1Model"`
	SubFunction2Model Master  `json:"subFunction2Model"`
	LanguageCode      string  `json:"LanguageCode"`
	Manager           Manager `json:"manager"`
}
