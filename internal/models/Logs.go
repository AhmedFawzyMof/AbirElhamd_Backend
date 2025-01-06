package models

import (
	"database/sql"
	"fmt"
)

type Logs struct {
	Id             int    `json:"id"`
	CaseId         string `json:"case_id"`
	Operation_type string `json:"operation_type"`
	UserId         string `json:"user_id"`
}

func CreateLogs(db *sql.DB, case_id, operation_type, user_id string) error {
	_, err := db.Exec("INSERT INTO logs (case_id, operation_type, user_id) VALUES (?, ?, ?)", case_id, operation_type, user_id)

	if err != nil {
		return err
	}

	return nil
}

func (l Logs) GetAll(db *sql.DB) ([]Logs, error) {
	var logs []Logs
	rows, err := db.Query("SELECT logs.id, logs.operation_type, Users.username as user_id, cases.case_name as case_id FROM logs LEFT JOIN Users ON Users.id = logs.user_id LEFT JOIN cases ON cases.id = logs.case_id ")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	for rows.Next() {
		log := Logs{}
		if err := rows.Scan(&log.Id, &log.Operation_type, &log.UserId, &log.CaseId); err != nil {
			fmt.Println(err)
			return nil, err
		}
		logs = append(logs, log)
	}

	return logs, nil

}
