package models

type PaymentOrder struct {
	Sender     string  `xml:"Payer" json:"payer"`
	SenderAcct string  `xml:"PayerAcct" json:"payerAcct"`
	Rcvr       string  `xml:"Payee" json:"payee"`
	RcvrAcct   string  `xml:"PayeeAcct" json:"payeeAcct"`
	Amount     float64 `xml:"Amount" json:"amount"`
	Status     string  `xml:"Status" json:"status"`
	PoNumber   string  `json:"poNumber"`
}
