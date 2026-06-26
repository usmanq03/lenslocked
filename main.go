package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl, err := template.ParseFiles("templates/home.gohtml")
	if err != nil {
		//panic(err) //TODO: Remove the panic
		//hanlde the error
		log.Printf("pasing template: %v", err)
		http.Error(w, "There was an error parsing the template.", http.StatusInternalServerError)
		return //return is important so it doesnt go below and execute tpl.execute
	}

	err = tpl.Execute(w, nil)
	if err != nil {
		//panic(err) //TODO: Remove the panic
		//handling the error below
		log.Printf("executing template: %v", err)
		http.Error(w, "There was an error executing the template.", http.StatusInternalServerError)
		return
	}
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h1>Contact Page</h1><p>To get in touch, email me at <a href="mailto:me@uqamar.co">me@uqamar.co</a>`)
}

// func pathHandler(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 	case "/contact":
// 		contactHandler(w, r)
// 	default:
// 		w.WriteHeader(404)
// 		fmt.Fprint(w, "<p>This page does not exist</p>")
// 	}
// }

func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	//http.HandleFunc("/", pathHandler)
	//http.HandleFunc("/contact", contactHandler)
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
