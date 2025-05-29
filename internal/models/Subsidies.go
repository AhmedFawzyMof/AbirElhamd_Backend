package models

import (
	"database/sql"
	"fmt"
	"time"
)

type Subsidies struct {
	Id                                            sql.NullInt64  `json:"id"`
	Grants_from_outside_the_association           sql.NullString `json:"grants_from_outside_the_association"`
	Grants_from_outside_the_association_financial sql.NullString `json:"grants_from_outside_the_association_financial"`
	Grants_from_the_association_financial         sql.NullString `json:"grants_from_the_association_financial"`
	Grants_from_the_association_inKind            sql.NullString `json:"grants_from_the_association_inKind"`
	Total_Subsidies                               sql.NullInt64  `json:"total_Subsidies"`
	End_Of_Payment_Date                           sql.NullTime   `json:"end_of_payment_date"`
	Note                                          sql.NullString `json:"note"`
	Case_id                                       sql.NullString `json:"case_id"`
}

func (s Subsidies) Add(db *sql.DB) error {
	_, err := db.Exec("INSERT INTO subsidies (grants_from_outside_the_association, grants_from_outside_the_association_financial, grants_from_the_association_financial, grants_from_the_association_inKind, total_Subsidies, end_of_payment_date, note, case_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", s.Grants_from_outside_the_association, s.Grants_from_outside_the_association_financial, s.Grants_from_the_association_financial, s.Grants_from_the_association_inKind, s.Total_Subsidies, s.End_Of_Payment_Date, s.Note, s.Case_id)

	if err != nil {
		fmt.Println("subsidies", err)
		return err
	}

	_, err = db.Exec("UPDATE cases SET updated_at = ? WHERE id = ?", time.Now(), s.Case_id)

	if err != nil {
		fmt.Println("subsidies", err)
		return err
	}

	return nil
}

func (s Subsidies) UPDATE(db *sql.DB) error {
	_, err := db.Exec("UPDATE subsidies SET grants_from_outside_the_association=?, grants_from_outside_the_association_financial=?, grants_from_the_association_financial=?, grants_from_the_association_inKind=?, total_Subsidies=?, end_of_payment_date = ?, note = ?, case_id = ? WHERE id=?", s.Grants_from_outside_the_association, s.Grants_from_outside_the_association_financial, s.Grants_from_the_association_financial, s.Grants_from_the_association_inKind, s.Total_Subsidies, s.End_Of_Payment_Date, s.Note, s.Case_id, s.Id)

	if err != nil {
		fmt.Println("subsidies", err)
		return err
	}

	_, err = db.Exec("UPDATE cases SET updated_at = ? WHERE id = ?", time.Now(), s.Case_id)

	if err != nil {
		fmt.Println("subsidies", err)
		return err
	}

	return nil
}

func (s Subsidies) DELETE(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM subsidies WHERE id = ?", s.Id)

	if err != nil {
		fmt.Println("subsidies", err)
		return err
	}

	return nil
}

func (s Subsidies) GetByCaseID(db *sql.DB) ([]Subsidies, error) {
	subsidies := []Subsidies{}
	rows, err := db.Query("SELECT * FROM `subsidies` WHERE case_id = ?", s.Case_id)
	if err != nil {
		fmt.Println("subsidies", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var subsidy Subsidies
		if err := scanQuerySubsidies(rows, &subsidy); err != nil {
			fmt.Println("subsidies", err)
			return nil, err
		}
		subsidies = append(subsidies, subsidy)
	}
	return subsidies, nil
}

func scanQuerySubsidies(rows *sql.Rows, subsidy *Subsidies) error {
	return rows.Scan(&subsidy.Id, &subsidy.Grants_from_outside_the_association, &subsidy.Grants_from_outside_the_association_financial, &subsidy.Grants_from_the_association_financial, &subsidy.Grants_from_the_association_inKind, &subsidy.Total_Subsidies, &subsidy.End_Of_Payment_Date, &subsidy.Note, &subsidy.Case_id)
}
