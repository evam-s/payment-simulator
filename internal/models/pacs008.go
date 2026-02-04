package models

type Pacs008 struct {
	FIToFICstmrCdtTrf FIToFICstmrCdtTrf `xml:"Document>FIToFICstmrCdtTrf" json:"fiToFiCustomerCreditTransfer"`
}

type FIToFICstmrCdtTrf struct {
	GrpHdr      GroupHeader               `xml:"GrpHdr" json:"groupHeader"`
	CdtTrfTxInf CreditTransferTransaction `xml:"CdtTrfTxInf" json:"creditTransferTransaction"`
	SplmtryData []SplmtryData             `xml:"SplmtryData" json:"supplementaryData"`
}

type GroupHeader struct {
	MsgId             string   `xml:"MsgId" json:"msgId"`
	CreDtTm           string   `xml:"CreDtTm" json:"creationDateTime"`
	NbOfTxs           int      `xml:"NbOfTxs" json:"numberOfTxs"`
	SttlmInf          SttlmInf `xml:"SttlmInf" json:"settlementInfo"`
	XpryDtTm          string   `xml:"XpryDtTm,omitempty" json:"expiryDateTime,omitempty"`
	BtchBookg         bool     `xml:"BtchBookg,omitempty" json:"batchBooking,omitempty"`
	CtrlSum           float64  `xml:"CtrlSum,omitempty" json:"controlSum,omitempty"`
	TtlIntrBkSttlmAmt Amount   `xml:"TtlIntrBkSttlmAmt,omitempty" json:"totalSettlementAmount,omitzero"`
	IntrBkSttlmDt     string   `xml:"IntrBkSttlmDt,omitempty" json:"settlementDate,omitempty"`
	PmtTpInf          PmtTpInf `xml:"PmtTpInf,omitempty" json:"paymentTypeInfo,omitzero"`
	InstgAgt          Agent    `xml:"InstgAgt,omitempty" json:"instructingAgent,omitzero"`
	InstdAgt          Agent    `xml:"InstdAgt,omitempty" json:"instructedAgent,omitzero"`
}

type SttlmInf struct {
	SttlmMtd string `xml:"SttlmMtd" json:"settlementMethod"`
}

type PmtTpInf struct {
	InstrPrty string    `xml:"InstrPrty,omitempty" json:"instructionPriority,omitempty"`
	ClrChanl  string    `xml:"ClrChanl,omitempty" json:"clearingChannel,omitempty"`
	SvcLvl    []SvcLvl  `xml:"SvcLvl,omitempty" json:"serviceLevel,omitempty"`
	LclInstrm LclInstrm `xml:"LclInstrm,omitempty" json:"localInstrument,omitzero"`
	CtgyPurp  CtgyPurp  `xml:"CtgyPurp,omitempty" json:"categoryPurpose,omitzero"`
}

type SvcLvl struct {
	Cd    string `xml:"Cd,omitempty" json:"code,omitempty"`
	Prtry string `xml:"Prtry,omitempty" json:"proprietary,omitempty"`
}

type LclInstrm struct {
	Cd    string `xml:"Cd,omitempty" json:"code,omitempty"`
	Prtry string `xml:"Prtry,omitempty" json:"proprietary,omitempty"`
}

type CtgyPurp struct {
	Cd    string `xml:"Cd,omitempty" json:"code,omitempty"`
	Prtry string `xml:"Prtry,omitempty" json:"proprietary,omitempty"`
}

type Agent struct {
	FinInstnId FinInstnId `xml:"FinInstnId" json:"financialInstitutionId"`
	BrnchId    BrnchId    `xml:"BrnchId,omitempty" json:"branchId,omitzero"`
}

type FinInstnId struct {
	BICFI       string       `xml:"BICFI,omitempty" json:"bicfi,omitempty"`
	ClrSysMmbId ClrSysMmbId  `xml:"ClrSysMmbId,omitempty" json:"clearingSystemMemberId,omitzero"`
	LEI         string       `xml:"LEI,omitempty" json:"lei,omitempty"`
	Nm          string       `xml:"Nm,omitempty" json:"name,omitempty"`
	PstlAdr     PstlAdr      `xml:"PstlAdr,omitempty" json:"postalAddress,omitzero"`
	Othr        FinInstnOthr `xml:"Othr,omitempty" json:"other,omitzero"`
}

