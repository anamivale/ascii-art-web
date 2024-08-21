package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)


func StatusBadRequest(w http.ResponseWriter){
	w.WriteHeader(http.StatusBadRequest)
	if _, err := os.Stat("template/400.html"); os.IsNotExist(err) {
		StatusNotFound(w)
		return
	}
		tmp, err := template.ParseFiles("template/400.html")
		if err != nil {
			http.Error(w, "Internal Server Error1", http.StatusInternalServerError)

		}
		tmp.Execute(w, nil)
		
}
func StatusMethodNotAllowed(w http.ResponseWriter){
	w.WriteHeader(http.StatusMethodNotAllowed)
	if _, err := os.Stat("template/405.html"); os.IsNotExist(err) {
		StatusNotFound(w)
		return
	}
		tmp, err := template.ParseFiles("template/405.html")
		if err != nil {
			http.Error(w, "Internal Server Error1", http.StatusInternalServerError)

		}
		tmp.Execute(w, nil)
}
func StatusNotFound(w http.ResponseWriter){
	w.WriteHeader(http.StatusNotFound)
	if _, err := os.Stat("template/404.html"); os.IsNotExist(err) {
		StatusNotFound(w)
		return
	}
		tmp, err := template.ParseFiles("template/404.html")
		if err != nil {
			http.Error(w, "Internal Server Error1", http.StatusInternalServerError)

		}
		tmp.Execute(w, nil)
		
}

func StatusInternalServerError(w http.ResponseWriter){
	w.WriteHeader(http.StatusInternalServerError)
	if _, err := os.Stat("template/500.html"); os.IsNotExist(err) {
		StatusNotFound(w)
		return
	}
		tmp, err := template.ParseFiles("template/500.html")
		if err != nil {
			http.Error(w, "Internal Server Error1", http.StatusInternalServerError)

		}
		tmp.Execute(w, nil)
		
}

func ErrorHandlingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Unhandled error: %v", err)
				StatusInternalServerError(w)
			}
		}()
		next.ServeHTTP(w, r)
	})
}