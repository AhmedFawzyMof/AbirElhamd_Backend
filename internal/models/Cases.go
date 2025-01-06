package models

import (
	"database/sql"
	"fmt"
	"time"
)

type Cases struct {
	Id                            sql.NullString `json:"id"`
	Case_name                     sql.NullString `json:"case_name"`
	National_id                   sql.NullString `json:"national_id"`
	Devices_needed_for_the_case   sql.NullString `json:"devices_needed_for_the_case"`
	Total_income                  sql.NullInt32  `json:"total_income"`
	Fixed_expenses                sql.NullInt32  `json:"fixed_expenses"`
	Pension_from_husband          sql.NullString `json:"pension_from_husband"`
	Pension_from_father           sql.NullString `json:"pension_from_father"`
	Debts                         sql.NullString `json:"debts"`
	Case_type                     sql.NullString `json:"case_type"`
	Date_of_birth                 sql.NullString `json:"date_of_birth"`
	Age                           sql.NullInt32  `json:"age"`
	Gender                        sql.NullString `json:"gender"`
	Job                           sql.NullString `json:"job"`
	Social_situation              sql.NullString `json:"social_situation"`
	Address_from_national_id_card sql.NullString `json:"address_from_national_id_card"`
	Actual_address                sql.NullString `json:"actual_address"`
	District                      sql.NullString `json:"district"`
	PhoneNumbers                  sql.NullString `json:"phone_numbers"`
	Subsidies_id                  sql.NullInt32  `json:"subsidies_id"`
	Social_status                 sql.NullInt32  `json:"social_status"`
	Husband_id                    sql.NullInt32  `json:"husband_id"`
	Created_at                    sql.NullString `json:"created_at"`
	Updated_at                    sql.NullString `json:"updated_at"`
	Deleted                       sql.NullBool   `json:"deleted"`
	Date_Of_Social_situation      sql.NullTime   `json:"date_of_social_situation"`
	Case_entry_date               sql.NullTime   `json:"case_entry_date"`
	Status_search_update_date     sql.NullTime   `json:"status_search_update_date"`
	Field_research_history        sql.NullTime   `json:"field_research_history"`
}


type Relative struct {
	Relative_id               sql.NullInt32
	Relative_type             sql.NullString
	Relative_name             sql.NullString
	Relative_national_id      sql.NullString
	Relative_date_of_birth    sql.NullString
	Relative_age              sql.NullInt32
	Relative_gender           sql.NullString
	Relative_job              sql.NullString
	Relative_social_situation sql.NullString
	Relative_health_status    sql.NullString
	Relative_education        sql.NullString
}

type CaseDitails struct {
	Id                                            sql.NullString
	Case_name                                     sql.NullString
	National_id                                   sql.NullString
	Devices_needed_for_the_case                   sql.NullString
	Total_income                                  sql.NullInt32
	Fixed_expenses                                sql.NullInt32
	Pension_from_husband                          sql.NullInt32
	Pension_from_father                           sql.NullInt32
	Debts                                         sql.NullString
	Case_type                                     sql.NullString
	Date_of_birth                                 sql.NullString
	Age                                           sql.NullInt32
	Gender                                        sql.NullString
	Job                                           sql.NullString
	Social_situation                              sql.NullString
	Address_from_national_id_card                 sql.NullString
	Actual_address                                sql.NullString
	District                                      sql.NullString
	PhoneNumbers                                  sql.NullString
	Created_at                                    sql.NullString
	Updated_at                                    sql.NullString
	Husband_id                                    sql.NullInt32
	Husband_name                                  sql.NullString
	Husband_national_id                           sql.NullString
	Husband_date_of_birth                         sql.NullString
	Husband_age                                   sql.NullInt32
	Husband_gender                                sql.NullString
	Social_status                                 sql.NullInt32
	Properties                                    sql.NullString
	Health_status                                 sql.NullString
	Education                                     sql.NullString
	Number_of_family_members                      sql.NullInt32
	Number_of_registered_children                 sql.NullInt32
	Total_number_of_children                      sql.NullInt32
	Subsidies_id                                  sql.NullInt32
	Grants_from_outside_the_association           sql.NullString
	Grants_from_outside_the_association_financial sql.NullString
	Grants_from_the_association_financial         sql.NullString
	Grants_from_the_association_inKind            sql.NullString
	Total_Subsidies                               sql.NullInt32
	End_of_payment_date                           sql.NullString
	Note                                          sql.NullString
	Date_Of_Social_situation                      sql.NullTime `json:"date_of_social_situation"`
	Case_entry_date                               sql.NullTime `json:"case_entry_date"`
	Status_search_update_date                     sql.NullTime `json:"status_search_update_date"`
	Field_research_history                        sql.NullTime `json:"field_research_history"`
	Relatives                                     []Relative
}

