package handlers

import (
	"abir-el-hamd/internal/config"
	"abir-el-hamd/internal/middleware"
	"abir-el-hamd/internal/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func UploadCaseFiles(w http.ResponseWriter, r *http.Request) {
	token, err := middleware.GetToken(r)

	if err != nil {
		fmt.Println(err)
		middleware.ErrorResopnse(w, err)
		return
	}

	userData, err := middleware.ValidateToken(token)

	if err != nil {
		fmt.Println(err)
		middleware.ErrorResopnse(w, err)
		return
	}
	id := r.PathValue("id")

	database := config.Database()
	defer database.Close()


	err = r.ParseMultipartForm(30 << 20)

	if err != nil {
		fmt.Println(err)
		middleware.ErrorResopnse(w, err)
		return
	}

	path := fmt.Sprintf("./uploads/%s",id)

	fmt.Println(path)

	err = os.MkdirAll(path, 0755)

	if err != nil {
		fmt.Println(err)
		middleware.ErrorResopnse(w, err)
		return
	}

	files := r.MultipartForm.File["files"]

	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			fmt.Println(err)
			middleware.ErrorResopnse(w, err)
			return
		}
		defer file.Close()

		dst, err := os.Create("./uploads/" + id + "/" + fileHeader.Filename)
		if err != nil {
			fmt.Println(err)
			middleware.ErrorResopnse(w, err)
			return
		}
		defer dst.Close()

		if _, err := io.Copy(dst, file); err != nil {
			fmt.Println(err)
			middleware.ErrorResopnse(w, err)
			return
		}
	}

	if err := models.CreateLogs(database, id, "إنشاء حالة جديدة", userData.Id); err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}

	Response := map[string]interface{}{
		"success": "تم إنشاء حالة جديدة بنجاح",
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(Response); err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}
		
}
