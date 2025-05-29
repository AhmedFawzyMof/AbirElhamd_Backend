package handlers

import (
	"abir-el-hamd/internal/config"
	"abir-el-hamd/internal/middleware"
	"abir-el-hamd/internal/models"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
)

func HomeApi(w http.ResponseWriter, r *http.Request) {
	database := config.Database()
	defer database.Close()

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))

	if err != nil {
		er := errors.New("invalid number")
		middleware.ErrorResopnse(w, er)
		return
	}

	from, err := strconv.Atoi(r.URL.Query().Get("from"))

	if err != nil {
		er := errors.New("invalid number")
		middleware.ErrorResopnse(w, er)
		return
	}

	to, err := strconv.Atoi(r.URL.Query().Get("to"))

	if err != nil {
		er := errors.New("invalid number")
		middleware.ErrorResopnse(w, er)
		return
	}

	district := r.URL.Query().Get("district")

	offset := limit - 30

	CasesTable := models.Cases{}

	Cases, err := CasesTable.GetAll(database, 30, offset, from, to, district)

	if err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}

	Districts, err := CasesTable.GetAllDistinct(database)

	if err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}

	pages, err := CasesTable.NumberOfPages(database, district, from, to)

	if err != nil {
		fmt.Println(err)
		middleware.ErrorResopnse(w, err)
		return
	}

	Res := map[string]interface{}{
		"Cases":     Cases,
		"Pages":     pages,
		"Districts": Districts,
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(Res); err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}
}

func DeletedCases(w http.ResponseWriter, r *http.Request) {
	database := config.Database()
	defer database.Close()

	CasesTable := models.Cases{}

	Cases, err := CasesTable.DeletedCases(database)

	if err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(Cases); err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}
}

func FilterKids(w http.ResponseWriter, r *http.Request) {
	database := config.Database()
	defer database.Close()

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))

	if err != nil {
		er := errors.New("invalid number")
		middleware.ErrorResopnse(w, er)
		return
	}

	from, err := strconv.Atoi(r.URL.Query().Get("from"))

	if err != nil {
		er := errors.New("invalid number")
		middleware.ErrorResopnse(w, er)
		return
	}

	to, err := strconv.Atoi(r.URL.Query().Get("to"))

	if err != nil {
		er := errors.New("invalid number")
		middleware.ErrorResopnse(w, er)
		return
	}

	district := r.URL.Query().Get("district")

	offset := limit - 30

	FilterdCases := models.Cases{}
	CasesTable := models.Cases{}
	RelativesTable := models.Relatives{}

	Cases, err := FilterdCases.FilterCasesByRelativeAge(database, district, from, to, 30, offset)

	if err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}

	Districts, err := CasesTable.GetAllDistinct(database)

	if err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}

	ids := []any{}
	for _, c := range Cases {
		ids = append(ids, c.Id.String)
	}

	relatives, err := RelativesTable.GetByCasesIDS(database, ids)

	if err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}

	Res := map[string]interface{}{
		"Cases":     Cases,
		"Pages":     len(Cases),
		"Districts": Districts,
		"Relatives": relatives,
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(Res); err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}
}

