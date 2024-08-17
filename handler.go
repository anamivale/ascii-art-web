package main

import (
	"ascii/art"
	"html/template"
	"log"
	"net/http"
	"os"
)
func Handler(w http.ResponseWriter, r *http.Request) {
	// Check if the requested URL path is "/"
	if !(r.URL.Path == "/" || r.URL.Path == "/about"){
		StatusNotFound(w)
		return
	}
	var goData Pagedata // Create an empty Pagedata struct to hold the form data and output
	if r.Method == http.MethodPost { // Handle form submission
		r.ParseForm() // Parse the form data
		// Retrieve input text and banner values from the form
		input := r.FormValue("inputText")
		banner := r.FormValue("banner")
		// Generate ASCII art based on input text and banner
		output := art.AsciiArt(input, banner)
		// Populate the Pagedata struct with the input and output
		goData = Pagedata{
			Output:    output,
			InputText: input,
			Banner:    banner,
		}
		// Handle specific error conditions based on the output
		if input == "" || banner == "" || output == "internal server error" {
			StatusBadRequest(w)
			return
		} else if output == "Not Found" {
			StatusNotFound(w)
			return
		}
	}
	// Check if the HTML template file exists
	if _, err := os.Stat("template/index.html"); os.IsNotExist(err) {
		StatusNotFound(w)
		return
	}
	// Parse the HTML template file
	tmplt, err := template.ParseFiles("template/index.html")
	if err != nil {
		StatusInternalServerError(w)
		return
	}
	// Execute the template with the Pagedata struct
	err = tmplt.Execute(w, goData)
	if err != nil {
		log.Printf("Template execution error: %v", err)
		StatusInternalServerError(w)
		return
	}
}
func AboutHandler(w http.ResponseWriter, r *http.Request){
	if _, err := os.Stat("template/about.html"); os.IsNotExist(err) {
		StatusNotFound(w)
		return
	}
	tmplt, err := template.ParseFiles("template/about.html")
	if err != nil {
		StatusInternalServerError(w)
		return
	}
	err = tmplt.Execute(w, nil)
	if err != nil {
		log.Printf("Template execution error: %v", err)
		StatusInternalServerError(w)
		return
	}
}