type ClrSysMmbId struct {
	ClrSysId ClrSysId `xml:"ClrSysId,omitempty" json:"clearingSystemId,omitzero"`
	MmbId    string   `xml:"MmbId,omitempty" json:"memberId,omitempty"`
}

type ClrSysId struct {
	Cd    string `xml:"Cd,omitempty" json:"code,omitempty"`
	Prtry string `xml:"Prtry,omitempty" json:"proprietary,omitempty"`
}

type BrnchId struct {
	Id      string  `xml:"Id,omitempty" json:"id,omitempty"`
	LEI     string  `xml:"LEI,omitempty" json:"lei,omitempty"`
	Nm      string  `xml:"Nm,omitempty" json:"name,omitempty"`
	PstlAdr PstlAdr `xml:"PstlAdr,omitempty" json:"postalAddress,omitzero"`
}

type CreditTransferTransaction struct {
	PmtId             PaymentId `xml:"PmtId" json:"paymentId"`
	PmtTpInf          PmtTpInf  `xml:"PmtTpInf,omitempty" json:"paymentTypeInfo,omitzero"`
	IntrBkSttlmAmt    Amount    `xml:"IntrBkSttlmAmt" json:"settlementAmount"`
	IntrBkSttlmDt     string    `xml:"IntrBkSttlmDt" json:"settlementDate"`
	SttlmPrty         string    `xml:"SttlmPrty" json:"settlementPriority"`
	SttlmTmIndctn     SttlmTmIndctn
	SttlmTmReq        SttlmTmReq
	AddtlDtTm         AddtlDtTm
	InstdAmt          Amount
	XchgRate          float64
	AgrdRate          AgrdRate
	ChrgBr            string
	ChrgsInf          []ChrgsInf
	MndtRltdInf       MndtRltdInf
	PmtSgntr          PmtSgntr
	PrvsInstgAgt1     Agent
	PrvsInstgAgt1Acct Account
	PrvsInstgAgt2     Agent
	PrvsInstgAgt2Acct Account
	PrvsInstgAgt3     Agent
	PrvsInstgAgt3Acct Account
	InstgAgt          Agent
	InstdAgt          Agent
	IntrmyAgt1        Agent
	IntrmyAgt1Acct    Account
	IntrmyAgt2        Agent
	IntrmyAgt2Acct    Account
	IntrmyAgt3        Agent
	IntrmyAgt3Acct    Account
	UltmtDbtr         Party
	InitgPty          Party
	Dbtr              Party   `xml:"Dbtr" json:"debtor"`
	DbtrAcct          Account `xml:"DbtrAcct" json:"debtorAccount"`
	DbtrAgt           Agent
	DbtrAgtAcct       Account
	CdtrAgt           Agent
	CdtrAgtAcct       Account
	Cdtr              Party   `xml:"Cdtr" json:"creditor"`
	CdtrAcct          Account `xml:"CdtrAcct" json:"creditorAccount"`
	UltmtCdtr         Party
	InstrForCdtrAgt   []InstrForCdtrAgt
	InstrForNxtAgt    []InstrForNxtAgt
	Purp              Purp
	RgltryRptg        []RgltryRptg
	Tax               Tax
	RltdRmtInf        []RltdRmtInf
	RmtInf            RmtInf        `xml:"RmtInf,omitempty" json:"remittanceInfo,omitzero"`
	SplmtryData       []SplmtryData `xml:"SplmtryData" json:"supplementaryData"`
}

type PaymentId struct {
	InstrId    string `xml:"InstrId,omitempty" json:"instrId,omitempty"`
	EndToEndId string `xml:"EndToEndId,omitempty" json:"endToEndId,omitempty"`
	TxId       string `xml:"TxId,omitempty" json:"transactionId,omitempty"`
	UETR       string `xml:"UETR,omitempty" json:"uetr,omitempty"`
	ClrSysRef  string `xml:"ClrSysRef,omitempty" json:"clearingSystemReference,omitempty"`
}

