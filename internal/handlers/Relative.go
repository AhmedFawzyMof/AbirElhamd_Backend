package handlers

import (
	"abir-el-hamd/internal/config"
	"abir-el-hamd/internal/middleware"
	"abir-el-hamd/internal/models"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
)

func AddRelative(w http.ResponseWriter, r *http.Request) {
	database := config.Database()
	defer database.Close()

	data := make(map[string]interface{})

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}

	id := data["case_id"].(string)

	age, err := strconv.Atoi(data["relative_age"].(string))

	if err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}

	relative := models.Relatives{
		Type:             sql.NullString{String: data["relative_type"].(string), Valid: true},
		Name:             sql.NullString{String: data["relative_name"].(string), Valid: true},
		National_id:      sql.NullString{String: data["relative_national_id"].(string), Valid: true},
		Date_of_birth:    sql.NullString{String: data["relative_date_of_birth"].(string), Valid: true},
		Age:              sql.NullInt32{Int32: int32(age), Valid: true},
		Gender:           sql.NullString{String: data["relative_gender"].(string), Valid: true},
		Job:              sql.NullString{String: data["relative_job"].(string), Valid: true},
		Social_situation: sql.NullString{String: data["relative_social_situation"].(string), Valid: true},
		Health_status:    sql.NullString{String: data["relative_health_status"].(string), Valid: true},
		Education:        sql.NullString{String: data["relative_education"].(string), Valid: true},
	}

	err = relative.Add(database, id)
	if err != nil {
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

func UpdateRelative(w http.ResponseWriter, r *http.Request) {
	database := config.Database()
	defer database.Close()

	data := make(map[string]interface{})

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}

	case_id := data["case_id"].(string)
	id := int(data["id"].(float64))
	age, err := strconv.Atoi(data["relative_age"].(string))

	if err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}

	relative := models.Relatives{
		Id:               sql.NullInt32{Int32: int32(id), Valid: true},
		Type:             sql.NullString{String: data["relative_type"].(string), Valid: true},
		Name:             sql.NullString{String: data["relative_name"].(string), Valid: true},
		National_id:      sql.NullString{String: data["relative_national_id"].(string), Valid: true},
		Date_of_birth:    sql.NullString{String: data["relative_date_of_birth"].(string), Valid: true},
		Age:              sql.NullInt32{Int32: int32(age), Valid: true},
		Gender:           sql.NullString{String: data["relative_gender"].(string), Valid: true},
		Job:              sql.NullString{String: data["relative_job"].(string), Valid: true},
		Social_situation: sql.NullString{String: data["relative_social_situation"].(string), Valid: true},
		Health_status:    sql.NullString{String: data["relative_health_status"].(string), Valid: true},
		Education:        sql.NullString{String: data["relative_education"].(string), Valid: true},
	}

	err = relative.UPDATE(database, case_id)
	if err != nil {
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
func DeleteRelative(w http.ResponseWriter, r *http.Request) {
	database := config.Database()
	defer database.Close()

	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}

	sub := models.Relatives{
		Id: sql.NullInt32{Int32: int32(id), Valid: true},
	}

	if err := sub.DELETE(database); err != nil {
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
