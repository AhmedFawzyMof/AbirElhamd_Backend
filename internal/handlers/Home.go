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
	"log"
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

	FilterdCases := models.FilterdCases{}

	CasesTable := models.Cases{}

	Cases, err := FilterdCases.FilterCasesByRelativeAge(database, district, from, to, limit, offset)

	if err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}

	Districts, err := CasesTable.GetAllDistinct(database)

	if err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}

	Res := map[string]interface{}{
		"Cases":     Cases,
		"Pages":     len(Cases),
		"Districts": Districts,
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

	err = os.MkdirAll("./uploads/"+id, 0755)

	if err != nil {
		fmt.Println(err)
		middleware.ErrorResopnse(w, err)
		return
	}

	files := r.MultipartForm.File["files"]

	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {

			middleware.ErrorResopnse(w, err)
			return
		}
		defer file.Close()

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

	totalIncome, err := strconv.Atoi(r.FormValue("total_income"))
	if err != nil {
		fmt.Println(err)
		log.Println("Error converting total_income:", err)
		return
	}

	fixedExpenses, err := strconv.Atoi(r.FormValue("fixed_expenses"))
	if err != nil {
		log.Println("Error converting fixed_expenses:", err)
		return
	}

	Age, err := strconv.Atoi(r.FormValue("age"))
	if err != nil {
		log.Println("Error converting age:", err)
		return
	}

	now := time.Now()
	createdAt := sql.NullString{String: now.Format(time.RFC3339), Valid: true}

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
		Age:                           sql.NullInt32{Int32: int32(Age), Valid: Age != 0},
		Gender:                        sql.NullString{String: r.FormValue("gender"), Valid: r.FormValue("gender") != ""},
		Job:                           sql.NullString{String: r.FormValue("job"), Valid: r.FormValue("job") != ""},
		Social_situation:              sql.NullString{String: r.FormValue("social_situation"), Valid: r.FormValue("social_situation") != ""},
		Address_from_national_id_card: sql.NullString{String: r.FormValue("address_from_national_id_card"), Valid: r.FormValue("address_from_national_id_card") != ""},
		Actual_address:                sql.NullString{String: r.FormValue("actual_address"), Valid: r.FormValue("actual_address") != ""},
		District:                      sql.NullString{String: r.FormValue("district"), Valid: r.FormValue("district") != ""},
		PhoneNumbers:                  sql.NullString{String: r.FormValue("phone_numbers"), Valid: r.FormValue("phone_numbers") != ""},
		Created_at:                    createdAt,
		Updated_at:                    createdAt,
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

	Case := models.CaseDitails{
		Id: sql.NullString{
			String: id,
			Valid:  true,
		},
	}

	cas, err := Case.Get(db)

	if err != nil {
		if err.Error() != "sql: no rows in result set" {
			fmt.Println(err)
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
		"Data":     cas,
		"hasFiles": exists,
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

	var casesUpdate map[string]interface{}

	if err := json.NewDecoder(r.Body).Decode(&casesUpdate); err != nil {
		fmt.Println(err)
		middleware.ErrorResopnse(w, err)
		return
	}

	age, err := strconv.Atoi(casesUpdate["age"].(string))

	if err != nil {
		fmt.Println(err)
		middleware.ErrorResopnse(w, err)
		return
	}

	total_income, err := strconv.Atoi(casesUpdate["total_income"].(string))

	if err != nil {
		fmt.Println(err)
		middleware.ErrorResopnse(w, err)
		return
	}

	now := time.Now()
	updated_at := sql.NullString{String: now.Format(time.RFC3339), Valid: true}

	cases := models.Cases{
		Id:                            sql.NullString{String: casesUpdate["id"].(string), Valid: casesUpdate["id"].(string) != ""},
		Case_name:                     sql.NullString{String: casesUpdate["case_name"].(string), Valid: casesUpdate["case_name"].(string) != ""},
		National_id:                   sql.NullString{String: casesUpdate["national_id"].(string), Valid: casesUpdate["national_id"].(string) != ""},
		Devices_needed_for_the_case:   sql.NullString{String: casesUpdate["devices_needed_for_the_case"].(string), Valid: casesUpdate["devices_needed_for_the_case"].(string) != ""},
		Total_income:                  sql.NullInt32{Int32: int32(total_income), Valid: total_income != 0},
		Age:                           sql.NullInt32{Int32: int32(age), Valid: age != 0},
		Gender:                        sql.NullString{String: casesUpdate["gender"].(string), Valid: casesUpdate["gender"].(string) != ""},
		Job:                           sql.NullString{String: casesUpdate["job"].(string), Valid: casesUpdate["job"].(string) != ""},
		Social_situation:              sql.NullString{String: casesUpdate["social_situation"].(string), Valid: casesUpdate["social_situation"].(string) != ""},
		Address_from_national_id_card: sql.NullString{String: casesUpdate["address_from_national_id_card"].(string), Valid: casesUpdate["address_from_national_id_card"].(string) != ""},
		Actual_address:                sql.NullString{String: casesUpdate["actual_address"].(string), Valid: casesUpdate["actual_address"].(string) != ""},
		District:                      sql.NullString{String: casesUpdate["district"].(string), Valid: casesUpdate["district"].(string) != ""},
		Debts:                         sql.NullString{String: casesUpdate["debts"].(string), Valid: casesUpdate["debts"].(string) != ""},
		Pension_from_husband:          sql.NullString{String: casesUpdate["pension_from_husband"].(string), Valid: casesUpdate["pension_from_husband"].(string) != ""},
		Pension_from_father:           sql.NullString{String: casesUpdate["pension_from_father"].(string), Valid: casesUpdate["pension_from_father"].(string) != ""},
		PhoneNumbers:                  sql.NullString{String: casesUpdate["phone_numbers"].(string), Valid: casesUpdate["phone_numbers"].(string) != ""},
		Updated_at:                    updated_at,
	}


	db := config.Database()
	defer db.Close()

	if err := cases.Update(db); err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}

	if err := models.CreateLogs(db, cases.Id.String, "تعديل على معلومات الحالة", userData.Id); err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}

	Response := map[string]interface{}{"status": "تمت العملية بنجاح"}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(Response); err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}
}