type FilterdCases struct {
	Id                            sql.NullString `json:"id"`
	Case_name                     sql.NullString `json:"case_name"`
	National_id                   sql.NullString `json:"national_id"`
	Devices_needed_for_the_case   sql.NullString `json:"devices_needed_for_the_case"`
	Total_income                  sql.NullInt32  `json:"total_income"`
	Fixed_expenses                sql.NullInt32  `json:"fixed_expenses"`
	Pension_from_husband          sql.NullString `json:"pension_from_husband"`
	Pension_from_father           sql.NullString `json:"pension_from_father"`
	Debts                         sql.NullString `json:"debts"`
	Case_type                     sql.NullString `json:"case_type"`
	Date_of_birth                 sql.NullString `json:"date_of_birth"`
	Age                           sql.NullInt32  `json:"age"`
	Gender                        sql.NullString `json:"gender"`
	Job                           sql.NullString `json:"job"`
	Social_situation              sql.NullString `json:"social_situation"`
	Address_from_national_id_card sql.NullString `json:"address_from_national_id_card"`
	Actual_address                sql.NullString `json:"actual_address"`
	District                      sql.NullString `json:"district"`
	PhoneNumbers                  sql.NullString `json:"phone_numbers"`
	Subsidies_id                  sql.NullInt32  `json:"subsidies_id"`
	Social_status                 sql.NullInt32  `json:"social_status"`
	Husband_id                    sql.NullInt32  `json:"husband_id"`
	Created_at                    sql.NullString `json:"created_at"`
	Updated_at                    sql.NullString `json:"updated_at"`
	Deleted                       sql.NullBool   `json:"deleted"`
	Date_Of_Social_situation      sql.NullTime   `json:"date_of_social_situation"`
	Case_entry_date               sql.NullTime   `json:"case_entry_date"`
	Status_search_update_date     sql.NullTime   `json:"status_search_update_date"`
	Field_research_history        sql.NullTime   `json:"field_research_history"`
	Relative_id                   sql.NullInt32  `json:"relative_id"`
	Relative_type                 sql.NullString `json:"relative_type"`
	Relative_name                 sql.NullString `json:"relative_name"`
	Relative_national_id          sql.NullString `json:"relative_national_id"`
	Relative_date_of_birth        sql.NullString `json:"relative_date_of_birth"`
	Relative_age                  sql.NullInt32  `json:"relative_age"`
	Relative_gender               sql.NullString `json:"relative_gender"`
	Relative_job                  sql.NullString `json:"relative_job"`
	Relative_social_situation     sql.NullString `json:"relative_social_situation"`
	Relative_health_status        sql.NullString `json:"relative_health_status"`
	Relative_education            sql.NullString `json:"relative_education"`
}

func (ca Cases) Create(db *sql.DB) error {
	_, err := db.Exec("INSERT INTO `cases` (`id`, `case_name`, `national_id`, `devices_needed_for_the_case`, `total_income`, `fixed_expenses`, `pension_from_husband`, `pension_from_father`, `debts`, `case_type`, `date_of_birth`, `age`, `gender`, `job`, `social_situation`, `address_from_national_id_card`, `actual_address`, `district`, `phone_numbers`, `date_of_social_situation`, `created_at`, `updated_at`, `case_entry_date`, `status_search_update_date`, `field_research_history`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		ca.Id, ca.Case_name, ca.National_id, ca.Devices_needed_for_the_case, ca.Total_income, ca.Fixed_expenses,
		ca.Pension_from_husband, ca.Pension_from_father, ca.Debts, ca.Case_type,
		ca.Date_of_birth, ca.Age, ca.Gender, ca.Job, ca.Social_situation,
		ca.Address_from_national_id_card, ca.Actual_address, ca.District, ca.PhoneNumbers, ca.Date_Of_Social_situation, ca.Created_at, ca.Updated_at, ca.Case_entry_date, ca.Status_search_update_date, ca.Field_research_history)

	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("error: %v", err)
	}

	return nil
}

