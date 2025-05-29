package models

import (
	"database/sql"
	"fmt"
	"time"
)

type Husband struct {
	Id            sql.NullInt32  `json:"id"`
	Name          sql.NullString `json:"name"`
	National_id   sql.NullString `json:"national_id"`
	Date_of_birth sql.NullString `json:"date_of_birth"`
	Age           sql.NullInt32  `json:"age"`
	Gender        sql.NullString `json:"gender"`
	Case_id       sql.NullString `json:"case_id"`
}

func (h Husband) Add(db *sql.DB) error {
	_, err := db.Exec("INSERT INTO husband (name, national_id, date_of_birth, age, gender, case_id) VALUES (?, ?, ?, ?, ?, ?)", h.Name, h.National_id, h.Date_of_birth, h.Age, h.Gender, h.Case_id)

	if err != nil {
		fmt.Println("husband", err)
		return err
	}
	_, err = db.Exec("UPDATE cases SET updated_at = ? WHERE id = ?", time.Now(), h.Case_id)

	if err != nil {
		fmt.Println("husband", err)
		return err
	}

	return nil
}

func (h Husband) UPDATE(db *sql.DB) error {
	_, err := db.Exec("UPDATE husband SET name = ?, national_id = ?, date_of_birth = ?, age = ?, gender = ? WHERE id = ?", h.Name, h.National_id, h.Date_of_birth, h.Age, h.Gender, h.Id)

	if err != nil {
		fmt.Println("husband", err)
		return err
	}

	_, err = db.Exec("UPDATE cases SET updated_at = ? WHERE id = ?", time.Now(), h.Case_id)

	if err != nil {
		fmt.Println("husband", err)
		return err
	}

	return nil
}

func (h Husband) DELETE(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM husband WHERE id = ?", h.Id)

	if err != nil {
		fmt.Println("husband", err)
		return err
	}

	return nil
}

func (h Husband) GetByCaseID(db *sql.DB) ([]Husband, error) {
	husbands := []Husband{}
	rows, err := db.Query("SELECT * FROM `husband` WHERE case_id = ?", h.Case_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var husband Husband
		if err := scanQueryHusband(rows, &husband); err != nil {
			fmt.Println("husband", err)
			return nil, err
		}
		husbands = append(husbands, husband)
	}
	return husbands, nil
}

func scanQueryHusband(rows *sql.Rows, h *Husband) error {
	return rows.Scan(&h.Id, &h.Name, &h.National_id, &h.Date_of_birth, &h.Age, &h.Gender, &h.Case_id)
}
