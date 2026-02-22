package isomodels

import "encoding/xml"

type Pacs002 struct {
	XMLName         xml.Name        `xml:"Document"`
	Xmlns           string          `xml:"xmlns,attr"`
	SchemaLocation  string          `xml:"schemaLocation,attr"`
	XmlnsXsi        string          `xml:"xsi,attr"`
	FIToFIPmtStsRpt FIToFIPmtStsRpt `xml:"FIToFIPmtStsRpt" binding:"required"`
}

type FIToFIPmtStsRpt struct {
	GrpHdr            GrpHdrPacs002        `xml:"GrpHdr" binding:"required"`
	OrgnlGrpInfAndSts []*OrgnlGrpInfAndSts `xml:"OrgnlGrpInfAndSts,omitempty" binding:"dive"`
	TxInfAndSts       []*TxInfAndSts       `xml:"TxInfAndSts,omitempty" binding:"dive"`
	SplmtryData       []*SplmtryData       `xml:"SplmtryData,omitempty" binding:"omitempty,dive"`
}

type GrpHdrPacs002 struct {
	MsgId       string       `xml:"MsgId" binding:"required,max=35"`
	CreDtTm     string       `xml:"CreDtTm" binding:"required,isoDateTime"`
	InstgAgt    *Agent       `xml:"InstgAgt,omitempty"`
	InstdAgt    *Agent       `xml:"InstdAgt,omitempty"`
	OrgnlBizQry *OrgnlBizQry `xml:"OrgnlBizQry,omitempty" binding:"omitempty"`
}

type OrgnlBizQry struct {
	MsgId   string `xml:"MsgId" binding:"required,max=35"`
	MsgNmId string `xml:"MsgNmId" binding:"max=35"`
	CreDtTm string `xml:"CreDtTm" binding:"isoDateTime"`
}

type OrgnlGrpInfAndSts struct {
	OrgnlMsgId    string           `xml:"OrgnlMsgId" binding:"required,max=35"`
	OrgnlMsgNmId  string           `xml:"OrgnlMsgNmId" binding:"required,max=35"`
	OrgnlCreDtTm  string           `xml:"OrgnlCreDtTm,omitempty" binding:"isoDateTime"`
	OrgnlNbOfTxs  string           `xml:"OrgnlNbOfTxs,omitempty"`
	OrgnlCtrlSum  float64          `xml:"OrgnlCtrlSum,omitempty"`
	GrpSts        string           `xml:"GrpSts,omitempty"`
	StsRsnInf     []*StsRsnInf     `xml:"StsRsnInf" binding:"dive"`
	NbOfTxsPerSts []*NbOfTxsPerSts `xml:"NbOfTxsPerSts" binding:"dive"`
}

type TxInfAndSts struct {
	StsId             string             `xml:"StsId,omitempty" binding:"max=35"`
	OrgnlGrpInf       *OrgnlGrpInf       `xml:"OrgnlGrpInf,omitempty"`
	OrgnlInstrId      string             `xml:"OrgnlInstrId,omitempty" binding:"max=35"`
	OrgnlEndToEndId   string             `xml:"OrgnlEndToEndId,omitempty" binding:"max=35"`
	OrgnlTxId         string             `xml:"OrgnlTxId,omitempty" binding:"max=35"`
	OrgnlUETR         string             `xml:"OrgnlUETR,omitempty" binding:"uuid4"`
	TxSts             string             `xml:"TxSts,omitempty" binding:"omitempty,min=1,max=4"`
	StsRsnInf         []*StsRsnInf       `xml:"StsRsnInf,omitempty" binding:"dive"`
	ChrgsInf          []*ChrgsInf        `xml:"ChrgsInf,omitempty" binding:"dive"`
	AccptncDtTm       string             `xml:"AccptncDtTm,omitempty" binding:"max=35"`
	PrcgDt            *DtAndDtTmChoice   `xml:"PrcgDt,omitempty"`
	FctvIntrBkSttlmDt *DtAndDtTmChoice   `xml:"FctvIntrBkSttlmDt,omitempty"`
	AcctSvcrRef       string             `xml:"AcctSvcrRef,omitempty" binding:"max=35"`
	ClrSysRef         string             `xml:"ClrSysRef,omitempty" binding:"max=35"`
	CdtSttlmKey       string             `xml:"CdtSttlmKey,omitempty" binding:"regexp=^([0-9A-F]{2}){32}$"`
	InstgAgt          *Agent             `xml:"InstgAgt,omitempty"`
	InstdAgt          *Agent             `xml:"InstdAgt,omitempty"`
	OrgnlTxRef        *OrgnlTxRefPacs002 `xml:"OrgnlTxRef,omitempty"`
	SplmtryData       []*SplmtryData     `xml:"SplmtryData,omitempty" binding:"omitempty,dive"`
}

