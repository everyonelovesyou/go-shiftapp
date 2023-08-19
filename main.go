package main

import (
	"html/template"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	server.ListenAndServe()
}

func handler(w http.ResponseWriter, r *http.Request) {
	tpl := `<p>{{ . }}</p>`
	data := "Hello Kitty!"
	t := template.Must(template.New("hello").Parse(tpl))
	t.Execute(w, data)
}
