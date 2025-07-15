package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

var tmpl = template.Must(template.ParseFiles("templates/length.html"))

type PageData struct {
	Result string
}

func main() {
	http.HandleFunc("/", lengthHandler)

	fmt.Println("Server starting on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
	http.ListenAndServe(":8080", nil)
}

func lengthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl.Execute(w, nil)
		return
	}

	// POST request - handle from submission
	valueStr := r.FormValue("value")
	from := r.FormValue("from")
	to := r.FormValue("to")

	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		tmpl.Execute(w, PageData{Result: "Invalid input!"})
		return
	}

	// For now, just echo without conversion logic
	result := strconv.FormatFloat(value, 'f', -1, 64) + " " + from + " = " + to

	tmpl.Execute(w, PageData{Result: result})

}