func AddCase(w http.ResponseWriter, r *http.Request) {
	token, err := middleware.GetToken(r)

	if err != nil {
		fmt.Println(err)
		middleware.ErrorResopnse(w, err)
		return
	}

	userData, err := middleware.ValidateToken(token)

	if err != nil {
		fmt.Println(err)
		middleware.ErrorResopnse(w, err)
		return
	}

	db := config.Database()
	defer db.Close()

	err = r.ParseMultipartForm(30 << 20)

	if err != nil {
		fmt.Println(err)
		middleware.ErrorResopnse(w, err)
		return
	}
	id := uuid.New().String()

	files := r.MultipartForm.File["files"]

	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			middleware.ErrorResopnse(w, err)
			return
		}
		defer file.Close()

		err = os.MkdirAll("./uploads/"+id, 0755)

		if err != nil {
			fmt.Println(err)
			middleware.ErrorResopnse(w, err)
			return
		}

		dst, err := os.Create("./uploads/" + id + "/" + fileHeader.Filename)
		if err != nil {
			fmt.Println(err)
			middleware.ErrorResopnse(w, err)
			return
		}
		defer dst.Close()

		if _, err := io.Copy(dst, file); err != nil {
			fmt.Println(err)
			middleware.ErrorResopnse(w, err)
			return
		}
	}

	var totalIncome int
	var fixedExpenses int
	var age int
	var subsidiesID int
	var socialStatus int
	var husbandID int

	if r.FormValue("total_income") != "" {
		totalIncome, err = strconv.Atoi(r.FormValue("total_income"))
		if err != nil {
			fmt.Println("Error converting total_income:", err)
			middleware.ErrorResopnse(w, err)
			return
		}
	}

	if r.FormValue("fixed_expenses") != "" {
		fixedExpenses, err = strconv.Atoi(r.FormValue("fixed_expenses"))
		if err != nil {
			fmt.Println("Error converting fixed_expenses:", err)
			middleware.ErrorResopnse(w, err)
			return
		}
	}

	if r.FormValue("age") != "" {
		age, err = strconv.Atoi(r.FormValue("age"))
		if err != nil {
			fmt.Println("Error converting age:", err)
			middleware.ErrorResopnse(w, err)
			return
		}
	}

	if r.FormValue("subsidies_id") != "" {
		subsidiesID, err = strconv.Atoi(r.FormValue("subsidies_id"))
		if err != nil {
			fmt.Println("Error converting subsidies_id:", err)
			middleware.ErrorResopnse(w, err)
			return
		}
	}

	if r.FormValue("social_status") != "" {
		socialStatus, err = strconv.Atoi(r.FormValue("social_status"))
		if err != nil {
			fmt.Println("Error converting social_status:", err)
			middleware.ErrorResopnse(w, err)
			return
		}
	}

	if r.FormValue("husband_id") != "" {
		husbandID, err = strconv.Atoi(r.FormValue("husband_id"))
		if err != nil {
			fmt.Println("Error converting husband_id:", err)
			middleware.ErrorResopnse(w, err)
			return
		}
	}

	egyptLoc, err := time.LoadLocation("Africa/Cairo")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return
	}

	now := time.Now()
	currentTime := sql.NullString{String: now.In(egyptLoc).Format(time.RFC3339), Valid: true}

	var dateOfSocialSituation sql.NullTime
	if len(r.MultipartForm.Value["date_of_social_situation"]) > 0 {
		if r.MultipartForm.Value["date_of_social_situation"][0] != "" {
			parsedTime, err := time.ParseInLocation("2006-01-02", r.MultipartForm.Value["date_of_social_situation"][0], egyptLoc)
			if err == nil {
				dateOfSocialSituation = sql.NullTime{Time: parsedTime, Valid: true}
			}
		}
	}

	var caseEntryDate sql.NullTime
	if len(r.MultipartForm.Value["case_entry_date"]) > 0 {
		if r.MultipartForm.Value["case_entry_date"][0] != "" {
			parsedTime, err := time.ParseInLocation("2006-01-02", r.MultipartForm.Value["case_entry_date"][0], egyptLoc)
			if err == nil {
				caseEntryDate = sql.NullTime{Time: parsedTime, Valid: true}
			}
		}
	}

	var statusSearchUpdateDate sql.NullTime
	if len(r.MultipartForm.Value["status_search_update_date"]) > 0 {
		if r.MultipartForm.Value["status_search_update_date"][0] != "" {
			parsedTime, err := time.ParseInLocation("2006-01-02", r.MultipartForm.Value["status_search_update_date"][0], egyptLoc)
			if err == nil {
				statusSearchUpdateDate = sql.NullTime{Time: parsedTime, Valid: true}
			}
		}
	}

	var fieldResearchHistory sql.NullTime
	if len(r.MultipartForm.Value["field_research_history"]) > 0 {
		if r.MultipartForm.Value["field_research_history"][0] != "" {
			parsedTime, err := time.ParseInLocation("2006-01-02", r.MultipartForm.Value["field_research_history"][0], egyptLoc)
			if err == nil {
				fieldResearchHistory = sql.NullTime{Time: parsedTime, Valid: true}
			}
		}
	}

	Case := models.Cases{
		Id:                            sql.NullString{String: id, Valid: id != ""},
		Case_name:                     sql.NullString{String: r.FormValue("case_name"), Valid: r.FormValue("case_name") != ""},
		National_id:                   sql.NullString{String: r.FormValue("national_id"), Valid: r.FormValue("national_id") != ""},
		Devices_needed_for_the_case:   sql.NullString{String: r.FormValue("devices_needed_for_the_case"), Valid: r.FormValue("devices_needed_for_the_case") != ""},
		Total_income:                  sql.NullInt32{Int32: int32(totalIncome), Valid: totalIncome != 0},
		Fixed_expenses:                sql.NullInt32{Int32: int32(fixedExpenses), Valid: fixedExpenses != 0},
		Pension_from_husband:          sql.NullString{String: r.FormValue("pension_from_husband"), Valid: r.FormValue("pension_from_husband") != ""},
		Pension_from_father:           sql.NullString{String: r.FormValue("pension_from_father"), Valid: r.FormValue("pension_from_father") != ""},
		Debts:                         sql.NullString{String: r.FormValue("debts"), Valid: r.FormValue("debts") != ""},
		Case_type:                     sql.NullString{String: r.FormValue("case_type"), Valid: r.FormValue("case_type") != ""},
		Date_of_birth:                 sql.NullString{String: r.FormValue("date_of_birth"), Valid: r.FormValue("date_of_birth") != ""},
		Age:                           sql.NullInt32{Int32: int32(age), Valid: age != 0},
		Gender:                        sql.NullString{String: r.FormValue("gender"), Valid: r.FormValue("gender") != ""},
		Job:                           sql.NullString{String: r.FormValue("job"), Valid: r.FormValue("job") != ""},
		Social_situation:              sql.NullString{String: r.FormValue("social_situation"), Valid: r.FormValue("social_situation") != ""},
		Address_from_national_id_card: sql.NullString{String: r.FormValue("address_from_national_id_card"), Valid: r.FormValue("address_from_national_id_card") != ""},
		Actual_address:                sql.NullString{String: r.FormValue("actual_address"), Valid: r.FormValue("actual_address") != ""},
		District:                      sql.NullString{String: r.FormValue("district"), Valid: r.FormValue("district") != ""},
		PhoneNumbers:                  sql.NullString{String: r.FormValue("phone_numbers"), Valid: r.FormValue("phone_numbers") != ""},
		Subsidies_id:                  sql.NullInt32{Int32: int32(subsidiesID), Valid: subsidiesID != 0},
		Social_status:                 sql.NullInt32{Int32: int32(socialStatus), Valid: socialStatus != 0},
		Husband_id:                    sql.NullInt32{Int32: int32(husbandID), Valid: husbandID != 0},
		Created_at:                    currentTime,
		Updated_at:                    currentTime,
		Date_Of_Social_situation:      dateOfSocialSituation,
		Case_entry_date:               caseEntryDate,
		Status_search_update_date:     statusSearchUpdateDate,
		Field_research_history:        fieldResearchHistory,
		Living_expenses:               sql.NullString{String: r.FormValue("living_expenses"), Valid: r.FormValue("living_expenses") != ""},
	}

	if err := Case.Create(db); err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}

	if err := models.CreateLogs(db, id, "إنشاء حالة جديدة", userData.Id); err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(map[string]interface{}{"status": "تمت إضافة الحالة بنجاح"}); err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}
}

