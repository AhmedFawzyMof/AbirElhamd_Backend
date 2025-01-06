package models

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Users struct {
	Id       string    `json:"id"`
	Name     string    `json:"username"`
	Password string    `json:"password"`
	Role     string    `json:"role"`
	LoginAt  time.Time `json:"login_at"`
}

func AddUser(db *sql.DB, user Users) error {
	user.Id = uuid.New().String()
	passwordHash := sha256.Sum256([]byte(user.Password))
	user.Password = hex.EncodeToString(passwordHash[:])

	_, err := db.Exec("INSERT INTO Users (id, username, password, role, login_at) VALUES (?, ?, ?, ?, ?)", user.Id, user.Name, user.Password, user.Role, user.LoginAt)

	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	return nil
}

func GetAllUsers(db *sql.DB) ([]Users, error) {
	rows, err := db.Query("SELECT * FROM Users")

	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("error: %v", err)
	}

	defer rows.Close()

	users := []Users{}

	for rows.Next() {
		var user Users
		if err := rows.Scan(&user.Id, &user.Name, &user.Password, &user.LoginAt, &user.Role); err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("error: %v", err)
		}
		users = append(users, user)
	}

	return users, nil
}

func DeleteUser(db *sql.DB, id string) error {
	_, err := db.Exec("DELETE FROM Users WHERE id = ?", id)

	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	return nil
}

func (u Users) Login(db *sql.DB) (Users, error) {
	var User Users

	passwordHash := sha256.Sum256([]byte(u.Password))
	u.Password = hex.EncodeToString(passwordHash[:])

	row := db.QueryRow("SELECT id, username, password, role FROM Users WHERE username = ? AND password = ?", u.Name, u.Password)

	if err := row.Scan(&User.Id, &User.Name, &User.Password, &User.Role); err != nil {
		return Users{}, fmt.Errorf("error: %v", err)
	}

	_, err := db.Exec("UPDATE Users SET Login_At = ? WHERE id = ?", time.Now(), User.Id)

	if err != nil {
		return Users{}, fmt.Errorf("error: %v", err)
	}
	return User, nil

}
