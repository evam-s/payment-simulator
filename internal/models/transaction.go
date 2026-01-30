package models

type PaymentOrder struct {
	Sender     string  `xml:"Payer"`
	SenderAcct string  `xml:"PayerAcct"`
	Rcvr       string  `xml:"Payee"`
	RcvrAcct   string  `xml:"PayeeAcct"`
	Amount     float64 `xml:"Amount"`
	Status     string  `xml:"Status"`
}