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
		Properties:                    sql.NullString{String: data["properties"].(string), Valid: true},
		Health_status:                 sql.NullString{String: data["health_status"].(string), Valid: true},
		Education:                     sql.NullString{String: data["education"].(string), Valid: true},
		Number_of_family_members:      sql.NullInt64{Int64: int64(nofm), Valid: true},
		Number_of_registered_children: sql.NullInt64{Int64: int64(norc), Valid: true},
		Total_number_of_children:      sql.NullInt64{Int64: int64(tnoc), Valid: true},
		Case_id:                       sql.NullString{String: id, Valid: true},
	}

	if err := s.Add(database); err != nil {
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
		Id:                            sql.NullInt64{Int64: int64(id), Valid: true},
		Properties:                    sql.NullString{String: data["properties"].(string), Valid: true},
		Health_status:                 sql.NullString{String: data["health_status"].(string), Valid: true},
		Education:                     sql.NullString{String: data["education"].(string), Valid: true},
		Number_of_family_members:      sql.NullInt64{Int64: int64(nofm), Valid: true},
		Number_of_registered_children: sql.NullInt64{Int64: int64(norc), Valid: true},
		Total_number_of_children:      sql.NullInt64{Int64: int64(tnoc), Valid: true},
		Case_id:                       sql.NullString{String: case_id, Valid: true},
	}

	if err := s.UPDATE(database); err != nil {
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

	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}

	ss := models.SS{
		Id: sql.NullInt64{
			Int64: int64(id),
			Valid: true,
		},
	}

	if err := ss.DELETE(database); err != nil {
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