func (ca Cases) Update(db *sql.DB) error {

	_, err := db.Exec("UPDATE `cases` SET `case_name` = ?,`national_id` = ?,`devices_needed_for_the_case` = ?,`total_income` = ?,`fixed_expenses` = ?,`pension_from_husband` = ?,`pension_from_father` = ?,`debts` = ?,`case_type` = ?,`age` = ?,`gender` = ?,`job` = ?,`social_situation` = ?,`address_from_national_id_card` = ?,`actual_address` = ?,`district` = ?, `phone_numbers` = ?, `date_of_social_situation` = ?, `updated_at` = ?, `case_entry_date` = ?, `status_search_update_date` = ?, `field_research_history` = ? WHERE `id` = ?",
		ca.Case_name, ca.National_id, ca.Devices_needed_for_the_case, ca.Total_income, ca.Fixed_expenses,
		ca.Pension_from_husband, ca.Pension_from_father, ca.Debts, ca.Case_type,
		ca.Age, ca.Gender, ca.Job, ca.Social_situation,
		ca.Address_from_national_id_card, ca.Actual_address, ca.District, ca.PhoneNumbers, ca.Date_Of_Social_situation, ca.Updated_at, ca.Case_entry_date, ca.Status_search_update_date, ca.Field_research_history, ca.Id)

	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	return nil
}

func (ca Cases) Delete(db *sql.DB) error {
	_, err := db.Exec("UPDATE `cases` SET `deleted` = 1, `updated_at` = ? WHERE id = ?", time.Now().Format("2006-01-02 15:04:05"), ca.Id)

	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("error: %v", err)
	}

	return nil
}