func DeleteCase(w http.ResponseWriter, r *http.Request) {
	token, err := middleware.GetToken(r)

	if err != nil {
		fmt.Println(err)
		middleware.ErrorResopnse(w, err)
		return
	}

	userData, err := middleware.ValidateToken(token)

	if err != nil {
		fmt.Println(err)
		middleware.ErrorResopnse(w, err)
		return
	}

	middleware.VerifyAdmin(w, r)

	db := config.Database()
	defer db.Close()

	id := r.PathValue("id")

	Case := models.Cases{
		Id: sql.NullString{
			String: id,
			Valid:  true,
		},
	}

	if err := Case.Delete(db); err != nil {
		fmt.Println(err)
		middleware.ErrorResopnse(w, err)
		return
	}

	Response := map[string]interface{}{}
	Response["status"] = "تمت العملية بنجاح"

	if err := models.CreateLogs(db, Case.Id.String, "حذف الحالة", userData.Id); err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(Response); err != nil {
		fmt.Println(err)
		middleware.ErrorResopnse(w, err)
		return
	}
}

func CaseApi(w http.ResponseWriter, r *http.Request) {
	db := config.Database()
	defer db.Close()

	id := r.PathValue("id")

	Case := models.Cases{
		Id: sql.NullString{
			String: id,
			Valid:  true,
		},
	}

	cas, err := Case.Get(db)

	if err != nil {
		fmt.Println("case", err)
		if err.Error() != "sql: no rows in result set" {
			middleware.ErrorResopnse(w, err)
			return
		}
	}

	RelativesModel := models.Relatives{
		Case_id: sql.NullString{
			String: cas.Id.String,
			Valid:  true,
		},
	}

	relatives, err := RelativesModel.GetByCaseID(db, cas.Id.String)

	if err != nil {
		fmt.Println("relatives", err)
		middleware.ErrorResopnse(w, err)
		return
	}

	HusbandsModel := models.Husband{
		Case_id: sql.NullString{
			String: cas.Id.String,
			Valid:  true,
		},
	}

	husbands, err := HusbandsModel.GetByCaseID(db)

	if err != nil {
		fmt.Println("husband", err)
		middleware.ErrorResopnse(w, err)
		return
	}

	SubsidiesModel := models.Subsidies{
		Case_id: sql.NullString{
			String: cas.Id.String,
			Valid:  true,
		},
	}

	subsidies, err := SubsidiesModel.GetByCaseID(db)

	if err != nil {
		fmt.Println("subsidies", err)
		middleware.ErrorResopnse(w, err)
		return
	}

	SocialSituation := models.SS{
		Case_id: sql.NullString{
			String: cas.Id.String,
			Valid:  true,
		},
	}

	ss, err := SocialSituation.GetByCaseID(db)

	if err != nil {
		fmt.Println("ss", err)
		if err.Error() != "sql: no rows in result set" {
			middleware.ErrorResopnse(w, err)
			return
		}
	}

	path := fmt.Sprintf("uploads/%s", id)

	exists, err := middleware.Exists(path)

	errText := fmt.Sprintf("CreateFile %s: The system cannot find the file specified.", path)

	if err != nil {
		if err.Error() != errText {
			fmt.Println(err)
			middleware.ErrorResopnse(w, err)
			return
		}
	}

	Response := map[string]interface{}{
		"Case":         cas,
		"Relatives":    relatives,
		"Husbands":     husbands,
		"Subsidies":    subsidies,
		"SocialStatus": ss,
		"hasFiles":     exists,
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(Response); err != nil {
		fmt.Println(err)
		middleware.ErrorResopnse(w, err)
		return
	}
}

func SearchCase(w http.ResponseWriter, r *http.Request) {
	db := config.Database()
	defer db.Close()

	search := "%" + r.URL.Query().Get("search") + "%"

	Case := models.Cases{}

	cases, err := Case.Search(db, search)

	if err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(cases); err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}

}

