package models

import (
	"database/sql"
	"fmt"
	"time"
)

type SS struct {
	Id                            sql.NullInt64  `json:"id"`
	Properties                    sql.NullString `json:"properties"`
	Health_status                 sql.NullString `json:"health_status"`
	Education                     sql.NullString `json:"education"`
	Number_of_family_members      sql.NullInt64  `json:"number_of_family_members"`
	Number_of_registered_children sql.NullInt64  `json:"number_of_registered_children"`
	Total_number_of_children      sql.NullInt64  `json:"total_number_of_children"`
	Case_id                       sql.NullString `json:"case_id"`
}

func (s SS) Add(db *sql.DB) error {
	_, err := db.Exec(`INSERT INTO socialstatusofthecase (properties, health_status, education, number_of_family_members, number_of_registered_children, total_number_of_children, case_id) VALUES (?, ?, ?, ?, ?, ?, ?)`,
		s.Properties, s.Health_status, s.Education, s.Number_of_family_members, s.Number_of_registered_children, s.Total_number_of_children, s.Case_id)

	if err != nil {
		fmt.Println("ss", err)
		return err
	}

	_, err = db.Exec("UPDATE cases SET updated_at = ?, social_status = ? WHERE id = ?", time.Now(), s.Id, s.Case_id)

	if err != nil {
		fmt.Println("ss", err)
		return err
	}
	return nil
}

func (s SS) UPDATE(db *sql.DB) error {
	_, err := db.Exec(`UPDATE socialstatusofthecase SET properties = ?, health_status = ?, education = ?, number_of_family_members = ?, number_of_registered_children = ?, total_number_of_children = ? WHERE id = ?`, s.Properties, s.Health_status, s.Education, s.Number_of_family_members, s.Number_of_registered_children, s.Total_number_of_children, s.Id)

	if err != nil {
		fmt.Println("ss", err)
		return err
	}

	_, err = db.Exec("UPDATE cases SET updated_at = ? WHERE id = ?", time.Now(), s.Case_id)

	if err != nil {
		fmt.Println("ss", err)
		return err
	}

	return nil
}

func (s SS) DELETE(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM socialstatusofthecase WHERE id = ?", s.Id)

	if err != nil {
		fmt.Println("ss", err)
		return err
	}
	return nil
}

func (s SS) GetByCaseID(db *sql.DB) (SS, error) {
	ss := SS{}
	row := db.QueryRow("SELECT * FROM `socialstatusofthecase` WHERE case_id = ?", s.Case_id)

	if err := row.Scan(&ss.Id, &ss.Properties, &ss.Health_status, &ss.Education, &ss.Number_of_family_members, &ss.Number_of_registered_children, &ss.Total_number_of_children, &ss.Case_id); err != nil {
		fmt.Println("ss", err)
		return ss, err
	}
	return ss, nil
}
