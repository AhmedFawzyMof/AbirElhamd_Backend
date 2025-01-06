package models

import (
	"database/sql"
	"time"
)

type Relatives struct {
	Id                 int    `json:"id"`
	Type               string `json:"type"`
	Name               string `json:"name"`
	Date_of_birth      string `json:"date_of_birth"`
	Health_status      string `json:"health_status"`
	Education          string `json:"education"`
	National_id        string `json:"national_id"`
	Age                int    `json:"age"`
	Gender             string `json:"gender"`
	Job                string `json:"job"`
	Social_situation   string `json:"social_situation"`
}

func (ra Relatives) Add(db *sql.DB, id string) error {
	_, err := db.Exec("INSERT INTO relative (relative_type, name, date_of_birth, health_status, education, national_id, age, gender, job, social_situation, case_id) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", ra.Type, ra.Name, ra.Date_of_birth, ra.Health_status, ra.Education, ra.National_id, ra.Age, ra.Gender, ra.Job, ra.Social_situation, id)

	if err != nil {
		return err
	}

	_, err = db.Exec("UPDATE cases SET updated_at = ? WHERE id = ?", time.Now(), id)

	if err != nil {
		return err
	}

	return nil
}

func (ra Relatives) UPDATE(db *sql.DB, id string) error {
	_, err := db.Exec("UPDATE relative SET relative_type = ?, name = ?, national_id = ?, date_of_birth = ?, age = ?, gender = ? WHERE id = ?", ra.Type, ra.Name, ra.National_id, ra.Date_of_birth, ra.Age, ra.Gender, ra.Id)

	if err != nil {
		return err
	}

	_, err = db.Exec("UPDATE cases SET updated_at = ? WHERE id = ?", time.Now(), id)

	if err != nil {
		return err
	}

	return nil
}

func (ra Relatives) DELETE(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM relative WHERE id = ?", ra.Id)
	if err != nil {
		return err
	}
	return nil
}
