package handlers

import (
	"archive/zip"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func DownloadFiles(w http.ResponseWriter, r *http.Request) {
	files := "./uploads/" + r.PathValue("id")

	w.Header().Set("Content-Disposition", "attachment; filename=files.zip")
	w.Header().Set("Content-Type", "application/zip")

	zipWriter := zip.NewWriter(w)
	defer zipWriter.Close()

	err := filepath.Walk(files, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			fmt.Println(err)
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			fmt.Println(err)
			return err
		}
		defer file.Close()

		relPath, err := filepath.Rel(files, path)
		if err != nil {
			fmt.Println(err)
			return err
		}

		wr, err := zipWriter.Create(relPath)
		if err != nil {
			fmt.Println(err)
			return err
		}

		_, err = wr.Write([]byte(file.Name()))
		fmt.Println(err)

		return err
	})

	if err != nil {
		http.Error(w, "Failed to create ZIP", http.StatusInternalServerError)
		return
	}
}