func (ca CaseDitails) Get(db *sql.DB) (CaseDitails, error) {
	query := `
		SELECT DISTINCT
			cases.id, cases.case_name, cases.national_id, cases.devices_needed_for_the_case, cases.total_income,
			cases.fixed_expenses, cases.pension_from_husband, cases.pension_from_father, cases.debts, cases.case_type,
			cases.date_of_birth, cases.age, cases.gender, cases.job, cases.social_situation, cases.address_from_national_id_card,
			cases.actual_address, cases.district, cases.phone_numbers, cases.created_at, cases.updated_at, cases.date_of_social_situation, cases.case_entry_date, cases.status_search_update_date, cases.field_research_history,
			husband.id, husband.name AS husband_name, husband.national_id AS husband_national_id, husband.date_of_birth AS husband_date_of_birth, husband.age AS husband_age,
			husband.gender AS husband_gender, socialstatusofthecase.id, socialstatusofthecase.properties, socialstatusofthecase.health_status,
			socialstatusofthecase.education, socialstatusofthecase.number_of_family_members, socialstatusofthecase.number_of_registered_children,
			socialstatusofthecase.total_number_of_children, subsidies.id, subsidies.grants_from_outside_the_association,
			subsidies.grants_from_outside_the_association_financial, subsidies.grants_from_the_association_financial,
			subsidies.grants_from_the_association_inKind, subsidies.total_Subsidies, subsidies.end_of_payment_date, subsidies.note
		FROM cases
		LEFT JOIN subsidies ON subsidies.id = cases.subsidies_id
		LEFT JOIN socialstatusofthecase ON socialstatusofthecase.id = cases.social_status
		LEFT JOIN husband ON husband.id = cases.husband_id
		WHERE cases.id = ?
	`

	Case := db.QueryRow(query, ca.Id)

	var cas CaseDitails
	if err := Case.Scan(
		&cas.Id, &cas.Case_name, &cas.National_id, &cas.Devices_needed_for_the_case, &cas.Total_income,
		&cas.Fixed_expenses, &cas.Pension_from_husband, &cas.Pension_from_father, &cas.Debts, &cas.Case_type,
		&cas.Date_of_birth, &cas.Age, &cas.Gender, &cas.Job, &cas.Social_situation, &cas.Address_from_national_id_card,
		&cas.Actual_address, &cas.District, &cas.PhoneNumbers, &cas.Created_at, &cas.Updated_at, &cas.Date_Of_Social_situation,
		&cas.Case_entry_date, &cas.Status_search_update_date, &cas.Field_research_history, &cas.Husband_id, &cas.Husband_name, &cas.Husband_national_id,
		&cas.Husband_date_of_birth, &cas.Husband_age, &cas.Husband_gender, &cas.Social_status, &cas.Properties, &cas.Health_status,
		&cas.Education, &cas.Number_of_family_members, &cas.Number_of_registered_children, &cas.Total_number_of_children, &cas.Subsidies_id,
		&cas.Grants_from_outside_the_association, &cas.Grants_from_outside_the_association_financial, &cas.Grants_from_the_association_financial,
		&cas.Grants_from_the_association_inKind, &cas.Total_Subsidies, &cas.End_of_payment_date,&cas.Note,
	); err != nil {
		fmt.Println(err.Error(), "here")
		return CaseDitails{}, fmt.Errorf("error: %v", err)
	}

	relativesQuery := `
		SELECT DISTINCT relative.id, relative.relative_type, relative.name, relative.national_id, relative.date_of_birth, relative.age,
			relative.gender, relative.job, relative.social_situation, relative.health_status, relative.education
		FROM relative
		WHERE relative.case_id = ?
	`

	rows, err := db.Query(relativesQuery, cas.Id)
	if err != nil {
		fmt.Println(err)
		return cas, fmt.Errorf("error: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var rel Relative

		if err := rows.Scan(
			&rel.Relative_id, &rel.Relative_type, &rel.Relative_name, &rel.Relative_national_id, &rel.Relative_date_of_birth, &rel.Relative_age,
			&rel.Relative_gender, &rel.Relative_job, &rel.Relative_social_situation, &rel.Relative_health_status, &rel.Relative_education,
		); err != nil {
			fmt.Println(err)
			return cas, fmt.Errorf("error: %v", err)
		}
		cas.Relatives = append(cas.Relatives, rel)
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return cas, fmt.Errorf("error: %v", err)
	}

	return cas, nil
}

func (ca Cases) GetAll(db *sql.DB, limit, offset, from, to int, district string) ([]Cases, error) {
	cases := []Cases{}
	query := "SELECT * FROM `cases` WHERE deleted = 0 LIMIT ? OFFSET ? "

	var rows *sql.Rows
	var err error

	if from != 0 && to != 0 && district == "" {
		query = "SELECT * FROM `cases` WHERE age > ? AND age <= ? AND deleted = 0 LIMIT ? OFFSET ?"
		rows, err = db.Query(query, from, to, limit, offset)
	} else if from != 0 && to != 0 && district != "" {
		query = "SELECT * FROM `cases` WHERE district = ? AND age > ? AND age <= ? AND deleted = 0 LIMIT ? OFFSET ?"
		rows, err = db.Query(query, district, from, to, limit, offset)
	} else if from == 0 && to == 0 && district != "" {
		query = "SELECT * FROM `cases` WHERE district = ? AND deleted = 0 LIMIT ? OFFSET ?"
		rows, err = db.Query(query, district, limit, offset)
	} else {
		rows, err = db.Query(query, limit, offset)
	}
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("error: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var cas Cases

		if err := rows.Scan(&cas.Id, &cas.Case_name, &cas.National_id, &cas.Devices_needed_for_the_case, &cas.Total_income,
			&cas.Fixed_expenses, &cas.Pension_from_husband, &cas.Pension_from_father, &cas.Debts, &cas.Case_type,
			&cas.Date_of_birth, &cas.Age, &cas.Gender, &cas.Job, &cas.Social_situation, &cas.Address_from_national_id_card,
			&cas.Actual_address, &cas.District, &cas.PhoneNumbers, &cas.Social_status, &cas.Subsidies_id, &cas.Husband_id, &cas.Created_at, &cas.Updated_at, &cas.Deleted, &cas.Date_Of_Social_situation, &cas.Case_entry_date, &cas.Status_search_update_date, &cas.Field_research_history); err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("error: %v", err)
		}

		cases = append(cases, cas)
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("error: %v", err)
	}

	return cases, nil
}

