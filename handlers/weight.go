package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
	"unit-converter/converters"
)

var weightTmpl = template.Must(template.ParseFiles("templates/weight.html"))

func WeightHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		weightTmpl.Execute(w, nil)
		return
	}

	valueStr := r.FormValue("value")
	from := r.FormValue("from")
	to := r.FormValue("to")

	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		weightTmpl.Execute(w, PageData{Result: "Invalid input!"})
		return
	}

	convertedValue, err := converters.ConvertWeight(value, from, to)
	if err != nil {
		weightTmpl.Execute(w, PageData{Result: "Error: " + err.Error()})
		return
	}

	result := fmt.Sprintf("%.4f %s = %.4f %s", value, from, convertedValue, to)
	weightTmpl.Execute(w, PageData{Result: result})
}