type StsRsnInf struct {
	Orgtr    *Party   `xml:"Orgtr,omitempty"`
	Rsn      *Rsn     `xml:"Rsn,omitempty"`
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

type OrgnlGrpInf struct {
	OrgnlMsgId   string `xml:"OrgnlMsgId" binding:"required,max=35"`
	OrgnlMsgNmId string `xml:"OrgnlMsgNmId" binding:"required,max=35"`
	OrgnlCreDtTm string `xml:"OrgnlCreDtTm" binding:"isoDateTime"`
}

type DtAndDtTmChoice struct {
	Dt   string `xml:"Dt" binding:"required_without=DtTm,isoDate"`
	DtTm string `xml:"DtTm" binding:"required_without=Dt,isoDateTime"`
}

type OrgnlTxRefPacs002 struct {
	IntrBkSttlmAmt *Amount             `xml:"IntrBkSttlmAmt,omitempty"`
	Amt            *AmountChoice       `xml:"Amt,omitempty"`
	IntrBkSttlmDt  string              `xml:"IntrBkSttlmDt,omitempty" binding:"isoDate"`
	ReqdColltnDt   string              `xml:"ReqdColltnDt,omitempty" binding:"isoDate"`
	ReqdExctnDt    *DtAndDtTmChoice    `xml:"ReqdExctnDt,omitempty"`
	CdtrSchmeId    *Party              `xml:"CdtrSchmeId,omitempty"`
	SttlmInf       *SttlmInf           `xml:"SttlmInf,omitempty"`
	PmtTpInf       *PmtTpInfPacs002    `xml:"PmtTpInf,omitempty"`
	PmtMtd         string              `xml:"PmtMtd,omitempty" binding:"omitempty,oneof=CHK TRF DD TRA"`
	MndtRltdInf    *MndtRltdChoice     `xml:"MndtRltdInf,omitempty"`
	RmtInf         *RmtInf             `xml:"RmtInf,omitempty"`
	UltmtDbtr      *PartyOrAgentChoice `xml:"UltmtDbtr,omitempty"`
	Dbtr           *PartyOrAgentChoice `xml:"Dbtr,omitempty"`
	DbtrAcct       *Account            `xml:"DbtrAcct,omitempty"`
	DbtrAgt        *Agent              `xml:"DbtrAgt,omitempty"`
	DbtrAgtAcct    *Account            `xml:"DbtrAgtAcct,omitempty"`
	CdtrAgt        *Agent              `xml:"CdtrAgt,omitempty"`
	CdtrAgtAcct    *Account            `xml:"CdtrAgtAcct,omitempty"`
	Cdtr           *PartyOrAgentChoice `xml:"Cdtr,omitempty"`
	CdtrAcct       *Account            `xml:"CdtrAcct,omitempty"`
	UltmtCdtr      *PartyOrAgentChoice `xml:"UltmtCdtr,omitempty"`
	Purp           *Purp               `xml:"Purp,omitempty"`
}

type AmountChoice struct {
	InstdAmt *Amount  `xml:"InstdAmt,omitempty" binding:"omitempty,required_without=EqvtAmt"`
	EqvtAmt  *EqvtAmt `xml:"EqvtAmt,omitempty" binding:"omitempty,required_without=InstdAmt"`
}

type EqvtAmt struct {
	Amt      *Amount `xml:"Amt,omitempty" binding:"required"`
	CcyOfTrf string  `xml:"CcyOfTrf,omitempty" binding:"required,regexp=^[A-Z]{3}$"`
}

type PmtTpInfPacs002 struct {
	InstrPrty string     `xml:"InstrPrty,omitempty" binding:"omitempty,oneof=HIGH NORM"`
	ClrChanl  string     `xml:"ClrChanl,omitempty" binding:"omitempty,oneof=RTGS RTNS MPNS BOOK"`
	SvcLvl    []*SvcLvl  `xml:"SvcLvl,omitempty" binding:"dive"`
	LclInstrm *LclInstrm `xml:"LclInstrm,omitempty"`
	SeqTp     string     `xml:"SeqTp,omitempty" binding:"omitempty,oneof=FRST RCUR FNAL OOFF RPRE"`
	CtgyPurp  *CtgyPurp  `xml:"CtgyPurp,omitempty"`
}

type MndtRltdChoice struct {
	DrctDbtMndt *DrctDbtMndt `xml:"DrctDbtMndt,omitempty"`
	CdtTrfMndt  *CdtTrfMndt  `xml:"CdtTrfMndt,omitempty"`
}

type DrctDbtMndt struct {
	MndtId        string         `xml:"MndtId,omitempty" binding:"max=35"`
	DtOfSgntr     string         `xml:"DtOfSgntr,omitempty" binding:"isoDate"`
	AmdmntInd     bool           `xml:"AmdmntInd,omitempty"`
	AmdmntInfDtls *AmdmntInfDtls `xml:"AmdmntInfDtls,omitempty"`
	ElctrncSgntr  string         `xml:"ElctrncSgntr,omitempty" binding:"max=1025"`
	FrstColltnDt  string         `xml:"FrstColltnDt,omitempty" binding:"isoDate"`
	FnlColltnDt   string         `xml:"FnlColltnDt,omitempty" binding:"isoDate"`
	Frqcy         *Frqcy         `xml:"Frqcy,omitempty"`
	Rsn           *Rsn           `xml:"Rsn,omitempty"`
	TrckgDays     string         `xml:"TrckgDays,omitempty" binding:"regexp=^[0-9]{2}$"`
}

type AmdmntInfDtls struct {
	OrgnlMndtId      string   `xml:"OrgnlMndtId,omitempty" binding:"max=35"`
	OrgnlCdtrSchmeId *Party   `xml:"OrgnlCdtrSchmeId,omitempty"`
	OrgnlCdtrAgt     *Agent   `xml:"OrgnlCdtrAgt,omitempty"`
	OrgnlCdtrAgtAcct *Account `xml:"OrgnlCdtrAgtAcct,omitempty"`
	OrgnlDbtr        *Party   `xml:"OrgnlDbtr,omitempty"`
	OrgnlDbtrAcct    *Account `xml:"OrgnlDbtrAcct,omitempty"`
	OrgnlDbtrAgt     *Agent   `xml:"OrgnlDbtrAgt,omitempty"`
	OrgnlDbtrAgtAcct *Account `xml:"OrgnlDbtrAgtAcct,omitempty"`
	OrgnlFnlColltnDt string   `xml:"OrgnlFnlColltnDt,omitempty" binding:"isoDate"`
	OrgnlFrqcy       *Frqcy   `xml:"OrgnlFrqcy,omitempty"`
	OrgnlRsn         *Rsn     `xml:"OrgnlRsn,omitempty"`
	OrgnlTrckgDays   string   `xml:"OrgnlTrckgDays,omitempty" binding:"regexp=^[0-9]{2}$"`
}

type PartyOrAgentChoice struct {
	Pty *Party `xml:"Pty,omitempty" binding:"required_without=Agt"`
	Agt *Agent `xml:"Agt,omitempty" binding:"required_without=Pty"`
}