func (ca FilterdCases) FilterCasesByRelativeAge(db *sql.DB, district string, from, to, limit, offset int) ([]FilterdCases, error) {
	cases := []FilterdCases{}
	var rows *sql.Rows
	var err error

	sql := `SELECT DISTINCT
			cases.id, cases.case_name, cases.national_id, cases.devices_needed_for_the_case, cases.total_income,
			cases.fixed_expenses, cases.pension_from_husband, cases.pension_from_father, cases.debts, cases.case_type,
			cases.date_of_birth, cases.age, cases.gender, cases.job, cases.social_situation, cases.address_from_national_id_card,
			cases.actual_address, cases.district, cases.phone_numbers, cases.created_at, cases.updated_at, cases.date_of_social_situation, cases.case_entry_date, cases.status_search_update_date, cases.field_research_history,
			relative.relative_type, relative.name, relative.national_id, relative.date_of_birth, relative.age,
			relative.gender, relative.job, relative.social_situation, relative.health_status, relative.education FROM cases LEFT JOIN relative ON relative.case_id = cases.id LIMIT ? OFFSET ?`

	if from == 0 && to == 0 && district == "" {
		rows, err = db.Query(sql, limit, offset)
	}

	if from == 0 && to == 0 && district != "" {
		sql = `SELECT DISTINCT
			cases.id, cases.case_name, cases.national_id, cases.devices_needed_for_the_case, cases.total_income,
			cases.fixed_expenses, cases.pension_from_husband, cases.pension_from_father, cases.debts, cases.case_type,
			cases.date_of_birth, cases.age, cases.gender, cases.job, cases.social_situation, cases.address_from_national_id_card,
			cases.actual_address, cases.district, cases.phone_numbers, cases.created_at, cases.updated_at, cases.date_of_social_situation, cases.case_entry_date, cases.status_search_update_date, cases.field_research_history,
			relative.relative_type, relative.name, relative.national_id, relative.date_of_birth, relative.age,
			relative.gender, relative.job, relative.social_situation, relative.health_status, relative.education FROM cases LEFT JOIN relative ON relative.case_id = cases.id WHERE cases.district = ? LIMIT ? OFFSET ?`
		rows, err = db.Query(sql, district, limit, offset)
	}

	if from != 0 && to != 0 && district == "" {
		sql = `SELECT DISTINCT
			cases.id, cases.case_name, cases.national_id, cases.devices_needed_for_the_case, cases.total_income,
			cases.fixed_expenses, cases.pension_from_husband, cases.pension_from_father, cases.debts, cases.case_type,
			cases.date_of_birth, cases.age, cases.gender, cases.job, cases.social_situation, cases.address_from_national_id_card,
			cases.actual_address, cases.district, cases.phone_numbers, cases.created_at, cases.updated_at, cases.date_of_social_situation, cases.case_entry_date, cases.status_search_update_date, cases.field_research_history,
			relative.relative_type, relative.name, relative.national_id, relative.date_of_birth, relative.age,
			relative.gender, relative.job, relative.social_situation, relative.health_status, relative.education FROM cases LEFT JOIN relative ON relative.case_id = cases.id WHERE relative.age BETWEEN ? AND ? LIMIT ? OFFSET ?`
		rows, err = db.Query(sql, from, to, limit, offset)
	}

	if from != 0 && to != 0 && district != "" {
		sql = `SELECT DISTINCT
		cases.id, cases.case_name, cases.national_id, cases.devices_needed_for_the_case, cases.total_income,
		cases.fixed_expenses, cases.pension_from_husband, cases.pension_from_father, cases.debts, cases.case_type,
		cases.date_of_birth, cases.age, cases.gender, cases.job, cases.social_situation, cases.address_from_national_id_card,
		cases.actual_address, cases.district, cases.phone_numbers, cases.created_at, cases.updated_at, cases.date_of_social_situation, cases.case_entry_date, cases.status_search_update_date, cases.field_research_history,
		relative.relative_type, relative.name, relative.national_id, relative.date_of_birth, relative.age,
		relative.gender, relative.job, relative.social_situation, relative.health_status, relative.education FROM cases LEFT JOIN relative ON relative.case_id = cases.id WHERE relative.age BETWEEN ? AND ? AND cases.district = ? LIMIT ? OFFSET ?`

		rows, err = db.Query(sql, from, to, district, limit, offset)
	}

	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("error: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var cas FilterdCases

		if err := rows.Scan(&cas.Id, &cas.Case_name, &cas.National_id, &cas.Devices_needed_for_the_case, &cas.Total_income,
			&cas.Fixed_expenses, &cas.Pension_from_husband, &cas.Pension_from_father, &cas.Debts, &cas.Case_type,
			&cas.Date_of_birth, &cas.Age, &cas.Gender, &cas.Job, &cas.Social_situation, &cas.Address_from_national_id_card,
			&cas.Actual_address, &cas.District, &cas.PhoneNumbers, &cas.Created_at, &cas.Updated_at, &cas.Date_Of_Social_situation, &cas.Case_entry_date, &cas.Status_search_update_date, &cas.Field_research_history, &cas.Relative_type, &cas.Relative_name, &cas.Relative_national_id, &cas.Relative_date_of_birth, &cas.Relative_age, &cas.Relative_gender, &cas.Relative_job, &cas.Relative_social_situation, &cas.Relative_health_status, &cas.Relative_education); err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("error: %v", err)
		}
		cases = append(cases, cas)
	}

	return cases, nil
}

