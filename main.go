package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

// Structure that represents the customer's request
type CustomerRequest struct {
	Age      int     `json:"age"`
	CPF      string  `json:"cpf"`
	Name     string  `json:"name"`
	Income   float64 `json:"income"`
	Location string  `json:"location"`
}

// Structure representing a loan option in the response
type LoanDetail struct {
	Type         string  `json:"type"`
	InterestRate float64 `json:"interest_rate"`
}

// Structure that represents the response to the customer
type LoanResponse struct {
	Customer string       `json:"customer"`
	Loans    []LoanDetail `json:"loans"`
}

// Function that determines the types of loans available
func getAvailableLoans(c CustomerRequest) []LoanDetail {
	var loans []LoanDetail

	// Personal loan
	if c.Income <= 3000 {
		loans = append(loans, LoanDetail{Type: "Pessoal", InterestRate: 4})
	} else if c.Income > 3000 && c.Income <= 5000 && c.Age < 30 && strings.EqualFold(c.Location, "SP") {
		loans = append(loans, LoanDetail{Type: "PERSONAL", InterestRate: 4})
	}

	// Consignment loan
	if c.Income >= 5000 {
		loans = append(loans, LoanDetail{Type: "CONSIGNMENT", InterestRate: 2})
	}

	// Guaranteed loan
	if c.Income <= 3000 {
		loans = append(loans, LoanDetail{Type: "GUARANTEED", InterestRate: 3})
	} else if c.Income > 3000 && c.Income <= 5000 && c.Age < 30 && strings.EqualFold(c.Location, "SP") {
		loans = append(loans, LoanDetail{Type: "GUARANTEED", InterestRate: 3})
	}

	return loans
}

// Handler HTTP to process the request and determine loan options
func loanOptionsHandler(w http.ResponseWriter, r *http.Request) {
	var c CustomerRequest

	// Decode JSON request body into structure CustomerRequest
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Determine available loan terms
	loans := getAvailableLoans(c)

	// Prepare the answer
	response := LoanResponse{
		Customer: c.Name,
		Loans:    loans,
	}

	// Encode the response in JSON format
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	// Register the HTTP handler
	http.HandleFunc("/customer-loans", loanOptionsHandler)

	// Start the HTTP server on port 8080
	http.ListenAndServe(":8080", nil)
}
