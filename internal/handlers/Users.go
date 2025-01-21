package handlers

import (
	"abir-el-hamd/internal/config"
	"abir-el-hamd/internal/middleware"
	"abir-el-hamd/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func LoginApi(w http.ResponseWriter, r *http.Request) {
	database := config.Database()
	defer database.Close()

	user := models.Users{}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}

	userData, err := user.Login(database)

	if err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}

	token, err := middleware.GenerateToken(userData)

	if err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	Response := map[string]interface{}{
		"user":  userData,
		"token": token,
	}
	if err := json.NewEncoder(w).Encode(Response); err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}
}

func CheckLogin(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		middleware.ErrorResopnse(w, fmt.Errorf("missing Authorization header"))
		return
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		middleware.ErrorResopnse(w, fmt.Errorf("invalid Authorization header format"))
		return
	}

	BearerToken := parts[1]

	userData, err := middleware.ValidateToken(BearerToken)
	if err != nil {
		fmt.Println(err)
		middleware.ErrorResopnse(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(userData); err != nil {
		fmt.Println(err)
		middleware.ErrorResopnse(w, err)
		return
	}
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	database := config.Database()
	defer database.Close()

	users, err := models.GetAllUsers(database)
	if err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(users); err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	database := config.Database()
	defer database.Close()

	userInterface := map[string]interface{}{}

	if err := json.NewDecoder(r.Body).Decode(&userInterface); err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}

	user := models.Users{}

	user.Name = userInterface["name"].(string)
	user.Password = userInterface["password"].(string)
	user.Role = userInterface["role"].(string)

	if err := models.AddUser(database, user); err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}

	Response := map[string]interface{}{
		"success": "تمت العملية بنجاح",
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(Response); err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	database := config.Database()
	defer database.Close()

	id := r.PathValue("id")
	if err := models.DeleteUser(database, id); err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}

	Response := map[string]interface{}{
		"success": "تمت العملية بنجاح",
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(Response); err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}
}
