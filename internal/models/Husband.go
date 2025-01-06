package models

import (
	"database/sql"
	"fmt"
	"time"
)

type Husband struct {
	Id                    int    `json:"id"`
	Husband_name          string `json:"Husband_name"`
	Husband_national_id   string `json:"Husband_national_id"`
	Husband_date_of_birth string `json:"Husband_date_of_birth"`
	Husband_age           int    `json:"Husband_age"`
	Husband_gender        string `json:"Husband_gender"`
}

func (h Husband) Add(db *sql.DB, id string) error {
	_, err := db.Exec("INSERT INTO `husband`(`name`, `national_id`, `date_of_birth`, `age`, `gender`) VALUES (?,?,?,?,?)", h.Husband_name, h.Husband_national_id, h.Husband_date_of_birth, h.Husband_age, h.Husband_gender)

	if err != nil {
		return err
	}

	stmt, err := db.Prepare("SELECT DISTINCT last_insert_rowid() FROM husband")

	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow().Scan(&h.Id)

	if err != nil {
		return err
	}

	_, err = db.Exec("UPDATE cases SET updated_at = ?, husband_id = ? WHERE id = ?", time.Now(), h.Id, id)

	if err != nil {
		return err
	}

	return nil
}

func (h Husband) UPDATE(db *sql.DB, id string) error {
	_, err := db.Exec("UPDATE husband SET name = ?, national_id = ?, date_of_birth = ?, age = ?, gender = ? WHERE id = ?", h.Husband_name, h.Husband_national_id, h.Husband_date_of_birth, h.Husband_age, h.Husband_gender, h.Id)

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

func (h Husband) DELETE(db *sql.DB, id string) error {
	_, err := db.Exec("DELETE FROM husband WHERE id = ?", id)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
