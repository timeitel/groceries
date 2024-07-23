package routes

import (
	"fmt"
	"net/http"
)

func PostItemHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	name := r.Form.Get("item")

	// Access form data
	// name := r.Form.Get("name")
	// age := r.Form.Get("age")
	fmt.Println("post item", name)
}
