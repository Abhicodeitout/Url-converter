package main


import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var urlMap := make(map[string]string)

func main(){
	http.HandleFunc("/", HandleFrm)
	http.HandleFunc("/shorturl", HandleShort)
	http.HandleFunc("/short/", HandleRdirect)
	http.ListenAndServe(":8080", nil)
}

func handleFrm(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		http.Redirect(w, r, "/shorturl", http.StatusSeeOther)
		return
	}

	// Serve the HTML form
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, `
		<!DOCTYPE html>
		<html>
		<head>
			<title>URL Shortener</title>
		</head>
		<body>
			<h2>URL Shortener</h2>
			<form method="post" action="/shorturl">
				<input type="url" name="url" placeholder="Enter a URL" required>
				<input type="submit" value="Shorturl">
			</form>
		</body>
		</html>
	`)
}


func handleShort(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	originalURL := r.FormValue("url")
	if originalURL == "" {
		http.Error(w, "URL parameter is missing", http.StatusBadRequest)
		return
	}

	shortKey := generateShortKey()
	urlMap[shortKey] = originalURL

	shrtEndURL := fmt.Sprintf("http://localhost:8080/short/%s", shortKey)

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, `
		<!DOCTYPE html>
		<html>
		<head>
			<title>URL Shortener</title>
		</head>
		<body>
			<h2>URL Shortener</h2>
			<p>Original URL: `, originalURL, `</p>
			<p>Shortened URL: <a href="`, shrtEndURL, `">`, shrtEndURL, `</a></p>
		</body>
		</html>
	`)
}


func handleRdirect(w http.ResponseWriter, r *http.Request) {
	shortKey := strings.TrimPrefix(r.URL.Path, "/short/")
	if shortKey == "" {
		http.Error(w, "Shortened key missing", http.StatusBadRequest)
		return
	}

	originalURL, found := urlMap[shortKey]
	if !found {
		http.Error(w, "Shortened key not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusMovedPermanently)
}
