package model

type Company struct {
	UserList    []User `json:"userList"`
	CompanyName string `json:"companyName"`
}
