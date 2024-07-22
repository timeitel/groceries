package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/a-h/templ"
	components "github.com/timeitel/groceries/internal/components/layout"
)

func main() {
	fs := http.FileServer(http.Dir("../static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	component := components.Layout("hi")
	http.Handle("/", templ.Handler(component))

	fmt.Println("Running on port 3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
