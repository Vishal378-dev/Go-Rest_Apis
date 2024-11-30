package types

type Employee struct {
	EmpID       string  `json:"empid"`
	Name        string  `json:"name" validate:"required"`
	Email       string  `json:"email" validate:"required"`
	PhoneNumber int     `json:"phonenumber" validate:"required"`
	Salary      float64 `json:"salary"`
}