func (ca Cases) GetAllDistinct(db *sql.DB) ([]string, error) {
	districts := []string{}

	rows, err := db.Query("SELECT DISTINCT district FROM `cases` WHERE district != ''")

	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("error: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var district string

		if err := rows.Scan(&district); err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("error: %v", err)
		}

		districts = append(districts, district)
	}

	return districts, nil
}

func (ca Cases) NumberOfPages(db *sql.DB, district string, from, to int) (int, error) {

	query := "SELECT COUNT(*) AS length FROM `cases`"

	var row *sql.Row

	if district != "" {
		query = "SELECT COUNT(*) AS length FROM `cases` WHERE district = ?"
		row = db.QueryRow(query, district)
	} else if from != 0 && to != 0 {
		query = "SELECT COUNT(DISTINCT cases.id) AS length FROM cases LEFT JOIN relative ON relative.case_id = cases.id WHERE relative.age > ? AND relative.age <= ?"
		row = db.QueryRow(query, from, to)
	} else {
		row = db.QueryRow(query)
	}

	var length int
	if err := row.Scan(&length); err != nil {
		return 0, fmt.Errorf("error: %v", err)
	}

	return length, nil
}

func (ca Cases) Search(db *sql.DB, SearchQuery string) ([]Cases, error) {
	cases := []Cases{}

	rows, err := db.Query("SELECT * FROM `cases` WHERE deleted = 0 AND case_name LIKE ? OR national_id LIKE ? OR devices_needed_for_the_case LIKE ? OR id LIKE ?", SearchQuery, SearchQuery, SearchQuery, SearchQuery)

	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("error: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var cas Cases

		if err := rows.Scan(&cas.Id, &cas.Case_name, &cas.National_id, &cas.Devices_needed_for_the_case, &cas.Total_income,
			&cas.Fixed_expenses, &cas.Pension_from_husband, &cas.Pension_from_father, &cas.Debts,
			&cas.Case_type, &cas.Date_of_birth, &cas.Age, &cas.Gender, &cas.Job, &cas.Social_situation,
			&cas.Address_from_national_id_card, &cas.Actual_address, &cas.District, &cas.PhoneNumbers, &cas.Subsidies_id,
			&cas.Social_status, &cas.Husband_id, &cas.Created_at, &cas.Updated_at, &cas.Deleted, &cas.Date_Of_Social_situation, &cas.Case_entry_date, &cas.Status_search_update_date, &cas.Field_research_history); err != nil {
			return nil, fmt.Errorf("error: %v", err)
		}

		cases = append(cases, cas)
	}

	return cases, nil
}

func (ca Cases) DeletedCases(db *sql.DB) ([]Cases, error) {
	cases := []Cases{}

	rows, err := db.Query("SELECT * FROM `cases` WHERE deleted = 1")

	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("error: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var cas Cases

		if err := rows.Scan(&cas.Id, &cas.Case_name, &cas.National_id, &cas.Devices_needed_for_the_case, &cas.Total_income,
			&cas.Fixed_expenses, &cas.Pension_from_husband, &cas.Pension_from_father, &cas.Debts,
			&cas.Case_type, &cas.Date_of_birth, &cas.Age, &cas.Gender, &cas.Job, &cas.Social_situation,
			&cas.Address_from_national_id_card, &cas.Actual_address, &cas.District, &cas.PhoneNumbers, &cas.Subsidies_id,
			&cas.Social_status, &cas.Husband_id, &cas.Created_at, &cas.Updated_at, &cas.Deleted, &cas.Date_Of_Social_situation, &cas.Case_entry_date, &cas.Status_search_update_date, &cas.Field_research_history); err != nil {
			return nil, fmt.Errorf("error: %v", err)
		}

		cases = append(cases, cas)
	}

	return cases, nil
}