type Amount struct {
	Value    float64 `xml:",chardata" json:"amountValue"`
	Currency string  `xml:"Ccy,attr" json:"amountCurrency"`
}

type SttlmTmIndctn struct {
	DbtDtTm string `xml:"DbtDtTm,omitempty" json:"debitDateTime,omitempty"`
	CdtDtTm string `xml:"CdtDtTm,omitempty" json:"creditDateTime,omitempty"`
}

type SttlmTmReq struct {
	CLSTm  string
	TillTm string
	FrTm   string
	RjctTm string
}

type AddtlDtTm struct {
	AccptncDtTm     string
	PoolgAdjstmntDt string
	XpryDtTm        string
}

type AgrdRate struct {
	UnitCcy         string
	QtdCcy          string
	PreAgrdXchgRate float64
	QtnDtTm         string
	QtId            string
	FXAgt           Agent
}

type ChrgsInf struct {
	Amt Amount
	Agt Agent
	Tp  ChrgsTp
}

type ChrgsTp struct {
	Cd    string
	Prtry ChrgsTpPrtry
}

type ChrgsTpPrtry struct {
	Id   string
	Issr string
}

type MndtRltdInf struct {
	MndtId       string
	Tp           MndtRltdInfTp
	DtOfSgntr    string
	DtOfVrfctn   string
	ElctrncSgntr []byte
	FrstPmtDt    string
	FnlPmtDt     string
	Frqcy        Frqcy
	Rsn          MndtRltdInfRsn
}

type MndtRltdInfTp struct {
	SvcLvl    []SvcLvl
	LclInstrm LclInstrm
	CtgyPurp  CtgyPurp
	Clssfctn  MndtRltdInfTpClssfctn
}

type MndtRltdInfTpClssfctn struct {
	Cd    string
	Prtry string
}

type Frqcy struct {
	Tp     string
	Prd    string
	PtInTm PtInTm
}

type FrqcyPrd struct {
	Tp        string
	CntPerPrd float64
}

type PtInTm struct {
	Tp     string
	PtInTm string
}

type MndtRltdInfRsn struct {
	Cd    string
	Prtry string
}

type PmtSgntr struct {
	ILPV4 string
	Sgntr string
}

type Party struct {
	Nm        string   `xml:"Nm,omitempty" json:"name,omitempty"`
	PstlAdr   PstlAdr  `xml:"PstlAdr,omitempty" json:"postalAddress,omitzero"`
	Id        PartyId  `xml:"Id,omitempty" json:"partyId,omitzero"`
	CtryOfRes string   `xml:"CtryOfRes,omitempty" json:"countryOfResidence,omitempty"`
	CtctDtls  CtctDtls `xml:"CtctDtls,omitempty" json:"contactDetails,omitzero"`
}

type FinInstnOthr struct {
	Id      string          `xml:"Id,omitempty" json:"id,omitempty"`
	Issr    string          `xml:"Issr,omitempty" json:"issuer,omitempty"`
	SchmeNm FinInstnSchmeNm `xml:"SchmeNm,omitempty" json:"schemeName,omitzero"`
}

type FinInstnSchmeNm struct {
	Cd    string `xml:"Cd,omitempty" json:"code,omitempty"`
	Prtry string `xml:"Prtry,omitempty" json:"proprietary,omitempty"`
}

