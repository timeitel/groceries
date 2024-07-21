package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	user "github.com/timeitel/groceries/internal/components"
)

func main() {
	component := user.Hello("Testie")

	http.Handle("/", templ.Handler(component))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
