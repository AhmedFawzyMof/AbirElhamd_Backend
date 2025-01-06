package handlers

import (
	"abir-el-hamd/internal/config"
	"abir-el-hamd/internal/middleware"
	"abir-el-hamd/internal/models"
	"encoding/json"
	"net/http"
)

func Logs(w http.ResponseWriter, r *http.Request) {
	database := config.Database()
	defer database.Close()

	log := models.Logs{}

	logs, err := log.GetAll(database)

	if err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}

	Response := map[string]interface{}{
		"logs": logs,
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(Response); err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}
}
