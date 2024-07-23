package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/a-h/templ"
	views "github.com/timeitel/groceries/internal/views/layout"
)

func main() {
	fs := http.FileServer(http.Dir("../static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.Handle("/", templ.Handler(views.Layout()))

	fmt.Println("Running on port 3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
