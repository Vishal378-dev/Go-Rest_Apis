package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/vishal/Rest_Apis/internal/configs"
)

type SQLITE struct {
	DBConn *sql.DB
}

func (sqlite *SQLITE) CreateEmployee(empid string, name string, email string, phonenumber int, salary int) (string, error) {
	statement, err := sqlite.DBConn.Prepare("INSERT INTO EMPLOYEE (EMPID,NAME, EMAIL,PHONENUMBER,SALARY) VALUES (?,?,?,?,?)")
	if err != nil {
		return "", err
	}
	defer statement.Close()
	res, err := statement.Exec(empid, name, email, phonenumber, salary)
	fmt.Println("the value is -- > ", res)
	if err != nil {
		return "", err
	}
	LastId, err := res.LastInsertId()
	if err != nil {
		return "", err
	}
	return string(LastId), nil
}

func New(cfg *configs.Config) (*SQLITE, error) {
	db, err := sql.Open("sqlite3", cfg.Storage_path)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS EMPLOYEE (
		EMPID TEXT,
		NAME TEXT, 
		EMAIL TEXT,
		PHONENUMBER INTEGER,
		SALARY INTEGER
	)`)
	if err != nil {
		fmt.Println("error while creating table")
		return nil, err
	}
	return &SQLITE{
		DBConn: db,
	}, nil
}
