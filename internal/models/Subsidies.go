package models

import (
	"database/sql"
	"fmt"
	"time"
)

type Subsidies struct {
	Id                                            int    `json:"id"`
	Grants_from_outside_the_association           string `json:"grants_from_outside_the_association"`
	Grants_from_outside_the_association_financial string `json:"grants_from_outside_the_association_financial"`
	Grants_from_the_association_financial         string `json:"grants_from_the_association_financial"`
	Grants_from_the_association_inKind            string `json:"grants_from_the_association_inKind"`
	Total_Subsidies                               int    `json:"total_Subsidies"`
	End_Of_Payment_Date                           string `json:"end_of_payment_date"`
	Note 										  string `json:"note"`
}

func (s Subsidies) Add(db *sql.DB, id string) error {
	_, err := db.Exec("INSERT INTO subsidies (grants_from_outside_the_association, grants_from_outside_the_association_financial, grants_from_the_association_financial, grants_from_the_association_inKind, total_Subsidies, end_of_payment_date, note) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", s.Grants_from_outside_the_association, s.Grants_from_outside_the_association_financial, s.Grants_from_the_association_financial, s.Grants_from_the_association_inKind, s.Total_Subsidies, s.End_Of_Payment_Date, s.Note)

	if err != nil {
		return err
	}

	stmt, err := db.Prepare("SELECT DISTINCT last_insert_rowid() FROM subsidies")

	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow().Scan(&s.Id)

	if err != nil {
		return err
	}

	_, err = db.Exec("UPDATE cases SET updated_at = ?, subsidies_id = ? WHERE id = ?", time.Now(), s.Id, id)

	if err != nil {
		return err
	}

	return nil
}

func (s Subsidies) UPDATE(db *sql.DB, id string) error {
	_, err := db.Exec("UPDATE subsidies SET grants_from_outside_the_association=?, grants_from_outside_the_association_financial=?, grants_from_the_association_financial=?, grants_from_the_association_inKind=?, total_Subsidies=?, end_of_payment_date = ?, note = ? WHERE id=?", s.Grants_from_outside_the_association, s.Grants_from_outside_the_association_financial, s.Grants_from_the_association_financial, s.Grants_from_the_association_inKind, s.Total_Subsidies, s.End_Of_Payment_Date, s.Note, s.Id)

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

func (s Subsidies) DELETE(db *sql.DB, id string) error {
	_, err := db.Exec("DELETE FROM subsidies WHERE id = ?", id)

	if err != nil {
		return err
	}

	return nil
}
