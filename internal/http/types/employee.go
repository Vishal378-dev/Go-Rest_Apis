package types

type Employee struct {
	EmpID       string  `json:"empid"`
	Name        string  `json:"name"`
	Email       string  `json:"email"`
	PhoneNumber int     `json:"phonenumber"`
	Salary      float64 `json:"salary"`
}
