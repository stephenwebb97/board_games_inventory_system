package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./Templates/layout.html", "./Templates/root.html"))
	tmpl.ExecuteTemplate(w, "root.html", nil)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	fmt.Fprintf(w, "Hi there, I love %s!", title)

}

func main() {
	http.Handle("/CSS/", http.FileServer(http.Dir("./")))
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/index.html", handler)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
