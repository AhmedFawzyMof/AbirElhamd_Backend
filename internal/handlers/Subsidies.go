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
	subsidies := models.Subsidies{
		Grants_from_outside_the_association:           data["grants_from_outside_the_association"].(string),
		Grants_from_outside_the_association_financial: data["grants_from_outside_the_association_financial"].(string),
		Grants_from_the_association_financial:         data["grants_from_the_association_financial"].(string),
		Grants_from_the_association_inKind:            data["grants_from_the_association_inKind"].(string),
		Total_Subsidies:                               total,
		End_Of_Payment_Date:                           data["end_of_payment_date"].(string),
		Note:                                          data["note"].(string),
	}

	if err := subsidies.Add(database, id); err != nil {
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

	subsidies := models.Subsidies{
		Id:                                  id,
		Grants_from_outside_the_association: data["grants_from_outside_the_association"].(string),
		Grants_from_outside_the_association_financial: data["grants_from_outside_the_association_financial"].(string),
		Grants_from_the_association_financial:         data["grants_from_the_association_financial"].(string),
		Grants_from_the_association_inKind:            data["grants_from_the_association_inKind"].(string),
		Total_Subsidies:                               total,
		End_Of_Payment_Date:                           data["end_of_payment_date"].(string),
		Note:                                          data["note"].(string),
	}

	if err := subsidies.UPDATE(database, case_id); err != nil {
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

	id := r.PathValue("id")

	sub := models.Subsidies{}

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
