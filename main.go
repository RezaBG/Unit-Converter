package main

import (
	"fmt"
	"net/http"
	"unit-converter/handlers"
)

func main() {
	http.HandleFunc("/", handlers.LengthHandler)
	http.HandleFunc("/weight", handlers.WeightHandler)
	http.HandleFunc("/temperature", handlers.TemperatureHandler)

	fmt.Println("Server starting on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
