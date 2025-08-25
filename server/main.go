package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	fmt.Println("Server running on http://localhost:8084")

	// Главная страница (index.html)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("../client/index.html"))
		tmpl.Execute(w, nil)
	})

	// Раздача статических файлов (CSS, JS)
	fs := http.FileServer(http.Dir("../client/assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8084", nil)
}
