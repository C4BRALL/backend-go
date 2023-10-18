package enums

type Status int

const (
	Active   Status = 1
	Inactive Status = 0
)

type UserType string

const (
	Customer    UserType = "customers"
	Salesperson UserType = "seller"
)
