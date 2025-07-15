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
	http.HandleFunc("/weight", weightHandler)
	http.HandleFunc("/temperature", temperatureHandler)

	fmt.Println("Server starting on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
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

func convertWeight(value float64, from string, to string) (float64, error) {
	conversions := map[string]float64{
		"mg":    0.001,
		"g":     1.0,
		"kg":    1000.0,
		"ounce": 28.3495,
		"pound": 453.592,
	}

	fromFactor, ok1 := conversions[from]
	toFactor, ok2 := conversions[to]

	if !ok1 || !ok2 {
		return 0, fmt.Errorf("invalid unit conversion: %s or %s", from, to)
	}

	grams := value * fromFactor
	converted := grams / toFactor

	return converted, nil
}

func convertTemperature(value float64, from string, to string) (float64, error) {
	if from == to {
		return value, nil
	}

	// Convert input to Celsius first
	var celsius float64
	switch from {
	case "C":
		celsius = value
	case "F":
		celsius = (value - 32) * 5 / 9
	case "K":
		celsius = value - 273.15
	default:
		return 0, fmt.Errorf("invalid temperature unit: %s", from)
	}

	// Convert Celsius to target unit
	switch to {
	case "C":
		return celsius, nil
	case "F":
		return celsius*9/5 + 32, nil
	case "K":
		return celsius + 273.15, nil
	default:
		return 0, fmt.Errorf("invalid temperature unit: %s", to)
	}

}

func weightHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/weight.html"))

	if r.Method == http.MethodGet {
		tmpl.Execute(w, nil)
		return
	}

	valueStr := r.FormValue("value")
	from := r.FormValue("from")
	to := r.FormValue("to")

	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		tmpl.Execute(w, PageData{Result: "Invalid input!"})
		return
	}

	convertedValue, err := convertWeight(value, from, to)
	if err != nil {
		tmpl.Execute(w, PageData{Result: "Error: " + err.Error()})
		return
	}

	result := fmt.Sprintf("%.2f %s = %.2f %s", value, from, convertedValue, to)
	tmpl.Execute(w, PageData{Result: result})
}

func temperatureHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/temperature.html"))

	if r.Method == http.MethodGet {
		tmpl.Execute(w, nil)
		return
	}

	valueStr := r.FormValue("value")
	from := r.FormValue("from")
	to := r.FormValue("to")

	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		tmpl.Execute(w, PageData{Result: "Invalid input!"})
		return
	}

	convertedValue, err := convertTemperature(value, from, to)
	if err != nil {
		tmpl.Execute(w, PageData{Result: "Error: " + err.Error()})
		return
	}

	result := fmt.Sprintf("%.2f %s = %.2f %s", value, from, convertedValue, to)
	tmpl.Execute(w, PageData{Result: result})
}
