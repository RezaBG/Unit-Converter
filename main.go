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

	// Perform conversion
	convertedLength, err := convertLength(value, from, to)
	if err != nil {
		tmpl.Execute(w, PageData{Result: "Error: " + err.Error()})
		return
	}

	tmpl.Execute(w, PageData{Result: fmt.Sprintf("%f %s = %f %s", value, from, convertedLength, to)})
}

func convertLength(value float64, from string, to string) (float64, error) {
	// Define conversion factors to meters
	conversions := map[string]float64{
		"mm":   0.001,
		"cm":   0.01,
		"m":    1.0,
		"km":   1000.0,
		"inch": 0.0254,
		"ft":   0.3048,
		"yard": 0.9144,
		"mile": 1609.34,
	}

	fromFactor, ok1 := conversions[from]
	toFactor, ok2 := conversions[to]

	if !ok1 || !ok2 {
		return 0, fmt.Errorf("invalid unit conversion: %s or %s", from, to)
	}

	// Convert input to meters
	meters := value * fromFactor
	// Convert meters to target unit
	converted := meters / toFactor

	return converted, nil
}
