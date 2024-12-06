package employee

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	db "github.com/vishal/Rest_Apis/DB"
	"github.com/vishal/Rest_Apis/internal/http/types"
	"github.com/vishal/Rest_Apis/internal/utils"
)

func NewEmployee(db db.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			var employee types.Employee
			err := json.NewDecoder(r.Body).Decode(&employee)
			if errors.Is(err, io.EOF) {
				// utils.WriteResponse(w, http.StatusBadRequest, utils.CommonError(err))
				utils.WriteResponse(w, http.StatusBadRequest, utils.CommonError(fmt.Errorf("empty body")))
				return
			}
			if err != nil {
				utils.WriteResponse(w, http.StatusBadGateway, utils.CommonError(err))
				return
			}
			slog.Info("Creating a New Employee")

			// request validation
			err = validator.New().Struct(employee)

			if err != nil {
				validateErrs := err.(validator.ValidationErrors)
				utils.WriteResponse(w, http.StatusBadRequest, utils.ValidationError(validateErrs))
				return
			}

			// storing student in db
			lastId, err := db.CreateEmployee(employee.EmpID, employee.Name, employee.Email, employee.PhoneNumber, int(employee.Salary))
			if err != nil {
				utils.CommonError(err)
			}

			utils.WriteResponse(w, http.StatusOK, lastId)
		} else {
			w.Write([]byte("wrong method"))
		}

	}
}
