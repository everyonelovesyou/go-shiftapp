package main

import (
	"html/template"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler)

	fileServer := http.FileServer(http.Dir("./public"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fileServer))

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	server.ListenAndServe()
}

func handler(w http.ResponseWriter, r *http.Request) {
	data := []string{"mon", "tue", "wed", "thu", "fri"}
	t := template.Must(template.ParseFiles("tpl/layout.html", "tpl/index.html", "tpl/component/day.html"))
	t.Execute(w, data)
}