type PstlAdr struct {
	AdrTp       AdressType `xml:"AdrTp,omitempty" json:"addressType,omitzero"`
	CareOf      string     `xml:"CareOf,omitempty" json:"careOf,omitempty"`
	Dept        string     `xml:"Dept,omitempty" json:"department,omitempty"`
	SubDept     string     `xml:"SubDept,omitempty" json:"subDepartment,omitempty"`
	StrtNm      string     `xml:"StrtNm,omitempty" json:"streetName,omitempty"`
	BldgNb      string     `xml:"BldgNb,omitempty" json:"buildingNumber,omitempty"`
	BldgNm      string     `xml:"BldgNm,omitempty" json:"buildingName,omitempty"`
	Flr         string     `xml:"Flr,omitempty" json:"floor,omitempty"`
	UnitNb      string     `xml:"UnitNb,omitempty" json:"unitNumber,omitempty"`
	PstBx       string     `xml:"PstBx,omitempty" json:"postBox,omitempty"`
	Room        string     `xml:"Room,omitempty" json:"room,omitempty"`
	PstCd       string     `xml:"PstCd,omitempty" json:"postalCode,omitempty"`
	TwnNm       string     `xml:"TwnNm,omitempty" json:"townName,omitempty"`
	TwnLctnNm   string     `xml:"TwnLctnNm,omitempty" json:"townLocationName,omitempty"`
	DstrctNm    string     `xml:"DstrctNm,omitempty" json:"districtName,omitempty"`
	CtrySubDvsn string     `xml:"CtrySubDvsn,omitempty" json:"countrySubdivision,omitempty"`
	Ctry        string     `xml:"Ctry,omitempty" json:"country,omitempty"`
	AdrLine     string     `xml:"AdrLine,omitempty" json:"addressLine,omitempty"`
}

type AdressType struct {
	Cd    string     `xml:"Cd,omitempty" json:"code,omitempty"`
	Prtry AdrTpPrtry `xml:"Prtry,omitempty" json:"proprietary,omitzero"`
}

type AdrTpPrtry struct {
	Id      string `xml:"Id,omitempty" json:"id,omitempty"`
	Issr    string `xml:"Issr,omitempty" json:"issuer,omitempty"`
	SchmeNm string `xml:"SchmeNm,omitempty" json:"schemeName,omitempty"`
}

type PartyId struct {
	OrgId  OrgId  `xml:"OrgId,omitempty" json:"orgId,omitzero"`
	PrvtId PrvtId `xml:"PrvtId,omitempty" json:"privateId,omitzero"`
}

type OrgId struct {
	AnyBIC string  `xml:"AnyBIC,omitempty" json:"anyBic,omitempty"`
	LEI    string  `xml:"LEI,omitempty" json:"lei,omitempty"`
	Othr   OrgOthr `xml:"Othr,omitempty" json:"other,omitzero"`
}

type OrgOthr struct {
	Id      string     `xml:"Id,omitempty" json:"id,omitempty"`
	Issr    string     `xml:"Issr,omitempty" json:"issuer,omitempty"`
	SchmeNm OrgSchmeNm `xml:"SchmeNm,omitempty" json:"schemeName,omitzero"`
}

type OrgSchmeNm struct {
	Cd    string `xml:"Cd,omitempty" json:"code,omitempty"`
	Prtry string `xml:"Prtry,omitempty" json:"proprietary,omitempty"`
}

type PrvtId struct {
	DtAndPlcOfBirth DateAndPlaceOfBirth `xml:"DtAndPlcOfBirth,omitempty" json:"dateAndPlaceOfBirth,omitzero"`
	Othr            PersonOthr          `xml:"Othr,omitempty" json:"other,omitzero"`
}

type PersonOthr struct {
	Id      string        `xml:"Id,omitempty" json:"id,omitempty"`
	Issr    string        `xml:"Issr,omitempty" json:"issuer,omitempty"`
	SchmeNm PersonSchmeNm `xml:"SchmeNm,omitempty" json:"schemeName,omitzero"`
}

type PersonSchmeNm struct {
	Cd    string `xml:"Cd,omitempty" json:"code,omitempty"`
	Prtry string `xml:"Prtry,omitempty" json:"proprietary,omitempty"`
}

type DateAndPlaceOfBirth struct {
	BirthDt     string `xml:"BirthDt,omitempty" json:"birthDate,omitempty"`
	PrvcOfBirth string `xml:"PrvcOfBirth,omitempty" json:"provinceOfBirth,omitempty"`
	CityOfBirth string `xml:"CityOfBirth,omitempty" json:"cityOfBirth,omitempty"`
	CtryOfBirth string `xml:"CtryOfBirth,omitempty" json:"countryOfBirth,omitempty"`
}

