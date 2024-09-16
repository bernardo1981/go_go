package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// Handler para servir la página HTML
func serveHTML(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

// Handler para calcular el resultado
func calculate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	expression := r.FormValue("expression")

	// Aquí, el manejo de errores y la lógica de evaluación es muy simple
	// en una aplicación real, se debe usar una biblioteca para evaluar expresiones matemáticas de manera segura.
	result := evaluateExpression(expression)

	fmt.Fprintf(w, "%s", result)
}

// Evalúa la expresión matemática (muy básico y no seguro para uso real)
func evaluateExpression(expression string) string {
	// Esto solo maneja expresiones muy simples y no es seguro para uso real.
	// En una aplicación real, usar una biblioteca para la evaluación de expresiones.
	expression = strings.TrimSpace(expression)
	result, err := eval(expression)
	if err != nil {
		return "Error: " + err.Error()
	}
	return fmt.Sprintf("%f", result)
}

// Simple función para evaluar expresiones (solo para demostración)
func eval(expr string) (float64, error) {
	// Este código no evalúa expresiones matemáticas reales.
	// Para fines demostrativos, esto solo convierte la cadena a un número.
	return strconv.ParseFloat(expr, 64)
}

func main() {
	http.HandleFunc("/", serveHTML)
	http.HandleFunc("/calculate", calculate)

	fmt.Println("Servidor iniciado en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
