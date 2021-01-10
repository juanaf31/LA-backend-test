package models

type Account struct{
	AccountNumber string `json:"account_number"`
	CustName string `json:"customer_name"`
	Balance int `json:"balance"`
}

type Transfer struct{
	Receiver string `json:"to_account_number"`
	Amount int `json:"amount"`
}

type Found struct {
	AccountNumber string
}

type Message struct {
	Msg string `json:"message"`
}