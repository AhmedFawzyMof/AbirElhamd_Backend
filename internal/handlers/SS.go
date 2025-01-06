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

func AddSS(w http.ResponseWriter, r *http.Request) {
	database := config.Database()
	defer database.Close()

	data := make(map[string]interface{})

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		fmt.Println(err)
		middleware.ErrorResopnse(w, err)
		return
	}


	id := data["case_id"].(string)

	nofm, err := strconv.Atoi(data["number_of_family_members"].(string))
	if err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}
	norc, err := strconv.Atoi(data["number_of_registered_children"].(string))
	if err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}
	tnoc, err := strconv.Atoi(data["notal_number_of_children"].(string))
	if err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}

	s := models.SS{
		Properties:                    data["properties"].(string),
		Health_status:                 data["health_status"].(string),
		Education:                     data["education"].(string),
		Number_of_family_members:      nofm,
		Number_of_registered_children: norc,
		Total_number_of_children:      tnoc,
	}

	if err := s.Add(database, id); err != nil {
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

func UpdateSS(w http.ResponseWriter, r *http.Request) {
	database := config.Database()
	defer database.Close()

	data := make(map[string]interface{})

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		fmt.Println(err)
		middleware.ErrorResopnse(w, err)
		return
	}


	case_id := data["case_id"].(string)

	id := int(data["id"].(float64))

	nofm, err := strconv.Atoi(data["number_of_family_members"].(string))
	if err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}
	norc, err := strconv.Atoi(data["number_of_registered_children"].(string))
	if err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}
	tnoc, err := strconv.Atoi(data["notal_number_of_children"].(string))
	if err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}

	s := models.SS{
		Id:                            id,
		Properties:                    data["properties"].(string),
		Health_status:                 data["health_status"].(string),
		Education:                     data["education"].(string),
		Number_of_family_members:      nofm,
		Number_of_registered_children: norc,
		Total_number_of_children:      tnoc,
	}

	if err := s.UPDATE(database, case_id); err != nil {
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


func DeleteSS(w http.ResponseWriter, r *http.Request) {
	database := config.Database()
	defer database.Close()

	id := r.PathValue("id")

	ss := models.SS{}

	if err := ss.DELETE(database, id); err != nil {
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