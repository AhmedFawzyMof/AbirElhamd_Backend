package models

import (
	"database/sql"
	"fmt"
	"time"
)

type Relatives struct {
	Id               sql.NullInt32  `json:"id"`
	Type             sql.NullString `json:"type"`
	Name             sql.NullString `json:"name"`
	Date_of_birth    sql.NullString `json:"date_of_birth"`
	Health_status    sql.NullString `json:"health_status"`
	Education        sql.NullString `json:"education"`
	National_id      sql.NullString `json:"national_id"`
	Age              sql.NullInt32  `json:"age"`
	Gender           sql.NullString `json:"gender"`
	Job              sql.NullString `json:"job"`
	Social_situation sql.NullString `json:"social_situation"`
	Case_id          sql.NullString `json:"case_id"`
}

func (ra Relatives) Add(db *sql.DB, id string) error {
	_, err := db.Exec("INSERT INTO relative (relative_type, name, date_of_birth, health_status, education, national_id, age, gender, job, social_situation, case_id) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", ra.Type, ra.Name, ra.Date_of_birth, ra.Health_status, ra.Education, ra.National_id, ra.Age, ra.Gender, ra.Job, ra.Social_situation, id)

	if err != nil {
		fmt.Println("relatives", err)
		return err
	}

	_, err = db.Exec("UPDATE cases SET updated_at = ? WHERE id = ?", time.Now(), id)

	if err != nil {
		fmt.Println("relatives", err)
		return err
	}

	return nil
}

func (ra Relatives) UPDATE(db *sql.DB, id string) error {
	_, err := db.Exec("UPDATE relative SET relative_type = ?, name = ?, national_id = ?, date_of_birth = ?, age = ?, gender = ? WHERE id = ?", ra.Type, ra.Name, ra.National_id, ra.Date_of_birth, ra.Age, ra.Gender, ra.Id)

	if err != nil {
		fmt.Println("relatives", err)
		return err
	}

	_, err = db.Exec("UPDATE cases SET updated_at = ? WHERE id = ?", time.Now(), id)

	if err != nil {
		fmt.Println("relatives", err)
		return err
	}

	return nil
}

func (ra Relatives) DELETE(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM relative WHERE id = ?", ra.Id)
	if err != nil {
		fmt.Println("relatives", err)
		return err
	}
	return nil
}

func (ra Relatives) GetAll(db *sql.DB, id string) ([]Relatives, error) {
	relatives := []Relatives{}
	rows, err := db.Query("SELECT * FROM `relative` WHERE case_id = ?", id)
	if err != nil {
		fmt.Println("relatives", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var rel Relatives
		if err := scanQueryRelative(rows, &rel); err != nil {
			fmt.Println("relatives", err)
			return nil, err
		}
		relatives = append(relatives, rel)
	}
	return relatives, nil
}

func (ra Relatives) GetByCasesIDS(db *sql.DB, ids []any) ([]Relatives, error) {
	relatives := []Relatives{}
	var sql string = "SELECT * FROM `relative` WHERE case_id IN ("

	for i := range ids {
		if i != len(ids)-1 {
			sql += "?,"
		}

		if i == len(ids)-1 {
			sql += "?)"
		}
	}

	sql += " GROUP BY case_id"

	rows, err := db.Query(sql, ids...)
	if err != nil {
		fmt.Println("relatives", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var rel Relatives
		if err := scanQueryRelative(rows, &rel); err != nil {
			fmt.Println("relatives", err)
			return nil, err
		}
		relatives = append(relatives, rel)
	}
	return relatives, nil
}

func (ra Relatives) GetByCaseID(db *sql.DB, id string) ([]Relatives, error) {
	relatives := []Relatives{}
	rows, err := db.Query("SELECT * FROM `relative` WHERE case_id = ?", id)
	if err != nil {
		fmt.Println("relatives", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var rel Relatives
		if err := scanQueryRelative(rows, &rel); err != nil {
			fmt.Println("relatives", err)
			return nil, err
		}
		relatives = append(relatives, rel)
	}
	return relatives, nil
}

func scanQueryRelative(rows *sql.Rows, rel *Relatives) error {
	return rows.Scan(&rel.Id, &rel.Type, &rel.Name, &rel.National_id, &rel.Date_of_birth, &rel.Age, &rel.Gender, &rel.Job, &rel.Social_situation, &rel.Health_status, &rel.Education, &rel.Case_id)
}
