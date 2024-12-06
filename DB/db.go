package db

type DB interface {
	CreateEmployee(empid string, name string, email string, phonenumber int, salary int) (string, error)
}
