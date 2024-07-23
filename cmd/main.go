package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/timeitel/groceries/internal/routes"
	views "github.com/timeitel/groceries/internal/views/layout"
)

func main() {
	fs := http.FileServer(http.Dir("../static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.Handle("/", templ.Handler(views.Layout()))
	http.HandleFunc("/item", routes.PostItemHandler)

	fmt.Println("Running on port 3000")

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
