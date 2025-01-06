package handlers

import (
	"abir-el-hamd/internal/config"
	"abir-el-hamd/internal/middleware"
	"abir-el-hamd/internal/models"
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
		Type:             data["relative_type"].(string),
		Name:             data["relative_name"].(string),
		National_id:      data["relative_national_id"].(string),
		Date_of_birth:    data["relative_date_of_birth"].(string),
		Age:              age,
		Gender:           data["relative_gender"].(string),
		Job:              data["relative_job"].(string),
		Social_situation: data["relative_social_situation"].(string),
		Health_status:    data["relative_health_status"].(string),
		Education:        data["relative_education"].(string),
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
		Id:               id,
		Type:             data["relative_type"].(string),
		Name:             data["relative_name"].(string),
		National_id:      data["relative_national_id"].(string),
		Date_of_birth:    data["relative_date_of_birth"].(string),
		Age:              age,
		Gender:           data["relative_gender"].(string),
		Job:              data["relative_job"].(string),
		Social_situation: data["relative_social_situation"].(string),
		Health_status:    data["relative_health_status"].(string),
		Education:        data["relative_education"].(string),
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
		Id: id,
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
