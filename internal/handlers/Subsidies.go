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
	"time"
)

func AddSubsidies(w http.ResponseWriter, r *http.Request) {
	database := config.Database()
	defer database.Close()

	data := make(map[string]interface{})

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		fmt.Println(err)
		middleware.ErrorResopnse(w, err)
		return
	}

	id := data["case_id"].(string)

	total, err := strconv.Atoi(data["total_subsidies"].(string))

	if err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}

	dateStr := data["end_of_payment_date"].(string)
	if len(dateStr) == 10 {
		dateStr += "T00:00:00Z"
	}

	t, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}

	end_Of_Payment_Date := sql.NullTime{Time: t, Valid: true}

	subsidies := models.Subsidies{
		Grants_from_outside_the_association:           sql.NullString{String: data["grants_from_outside_the_association"].(string), Valid: true},
		Grants_from_outside_the_association_financial: sql.NullString{String: data["grants_from_outside_the_association_financial"].(string), Valid: true},
		Grants_from_the_association_financial:         sql.NullString{String: data["grants_from_the_association_financial"].(string), Valid: true},
		Grants_from_the_association_inKind:            sql.NullString{String: data["grants_from_the_association_inKind"].(string), Valid: true},
		Total_Subsidies:                               sql.NullInt64{Int64: int64(total), Valid: true},
		End_Of_Payment_Date:                           end_Of_Payment_Date,
		Note:                                          sql.NullString{String: data["note"].(string), Valid: true},
		Case_id:                                       sql.NullString{String: id, Valid: true},
	}

	if err := subsidies.Add(database); err != nil {
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

func UpdateSubsidies(w http.ResponseWriter, r *http.Request) {
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

	total, err := strconv.Atoi(data["total_subsidies"].(string))

	if err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}

		dateStr := data["end_of_payment_date"].(string)
	if len(dateStr) == 10 {
		dateStr += "T00:00:00Z"
	}

	t, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}

	end_Of_Payment_Date := sql.NullTime{Time: t, Valid: true}

	subsidies := models.Subsidies{
		Id:                                  sql.NullInt64{Int64: int64(id), Valid: true},
		Grants_from_outside_the_association: sql.NullString{String: data["grants_from_outside_the_association"].(string), Valid: true},
		Grants_from_outside_the_association_financial: sql.NullString{String: data["grants_from_outside_the_association_financial"].(string), Valid: true},
		Grants_from_the_association_financial:         sql.NullString{String: data["grants_from_the_association_financial"].(string), Valid: true},
		Grants_from_the_association_inKind:            sql.NullString{String: data["grants_from_the_association_inKind"].(string), Valid: true},
		Total_Subsidies:                               sql.NullInt64{Int64: int64(total), Valid: true},
		End_Of_Payment_Date:                           end_Of_Payment_Date,
		Note:                                          sql.NullString{String: data["note"].(string), Valid: true},
		Case_id:                                       sql.NullString{String: case_id, Valid: true},
	}

	if err := subsidies.UPDATE(database); err != nil {
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
func DeleteSubsidies(w http.ResponseWriter, r *http.Request) {
	database := config.Database()
	defer database.Close()

	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		fmt.Println(err)
		middleware.ErrorResopnse(w, err)
		return
	}

	sub := models.Subsidies{
		Id: sql.NullInt64{
			Int64: int64(id),
			Valid: true,
		},
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
