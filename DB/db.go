package db

type DB interface {
	CreateEmployee(empid string, name string, email string, phonenumber int, salary int) (string, error)
}

// getBYID where r.pathvalue("id") --> {id } -->queryrow.scan(&parameter you want ot grab)  sql.errnorows