type CtctDtls struct {
	NmPrfx    string      `xml:"NmPrfx,omitempty" json:"namePrefix,omitempty"`
	Nm        string      `xml:"Nm,omitempty" json:"name,omitempty"`
	PhneNb    string      `xml:"PhneNb,omitempty" json:"phoneNumber,omitempty"`
	MobNb     string      `xml:"MobNb,omitempty" json:"mobileNumber,omitempty"`
	FaxNb     string      `xml:"FaxNb,omitempty" json:"faxNumber,omitempty"`
	URLAdr    string      `xml:"URLAdr,omitempty" json:"urlAddress,omitempty"`
	EmailAdr  string      `xml:"EmailAdr,omitempty" json:"emailAddress,omitempty"`
	EmailPurp string      `xml:"EmailPurp,omitempty" json:"emailPurpose,omitempty"`
	JobTitl   string      `xml:"JobTitl,omitempty" json:"jobTitle,omitempty"`
	Rspnsblty string      `xml:"Rspnsblty,omitempty" json:"responsibility,omitempty"`
	Dept      string      `xml:"Dept,omitempty" json:"department,omitempty"`
	Othr      ContactOthr `xml:"Othr,omitempty" json:"other,omitzero"`
	PrefrdMtd string      `xml:"PrefrdMtd,omitempty" json:"preferredMethod,omitempty"`
}

type ContactOthr struct {
	Id      string `xml:"Id,omitempty" json:"id,omitempty"`
	ChanlTp string `xml:"ChanlTp,omitempty" json:"channelType,omitempty"`
}

type Account struct {
	Id   AccountId    `xml:"Id" json:"accountId"`
	Tp   AccountType  `xml:"Tp,omitempty" json:"type,omitzero"`
	Ccy  string       `xml:"Ccy,omitempty" json:"currency,omitempty"`
	Nm   string       `xml:"Nm,omitempty" json:"name,omitempty"`
	Prxy AccountProxy `xml:"Prxy,omitempty" json:"proxy,omitzero"`
}

type AccountId struct {
	IBAN string      `xml:"IBAN,omitempty" json:"iban,omitempty"`
	Othr AccountOthr `xml:"Othr,omitempty" json:"other,omitzero"`
}

type AccountOthr struct {
	Id      string         `xml:"Id,omitempty" json:"id,omitempty"`
	Issr    string         `xml:"Issr,omitempty" json:"issuer,omitempty"`
	SchmeNm AccountSchmeNm `xml:"SchmeNm,omitempty" json:"schemeName,omitzero"`
}

type AccountSchmeNm struct {
	Cd    string `xml:"Cd,omitempty" json:"code,omitempty"`
	Prtry string `xml:"Prtry,omitempty" json:"proprietary,omitempty"`
}

type AccountType struct {
	Cd    string `xml:"Cd,omitempty" json:"code,omitempty"`
	Prtry string `xml:"Prtry,omitempty" json:"proprietary,omitempty"`
}

type AccountProxy struct {
	Tp string `xml:"Tp,omitempty" json:"type,omitempty"`
	Id string `xml:"Id,omitempty" json:"id,omitempty"`
}

type InstrForCdtrAgt struct {
	Cd       string
	InstrInf string
}

type InstrForNxtAgt struct {
	Cd       string
	InstrInf string
}

type Purp struct {
	Cd    string
	Prtry string
}

type RgltryRptg struct {
	DbtCdtRptgInd string
	Authrty       RgltryRptgAuthrty
	Dtls          []RgltryRptgStrd
}

type RgltryRptgAuthrty struct {
	Nm   string
	Ctry string
}

type RgltryRptgStrd struct {
	Tp   string
	Dt   string
	Ctry string
	Cd   string
	Amt  Amount
	Inf  []string
}

type Tax struct {
	Cdtr            TaxParty1
	Dbtr            TaxParty2
	UltmtDbtr       TaxParty2
	AdmstnZone      string
	RefNb           string
	Mtd             string
	TtlTaxblBaseAmt Amount
	TtlTaxAmt       Amount
	Dt              string
	SeqNb           int64
	Rcrd            []TaxRcrd
}

