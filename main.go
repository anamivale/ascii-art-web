package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// Pagedata struct holds the data to be passed to the HTML template
type Pagedata struct {
	Output    string
	InputText string
	Banner    string
}

// handler function processes the form submission and renders the HTML template

func main() {
	if len(os.Args) != 1 {
		fmt.Println("Usage: go run .")
		return
	}
	// Serve static files from the "static"
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Handle the root URL with the error handling middleware and the main handler
	http.Handle("/ascii-art", ErrorHandlingMiddleware(http.HandlerFunc(Handler)))
	http.Handle("/", ErrorHandlingMiddleware(http.HandlerFunc(Home)))
	http.Handle("/download", ErrorHandlingMiddleware(http.HandlerFunc(Download)))

	http.Handle("/about", ErrorHandlingMiddleware(http.HandlerFunc(AboutHandler)))

	// Print a message indicating the server is running and listen on port 1024
	fmt.Println("running on http://localhost:3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
