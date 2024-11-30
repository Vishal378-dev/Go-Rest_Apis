package employee

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/vishal/Rest_Apis/internal/http/types"
	"github.com/vishal/Rest_Apis/internal/utils"
)

func NewEmployee() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			var employee types.Employee
			err := json.NewDecoder(r.Body).Decode(&employee)
			if errors.Is(err, io.EOF) {
				utils.WriteResponse(w, http.StatusBadRequest, err.Error())
				return
			}
			slog.Info("Creating a New Employee")
			utils.WriteResponse(w, http.StatusOK, employee)
		} else {
			w.Write([]byte("wrong method"))
		}

	}
}