type TaxParty1 struct {
	TaxId  string
	RegnId string
	TaxTp  string
}

type TaxParty2 struct {
	TaxId   string
	RegnId  string
	TaxTp   string
	Authstn TaxPartyAuthstn
}

type TaxPartyAuthstn struct {
	Titl string
	Nm   string
}

type TaxRcrd struct {
	Tp       string
	Ctgy     string
	CtgyDtls string
	DbtrSts  string
	CertId   string
	FrmsCd   string
	Prd      TaxPrd
	TaxAmt   TaxAmt
	AddtlInf string
}

type TaxPrd struct {
	Yr     int16
	Tp     string
	FrToDt DatePrd
}

type DatePrd struct {
	FrDt string
	ToDt string
}

type TaxAmt struct {
	Rate         float64
	TaxblBaseAmt Amount
	TtlAmt       Amount
	Dtls         TaxAmtDtls
}

type TaxAmtDtls struct {
	Prd TaxPrd
	Amt Amount
}

type RltdRmtInf struct {
	RmtId       string
	RmtLctnDtls []RmtLctnDtls
}

type RmtLctnDtls struct {
	Mtd        string
	ElctrncAdr string
	PstlAdr    NmAndAdr
}

type NmAndAdr struct {
	Nm  string
	Adr PstlAdr
}
type RmtInf struct {
	Ustrd []string `xml:"Ustrd,omitempty" json:"ustrdRemittanceInfo,omitempty"`
	Strd  []RmtInfStrd
}
type RmtInfStrd struct {
	RfrdDocInf  []RfrdDocInf
	RfrdDocAmt  LineDtlsAmount
	CdtrRefInf  CdtrRefInf
	Invcr       Party
	Invcee      Party
	TaxRmt      Tax
	GrnshmtRmt  GrnshmtRmt
	AddtlRmtInf []string
}

type RfrdDocInf struct {
	Tp       RfrdDocInfTp
	Nb       string
	RltdDt   DtAndTp
	LineDtls []LineDtls
}

type RfrdDocInfTp struct {
	CdOrPrtry CdOrPrtry
	Issr      string
}

type CdOrPrtry struct {
	Cd    string
	Prtry string
}

type DtAndTp struct {
	Tp string
	Dt string
}

type LineDtls struct {
	Id   []LineDtlsId
	Desc string
	Amt  LineDtlsAmount
}

type LineDtlsId struct {
	Tp     LineDtlsIdTp
	Nb     string
	RltdDt string
}

type LineDtlsIdTp struct {
	CdOrPrtry CdOrPrtry
	Issr      string
}

type LineDtlsAmount struct {
	RmtAmtAndTp       []LineDtlsRmtAmtAndTp
	AdjstmntAmtAndRsn []LineDtlsAdjstmntAmtAndRsn
}

type LineDtlsRmtAmtAndTp struct {
	Tp  LineDtlsCdOrPrtry
	Amt Amount
}

type LineDtlsCdOrPrtry struct {
	Cd    string
	Prtry string
}

type LineDtlsAdjstmntAmtAndRsn struct {
	Amt       Amount
	CdtDbtInd string
	Rsn       string
	AddtlInf  string
}

type CdtrRefInf struct {
	Tp  CdtrRefInfTp
	Ref string
}

type CdtrRefInfTp struct {
	CdOrPrtry CdOrPrtry
	Issr      string
}

type GrnshmtRmt struct {
	Tp                GrnshmtRmtTp
	Grnshee           Party
	GrnshmtAdmstr     Party
	RefNb             string
	Dt                string
	RmtdAmt           Amount
	FmlyMdclInsrncInd bool
	MplyeeTermntnInd  bool
}

type GrnshmtRmtTp struct {
	CdOrPrtry CdOrPrtry
	Issr      string
}

type SplmtryData struct {
	PlcAndNm string `xml:"PlcAndNm,omitempty" json:"placeAndName,omitempty"`
	Envlp    []byte `xml:"Envlp,omitempty" json:"envelope,omitempty"`
}
