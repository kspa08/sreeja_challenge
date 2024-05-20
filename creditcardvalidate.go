package main

import (
	"fmt"
	"net/http"
	"regexp"
)

func isValidCreditCard(cardNumber string) bool {
	re := regexp.MustCompile(`^(?:4[0-9]{12}(?:-[0-9]{4})?|[56][0-9]{15})$`)
	return re.MatchString(cardNumber)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	cardNumber := r.FormValue("cardNumber")
	if cardNumber == "" {
		fmt.Fprintf(w, "Missing credit card number")
		return
	}

	if isValidCreditCard(cardNumber) {
		fmt.Fprintf(w, "Valid credit card number")
	} else {
		fmt.Fprintf(w, "Invalid credit card number")
	}
}

func main() {
	http.HandleFunc("/", handleRequest)
	fmt.Println("Starting web server on port 8080")
	http.ListenAndServe(":8080", nil)
}
