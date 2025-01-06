package middleware

import (
	"abir-el-hamd/internal/models"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(user models.Users) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.Id,
		"username":  user.Name,
		"login_at":  user.LoginAt,
		"role":   user.Role,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString([]byte("AbirElhamd"))
}

func ValidateToken(tokenStr string) (models.Users, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("AbirElhamd"), nil
	})
	if err != nil {
		return models.Users{}, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return models.Users{}, err
	}

	userId := claims["id"].(string)
	username := claims["username"].(string)
	loginAt, err := time.Parse(time.RFC3339, fmt.Sprint(claims["login_at"]))
	if err != nil {
		return models.Users{}, err
	}
	role := claims["role"].(string)

	return models.Users{
		Id:   userId,
		Name: username,
		LoginAt: loginAt,
		Role: role,
	}, nil
}

func Valid(BearerToken string) (bool, error) {
	_, err := ValidateToken(BearerToken)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	return true, nil
}

func GetToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("missing Authorization header")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", fmt.Errorf("invalid Authorization header format")
	}

	BearerToken := parts[1]

	return BearerToken, nil
}

func VerifyAdmin(w http.ResponseWriter, r *http.Request) {
	token, err := GetToken(r)

	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	ok, err := Valid(token)

	if !ok {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
}