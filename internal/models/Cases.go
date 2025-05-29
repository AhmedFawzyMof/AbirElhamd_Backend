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
	Living_expenses               sql.NullString  `json:"living_expenses"`
}

func (ca *Cases) Create(db *sql.DB) error {
	fmt.Println(ca)
	_, err := db.Exec("INSERT INTO `cases` (`id`, `case_name`, `national_id`, `devices_needed_for_the_case`, `total_income`, `fixed_expenses`, `pension_from_husband`, `pension_from_father`, `debts`, `case_type`, `date_of_birth`, `age`, `gender`, `job`, `social_situation`, `address_from_national_id_card`, `actual_address`, `district`, `phone_numbers`, `date_of_social_situation`, `created_at`, `updated_at`, `case_entry_date`, `status_search_update_date`, `field_research_history`, `living_expenses`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		ca.Id, ca.Case_name, ca.National_id, ca.Devices_needed_for_the_case, ca.Total_income, ca.Fixed_expenses,
		ca.Pension_from_husband, ca.Pension_from_father, ca.Debts, ca.Case_type,
		ca.Date_of_birth, ca.Age, ca.Gender, ca.Job, ca.Social_situation,
		ca.Address_from_national_id_card, ca.Actual_address, ca.District, ca.PhoneNumbers, ca.Date_Of_Social_situation, ca.Created_at, ca.Updated_at, ca.Case_entry_date, ca.Status_search_update_date, ca.Field_research_history, ca.Living_expenses)

	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("error: %v", err)
	}

	return nil
}

