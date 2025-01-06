package models

import (
	"database/sql"
	"fmt"
	"time"
)

type SS struct {
	Id                            int    `json:"id"`
	Properties                    string `json:"properties"`
	Health_status                 string `json:"health_status"`
	Education                     string `json:"education"`
	Number_of_family_members      int    `json:"number_of_family_members"`
	Number_of_registered_children int    `json:"number_of_registered_children"`
	Total_number_of_children      int    `json:"total_number_of_children"`
}

func (s SS) Add(db *sql.DB, id string) error {
	_, err := db.Exec(`INSERT INTO socialstatusofthecase (properties, health_status, education, number_of_family_members, number_of_registered_children, total_number_of_children) VALUES (?, ?, ?, ?, ?, ?)`,
		s.Properties, s.Health_status, s.Education, s.Number_of_family_members, s.Number_of_registered_children, s.Total_number_of_children)
	if err != nil {
		return err
	}

	stmt, err := db.Prepare("SELECT DISTINCT last_insert_rowid() FROM socialstatusofthecase")

	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow().Scan(&s.Id)

	if err != nil {
		return err
	}

	_, err = db.Exec("UPDATE cases SET updated_at = ?, social_status = ? WHERE id = ?", time.Now(), s.Id, id)

	if err != nil {
		return err
	}
	return nil
}

func (s SS) UPDATE(db *sql.DB, id string) error {
	_, err := db.Exec(`UPDATE socialstatusofthecase SET properties = ?, health_status = ?, education = ?, number_of_family_members = ?, number_of_registered_children = ?, total_number_of_children = ? WHERE id = ?`, s.Properties, s.Health_status, s.Education, s.Number_of_family_members, s.Number_of_registered_children, s.Total_number_of_children, s.Id)

	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = db.Exec("UPDATE cases SET updated_at = ? WHERE id = ?", time.Now(), id)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (s SS) DELETE(db *sql.DB, id string) error {
	_, err := db.Exec("DELETE FROM socialstatusofthecase WHERE id = ?", id)

	if err != nil {
		return err
	}
	return nil
}
