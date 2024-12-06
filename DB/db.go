package db

type DB interface {
	CreateEmployee(name string, email string, phonenumber int) (string, error)
}