func (ca Cases) Update(db *sql.DB) error {
	_, err := db.Exec(`
		UPDATE cases SET 
			case_name = ?,
			national_id = ?,
			devices_needed_for_the_case = ?,
			total_income = ?,
			fixed_expenses = ?,
			pension_from_husband = ?,
			pension_from_father = ?,
			debts = ?,
			case_type = ?,
			date_of_birth = ?,
			age = ?,
			gender = ?,
			job = ?,
			social_situation = ?,
			address_from_national_id_card = ?,
			actual_address = ?,
			district = ?,
			phone_numbers = ?,
			subsidies_id = ?,
			social_status = ?,
			husband_id = ?,
			updated_at = ?,
			deleted = ?,
			date_of_social_situation = ?,
			case_entry_date = ?,
			status_search_update_date = ?,
			field_research_history = ?,
			living_expenses = ?
		WHERE id = ?`,
		ca.Case_name,
		ca.National_id,
		ca.Devices_needed_for_the_case,
		ca.Total_income,
		ca.Fixed_expenses,
		ca.Pension_from_husband,
		ca.Pension_from_father,
		ca.Debts,
		ca.Case_type,
		ca.Date_of_birth,
		ca.Age,
		ca.Gender,
		ca.Job,
		ca.Social_situation,
		ca.Address_from_national_id_card,
		ca.Actual_address,
		ca.District,
		ca.PhoneNumbers,
		ca.Subsidies_id,
		ca.Social_status,
		ca.Husband_id,
		ca.Updated_at,
		ca.Deleted,
		ca.Date_Of_Social_situation,
		ca.Case_entry_date,
		ca.Status_search_update_date,
		ca.Field_research_history,
		ca.Living_expenses,
		ca.Id,
	)

	if err != nil {
		return fmt.Errorf("update error: %v", err)
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

func (ca Cases) Get(db *sql.DB) (Cases, error) {
	query := `SELECT * FROM cases WHERE id = ?`

	Case := db.QueryRow(query, ca.Id)

	var cas Cases
	if err := Case.Scan(&cas.Id, &cas.Case_name, &cas.National_id, &cas.Devices_needed_for_the_case, &cas.Total_income,
		&cas.Fixed_expenses, &cas.Pension_from_husband, &cas.Pension_from_father, &cas.Debts,
		&cas.Case_type, &cas.Date_of_birth, &cas.Age, &cas.Gender, &cas.Job, &cas.Social_situation,
		&cas.Address_from_national_id_card, &cas.Actual_address, &cas.District, &cas.PhoneNumbers, &cas.Subsidies_id,
		&cas.Social_status, &cas.Husband_id, &cas.Created_at, &cas.Updated_at, &cas.Deleted, &cas.Date_Of_Social_situation, &cas.Case_entry_date, &cas.Status_search_update_date, &cas.Field_research_history, &cas.Living_expenses); err != nil {
		fmt.Println(err.Error(), "here")
		return Cases{}, fmt.Errorf("error: %v", err)
	}

	return cas, nil
}

func (ca Cases) GetAll(db *sql.DB, limit, offset, from, to int, district string) ([]Cases, error) {
	cases := []Cases{}
	query := "SELECT * FROM `cases` WHERE deleted = 0"
	inputs := []interface{}{}

	if from != 0 && to != 0 {
		query += " AND age > ? AND age <= ?"
		inputs = append(inputs, from, to)
	}

	if district != "" {
		query += " AND district = ?"
		inputs = append(inputs, district)
	}

	query += " LIMIT ? OFFSET ?"
	inputs = append(inputs, limit, offset)

	rows, err := db.Query(query, inputs...)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("error: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var cas Cases

		if err := scanQueryCases(rows, &cas); err != nil {
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

func (ca Cases) FilterCasesByRelativeAge(db *sql.DB, district string, from, to, limit, offset int) ([]Cases, error) {
	cases := []Cases{}
	sql := `SELECT DISTINCT cases.id, cases.case_name, cases.national_id, cases.devices_needed_for_the_case, cases.total_income,
		cases.fixed_expenses, cases.pension_from_husband, cases.pension_from_father, cases.debts, cases.case_type,
		cases.date_of_birth, cases.age, cases.gender, cases.job, cases.social_situation, cases.address_from_national_id_card,
		cases.actual_address, cases.district, cases.phone_numbers, cases.subsidies_id, cases.social_status, cases.husband_id,
		cases.created_at, cases.updated_at, cases.deleted, cases.date_of_social_situation, cases.case_entry_date,
		cases.status_search_update_date, cases.field_research_history, cases.living_expenses
		FROM cases INNER JOIN relative ON cases.id = relative.case_id
		WHERE cases.deleted = 0`

	var inputs []interface{}

	if from != 0 && to != 0 {
		sql += ` AND relative.age > ? AND relative.age <= ?`
		inputs = append(inputs, from, to)
	}

	if district != "" {
		sql += ` AND cases.district = ?`
		inputs = append(inputs, district)
	}

	sql += ` LIMIT ? OFFSET ?`
	inputs = append(inputs, limit, offset)
	rows, err := db.Query(sql, inputs...)

	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("error: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var cas Cases

		if err := scanQueryCases(rows, &cas); err != nil {
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

	query := "SELECT COUNT(*) AS length FROM `cases` WHERE deleted = 0"
	inputs := []interface{}{}

	if district != "" {
		query += " AND district = ?"
		inputs = append(inputs, district)
	}

	if from != 0 && to != 0 {
		query += " AND age > ? AND age <= ?"
		inputs = append(inputs, from, to)
	}

	row := db.QueryRow(query, inputs...)

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

		if err := scanQueryCases(rows, &cas); err != nil {
			fmt.Println(err)
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

		if err := scanQueryCases(rows, &cas); err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("error: %v", err)
		}

		cases = append(cases, cas)
	}

	return cases, nil
}

func scanQueryCases(rows *sql.Rows, cas *Cases) error {
	return rows.Scan(&cas.Id, &cas.Case_name, &cas.National_id, &cas.Devices_needed_for_the_case, &cas.Total_income,
		&cas.Fixed_expenses, &cas.Pension_from_husband, &cas.Pension_from_father, &cas.Debts,
		&cas.Case_type, &cas.Date_of_birth, &cas.Age, &cas.Gender, &cas.Job, &cas.Social_situation,
		&cas.Address_from_national_id_card, &cas.Actual_address, &cas.District, &cas.PhoneNumbers, &cas.Subsidies_id,
		&cas.Social_status, &cas.Husband_id, &cas.Created_at, &cas.Updated_at, &cas.Deleted, &cas.Date_Of_Social_situation, &cas.Case_entry_date, &cas.Status_search_update_date, &cas.Field_research_history, &cas.Living_expenses)
}
