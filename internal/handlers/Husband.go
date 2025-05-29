package handlers

import (
	"abir-el-hamd/internal/config"
	"abir-el-hamd/internal/middleware"
	"abir-el-hamd/internal/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func AddHusband(w http.ResponseWriter, r *http.Request) {
	database := config.Database()
	defer database.Close()

	data := make(map[string]interface{})

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		fmt.Println(err)
		middleware.ErrorResopnse(w, err)
		return
	}

	id := data["case_id"].(string)
	age, err := strconv.Atoi(data["age"].(string))

	if err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}

	husband := models.Husband{
		Name:          sql.NullString{String: data["name"].(string), Valid: true},
		National_id:   sql.NullString{String: data["national_id"].(string), Valid: true},
		Date_of_birth: sql.NullString{String: data["date_of_birth"].(string), Valid: true},
		Age:           sql.NullInt32{Int32: int32(age), Valid: true},
		Gender:        sql.NullString{String: data["gender"].(string), Valid: true},
		Case_id:       sql.NullString{String: id, Valid: true},
	}

	if err := husband.Add(database); err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}

	Response := map[string]interface{}{
		"success": "تمت العملية بنجاح",
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(Response); err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}

}

func UpdateHusband(w http.ResponseWriter, r *http.Request) {
	database := config.Database()
	defer database.Close()

	data := make(map[string]interface{})

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		fmt.Println(err)
		middleware.ErrorResopnse(w, err)
		return
	}

	age, err := strconv.Atoi(data["age"].(string))

	if err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}

	case_id := data["case_id"].(string)

	id := int(data["id"].(float64))

	husband := models.Husband{
		Id:            sql.NullInt32{Int32: int32(id), Valid: true},
		Name:          sql.NullString{String: data["name"].(string), Valid: true},
		National_id:   sql.NullString{String: data["national_id"].(string), Valid: true},
		Date_of_birth: sql.NullString{String: data["date_of_birth"].(string), Valid: true},
		Age:           sql.NullInt32{Int32: int32(age), Valid: true},
		Gender:        sql.NullString{String: data["gender"].(string), Valid: true},
		Case_id:       sql.NullString{String: case_id, Valid: true},
	}

	if err := husband.UPDATE(database); err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}

	Response := map[string]interface{}{
		"success": "تمت العملية بنجاح",
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(Response); err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}
}

func DeleteHusband(w http.ResponseWriter, r *http.Request) {
	database := config.Database()
	defer database.Close()

	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}

	husband := models.Husband{
		Id: sql.NullInt32{
			Int32: int32(id),
			Valid: true,
		},
	}

	if err := husband.DELETE(database); err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}

	Response := map[string]interface{}{
		"message": "تمت العملية بنجاح",
	}

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(Response); err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}
}
