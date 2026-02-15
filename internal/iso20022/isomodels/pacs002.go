package isomodels

type Pacs002 struct {
	FIToFIPmtStsRpt FIToFIPmtStsRpt `xml:"FIToFIPmtStsRpt" binding:"required"`
}

// FIToFIPaymentStatusReportV15 is the main pacs.002 message
type FIToFIPmtStsRpt struct {
	GrpHdr            GrpHdrPacs002       `xml:"GrpHdr" binding:"required"`
	OrgnlGrpInfAndSts []OrgnlGrpInfAndSts `xml:"OrgnlGrpInfAndSts,omitempty"`
	TxInfAndSts       []TxInfAndSts       `xml:"TxInfAndSts,omitempty"`
	SplmtryData       []SplmtryData       `xml:"SplmtryData,omitempty"`
}

type GrpHdrPacs002 struct {
	MsgId       string      `xml:"MsgId" binding:"required,max=35"`
	CreDtTm     string      `xml:"CreDtTm" binding:"required"`
	InstgAgt    Agent       `xml:"InstgAgt,omitempty"`
	InstdAgt    Agent       `xml:"InstdAgt,omitempty"`
	OrgnlBizQry OrgnlBizQry `xml:"OrgnlBizQry,omitempty" binding:"omitempty"`
}

type OrgnlBizQry struct {
	MsgId   string `xml:"MsgId" binding:"required,max=35"`
	MsgNmId string `xml:"MsgNmId" binding:"max=35"`
	CreDtTm string `xml:"CreDtTm"`
}

type OrgnlGrpInfAndSts struct {
	OrgnlMsgId    string          `xml:"OrgnlMsgId" binding:"required,max=35"`
	OrgnlMsgNmId  string          `xml:"OrgnlMsgNmId" binding:"required,max=35"`
	OrgnlCreDtTm  string          `xml:"OrgnlCreDtTm,omitempty"`
	OrgnlNbOfTxs  string          `xml:"OrgnlNbOfTxs,omitempty"`
	OrgnlCtrlSum  float64         `xml:"OrgnlCtrlSum,omitempty"`
	GrpSts        string          `xml:"GrpSts,omitempty"`
	StsRsnInf     []StsRsnInf     `xml:"StsRsnInf"`
	NbOfTxsPerSts []NbOfTxsPerSts `xml:"NbOfTxsPerSts"`
}

type TxInfAndSts struct {
	StsId           string     `xml:"StsId,omitempty"`
	OrgnlInstrId    string     `xml:"OrgnlInstrId,omitempty"`
	OrgnlEndToEndId string     `xml:"OrgnlEndToEndId,omitempty"`
	TxSts           string     `xml:"TxSts,omitempty"`
	StsRsnInf       StsRsnInf  `xml:"StsRsnInf,omitempty"`
	ChrgsInf        []ChrgsInf `xml:"ChrgsInf,omitempty"`
	AccptncDtTm     string     `xml:"AccptncDtTm,omitempty"`
	InstgAgt        Agent      `xml:"InstgAgt,omitempty"`
	InstdAgt        Agent      `xml:"InstdAgt,omitempty"`
}

type StsRsnInf struct {
	Orgtr    Party    `xml:"Orgtr,omitempty"`
	Rsn      Rsn      `xml:"Rsn,omitempty"`
	AddtlInf []string `xml:"AddtlInf,omitempty"`
}
type Rsn struct {
	Cd    string `xml:"Cd,omitempty" binding:"omitempty,min=1,max=4,required_without=Prtry"`
	Prtry string `xml:"Prtry,omitempty" binding:"max=35,required_without=Cd"`
}

type NbOfTxsPerSts struct {
	DtldNbOfTxs string  `xml:"DtldNbOfTxs" binding:"required,max=15"`
	DtldSts     string  `xml:"DtldSts" binding:"required,min=1,max=4"`
	DtldCtrlSum float64 `xml:"DtldCtrlSum,omitempty"`
}
