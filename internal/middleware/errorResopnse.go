package middleware

import "net/http"

func ErrorResopnse(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
