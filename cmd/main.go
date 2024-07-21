package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	components "github.com/timeitel/groceries/internal/components/layout"
)

func main() {
	component := components.Layout("hi")

	http.Handle("/", templ.Handler(component))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
