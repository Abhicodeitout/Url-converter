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
	http.ListenAndServe("":3030", nil)
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
