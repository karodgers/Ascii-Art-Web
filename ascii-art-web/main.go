package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"ascii/handlers"
)

// PageData is a struct to hold data for rendering templates
type PageData struct {
	Result string
}

func main() {
	// Define route handlers
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiArtHandler)

	// Start the server
	fmt.Println("Starting server at :8080")
	fmt.Println("Server running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server failed to start:", err)
	}
}

// homeHandler handles requests to the root URL ("/")
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Serve 404 error if the path is not "/"
	if r.URL.Path != "/" {
		http.Error(w, "404 - not found", http.StatusNotFound)
		return
	}

	// Parse the template file
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute the template with no data
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// asciiArtHandler handles requests to the "/ascii-art" URL
func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	// Serve error if the method is not POST
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Get form values for text and banner options
	text := r.FormValue("message")
	banner := r.FormValue("banner-options") + ".txt"

	// Validate the inputs
	if text == "" || banner == "" {
		http.Error(w, "400: Please provide both text and banner", http.StatusBadRequest)
		return
	}

	// Check for non-ASCII characters in the text
	if err := handlers.ContainsNonASCII(text); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if the banner file exists
	if _, err := os.Stat(banner); os.IsNotExist(err) {
		http.Error(w, "500: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Validate the banner file checksum
	if err := handlers.ValidateFileChecksum(banner); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Generate ASCII art using the provided text and banner
	output := handlers.FunctionMain(text, banner)
	data := PageData{
		Result: string(output),
	}

	// Parse the template file
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute the template with the generated ASCII art
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
