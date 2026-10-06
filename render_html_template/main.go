package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const tmplAddr = "render_html_template/tmp.html"

type User struct {
	Name  string
	Email string
}

type Server struct {
	User      User
	Router    *http.ServeMux
	Srvr      *http.Server
	Templates *template.Template
}

func NewServer() *Server {

	tmpl, err := template.ParseFiles(tmplAddr)

	if err != nil {
		log.Fatalf("Error parsing template file: %v", err)
	}

	u := User{
		Name:  "Eric",
		Email: "eric.ekholm@whatsurboyupto.com",
	}

	return &Server{
		User:   u,
		Router: http.NewServeMux(),
		Srvr: &http.Server{
			Addr: ":8080",
		},
		Templates: tmpl,
	}
}

func (s *Server) handleIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	if err := s.Templates.Execute(w, s.User); err != nil {
		log.Printf("Template execution error: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

}

func main() {
	s := NewServer()

	s.Router.HandleFunc("GET /", s.handleIndex)

	s.Srvr.Handler = s.Router

	fmt.Printf("Starting server on %s\n", s.Srvr.Addr)

	if err := s.Srvr.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Listen and serve error: %v", err)
	}
}