func UpdateCase(w http.ResponseWriter, r *http.Request) {
	token, err := middleware.GetToken(r)
	if err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}

	userData, err := middleware.ValidateToken(token)
	if err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}

	db := config.Database()
	defer db.Close()

	var body map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		middleware.ErrorResopnse(w, errors.New("invalid JSON body"))
		return
	}

	getString := func(key string) sql.NullString {
		if val, ok := body[key].(string); ok && val != "" {
			return sql.NullString{String: val, Valid: true}
		}
		return sql.NullString{}
	}

	getInt := func(key string) sql.NullInt32 {
		if val, ok := body[key].(float64); ok {
			return sql.NullInt32{Int32: int32(val), Valid: true}
		}
		return sql.NullInt32{}
	}

	getBool := func(key string) sql.NullBool {
		if val, ok := body[key].(bool); ok {
			return sql.NullBool{Bool: val, Valid: true}
		}
		return sql.NullBool{}
	}

	getDate := func(key string) sql.NullTime {
		if val, ok := body[key].(string); ok && val != "" {
			parsed, err := time.Parse("2006-01-02", val)
			if err == nil {
				return sql.NullTime{Time: parsed, Valid: true}
			}
		}
		return sql.NullTime{}
	}

	caseData := models.Cases{
		Id:                            getString("id"),
		Case_name:                     getString("case_name"),
		National_id:                   getString("national_id"),
		Devices_needed_for_the_case:   getString("devices_needed_for_the_case"),
		Total_income:                  getInt("total_income"),
		Fixed_expenses:                getInt("fixed_expenses"),
		Pension_from_husband:          getString("pension_from_husband"),
		Pension_from_father:           getString("pension_from_father"),
		Debts:                         getString("debts"),
		Case_type:                     getString("case_type"),
		Date_of_birth:                 getString("date_of_birth"),
		Age:                           getInt("age"),
		Gender:                        getString("gender"),
		Job:                           getString("job"),
		Social_situation:              getString("social_situation"),
		Address_from_national_id_card: getString("address_from_national_id_card"),
		Actual_address:                getString("actual_address"),
		District:                      getString("district"),
		PhoneNumbers:                  getString("phone_numbers"),
		Subsidies_id:                  getInt("subsidies_id"),
		Social_status:                 getInt("social_status"),
		Husband_id:                    getInt("husband_id"),
		Deleted:                       getBool("deleted"),
		Date_Of_Social_situation:      getDate("date_of_social_situation"),
		Case_entry_date:               getDate("case_entry_date"),
		Status_search_update_date:     getDate("status_search_update_date"),
		Field_research_history:        getDate("field_research_history"),
		Living_expenses:               getString("living_expenses"),
	}

	if !caseData.Id.Valid || caseData.Id.String == "" {
		middleware.ErrorResopnse(w, errors.New("missing case id"))
		return
	}

	egyptLoc, err := time.LoadLocation("Africa/Cairo")
	if err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}

	caseData.Updated_at = sql.NullString{
		String: time.Now().In(egyptLoc).Format(time.RFC3339),
		Valid:  true,
	}

	err = caseData.Update(db)
	if err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}

	err = models.CreateLogs(db, caseData.Id.String, "تعديل على معلومات الحالة", userData.Id)
	if err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"status": "تم تعديل الحالة بنجاح"})
}
