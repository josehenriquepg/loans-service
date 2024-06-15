package main

import (
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
