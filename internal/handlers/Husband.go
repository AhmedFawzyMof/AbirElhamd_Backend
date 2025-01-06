package handlers

import (
	"abir-el-hamd/internal/config"
	"abir-el-hamd/internal/middleware"
	"abir-el-hamd/internal/models"
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
		Husband_name:          data["name"].(string),
		Husband_national_id:   data["national_id"].(string),
		Husband_date_of_birth: data["date_of_birth"].(string),
		Husband_age:           age,
		Husband_gender:        data["gender"].(string),
	}

	if err := husband.Add(database, id); err != nil {
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
		Id:                    id,
		Husband_name:          data["name"].(string),
		Husband_national_id:   data["national_id"].(string),
		Husband_date_of_birth: data["date_of_birth"].(string),
		Husband_age:           age,
		Husband_gender:        data["gender"].(string),
	}

	if err := husband.UPDATE(database, case_id); err != nil {
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

	id := r.PathValue("id")

	sub := models.Husband{}

	if err := sub.DELETE(database, id); err != nil {
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
