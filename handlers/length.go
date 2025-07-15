package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
	"unit-converter/converters"
)

type PageData struct {
	Result string
}

var lengthTmpl = template.Must(template.ParseFiles("templates/length.html"))

func LengthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		lengthTmpl.Execute(w, nil)
		return
	}

	valueStr := r.FormValue("value")
	from := r.FormValue("from")
	to := r.FormValue("to")

	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		lengthTmpl.Execute(w, PageData{Result: "Invalid input!"})
		return
	}

	convertedValue, err := converters.ConvertLength(value, from, to)
	if err != nil {
		lengthTmpl.Execute(w, PageData{Result: "Error: " + err.Error()})
		return
	}

	result := fmt.Sprintf("%.4f %s = %.4f %s", value, from, convertedValue, to)
	lengthTmpl.Execute(w, PageData{Result: result})

}
