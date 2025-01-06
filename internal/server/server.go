package server

import (
	"abir-el-hamd/internal/handlers"
	"abir-el-hamd/internal/middleware"
	"net/http"
	"path/filepath"
)

type Server struct {
	listenAddr string
	staticPath string
}

func NewServer(listenAddr string, staticPath string) *Server {
	return &Server{
		listenAddr: listenAddr,
		staticPath: staticPath,
	}
}

func (s *Server) Start() error {
	mux := http.NewServeMux()

	mux.HandleFunc("/api", handlers.HomeApi)
	mux.HandleFunc("/api/deleted-cases", handlers.DeletedCases)
	mux.HandleFunc("/api/logs", handlers.Logs)
	mux.HandleFunc("/api/filter", handlers.FilterKids)
	mux.HandleFunc("/api/users", handlers.GetAllUsers)
	mux.HandleFunc("POST /api/cases/add", handlers.AddCase)
	mux.HandleFunc("PUT /api/cases/edit", handlers.UpdateCase)
	mux.HandleFunc("DELETE /api/case/{id}", handlers.DeleteCase)
	mux.HandleFunc("DELETE /api/users/{id}", handlers.DeleteUser)
	mux.HandleFunc("DELETE /api/relative/{id}", handlers.DeleteRelative)
	mux.HandleFunc("DELETE /api/subsidies/{id}", handlers.DeleteSubsidies)
	mux.HandleFunc("DELETE /api/ss/{id}", handlers.DeleteSS)
	mux.HandleFunc("DELETE /api/husband/{id}", handlers.DeleteHusband)
	mux.HandleFunc("POST /api/relative/add", handlers.AddRelative)
	mux.HandleFunc("POST /api/relative/edit", handlers.UpdateRelative)
	mux.HandleFunc("POST /api/subsidies/add", handlers.AddSubsidies)
	mux.HandleFunc("POST /api/subsidies/edit", handlers.UpdateSubsidies)
	mux.HandleFunc("POST /api/ss/add", handlers.AddSS)
	mux.HandleFunc("POST /api/ss/edit", handlers.UpdateSS)
	mux.HandleFunc("POST /api/husband/add", handlers.AddHusband)
	mux.HandleFunc("POST /api/husband/edit", handlers.UpdateHusband)
	mux.HandleFunc("POST /api/users/add", handlers.AddUser)
	mux.HandleFunc("POST /api/case/upload/{id}", handlers.UploadCaseFiles)
	mux.HandleFunc("GET /api/case/{id}", handlers.CaseApi)
	mux.HandleFunc("/api/search", handlers.SearchCase)
	mux.HandleFunc("POST /api/login", handlers.LoginApi)
	mux.HandleFunc("/api/check-token", handlers.CheckLogin)
	mux.HandleFunc("GET /api/download/{id}", handlers.DownloadFiles)

	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.ServeFile(w, r, filepath.Join(s.staticPath, "index.html"))
		return
	}	})
	return http.ListenAndServe(s.listenAddr, middleware.CorsMiddleware(mux))
}
