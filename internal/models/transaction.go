package models

type PaymentOrder struct {
	Debtor         string   `xml:"Payer" json:"payer"`
	DebtorAcct     string   `xml:"PayerAcct" json:"payerAcct"`
	Creditor       string   `xml:"Payee" json:"payee"`
	CreditorAcct   string   `xml:"PayeeAcct" json:"payeeAcct"`
	Amount         float64  `xml:"Amount" json:"amount"`
	AmountCurrency string   `xml:"Currency" json:"currency"`
	Status         string   `xml:"Status" json:"status"`
	EntityId       string   `json:"entityId"`
	Errors         []string `xml:"Errors" json:"errors"`
	InstructionId  string
	MsgId          string
}

// func ConvertPacs008ToPaymentOrder(isoPacs Pacs008) PaymentOrder {
// 	txn := isoPacs.FIToFICstmrCdtTrf.CdtTrfTxInf
// 	grpHdr := isoPacs.FIToFICstmrCdtTrf.GrpHdr

// 	return PaymentOrder{
// 		Debtor:       txn.Dbtr.Name,
// 		DebtorAcct:   txn.DbtrAcct.IBAN,
// 		Creditor:     txn.Cdtr.Name,
// 		CreditorAcct: txn.CdtrAcct.IBAN,

// 		Amount:         txn.Amt.Value,
// 		AmountCurrency: txn.Amt.Currency,

// 		InstructionId: txn.PmtId.InstrId,
// 		Status:        txn.Status,

// 		SttlmDt:    tx.SttlmDt,
// 		Remittance: tx.RmtInf.Ustrd,

// 		MsgId:    grpHdr.MsgId,
// 		CreDtTm:  grpHdr.CreDtTm,
// 		NbOfTxs:  grpHdr.NbOfTxs,
// 		SttlmMtd: grpHdr.SttlmMtd,
// 	}
// }
