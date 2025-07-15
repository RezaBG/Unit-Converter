package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
	"unit-converter/converters"
)

var tempTmpl = template.Must(template.ParseFiles("templates/temperature.html"))

func TemperatureHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tempTmpl.Execute(w, nil)
		return
	}

	valueStr := r.FormValue("value")
	from := r.FormValue("from")
	to := r.FormValue("to")

	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		tempTmpl.Execute(w, PageData{Result: "Invalid input!"})
		return
	}

	convertedValue, err := converters.ConvertTemperature(value, from, to)
	if err != nil {
		tempTmpl.Execute(w, PageData{Result: "Error: " + err.Error()})
		return
	}

	result := fmt.Sprintf("%.2f %s = %.2f %s", value, from, convertedValue, to)
	tempTmpl.Execute(w, PageData{Result: result})
}
