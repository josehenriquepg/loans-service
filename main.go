package main

type CustomerRequest struct {
	Age      int     `json:"age"`
	CPF      string  `json:"cpf"`
	Name     string  `json:"name"`
	Income   float64 `json:"income"`
	Location string  `json:"location"`
}

type LoanDetail struct {
	Type         string  `json:"type"`
	InterestRate float64 `json:"interest_rate"`
}

type LoanResponse struct {
	Customer string       `json:"customer"`
	Loans    []LoanDetail `json:"loans"`
}
