package isomodels

type Pacs008 struct {
	FIToFICstmrCdtTrf fIToFICstmrCdtTrf `xml:"FIToFICstmrCdtTrf" json:"fiToFiCustomerCreditTransfer" binding:"required"`
}

type fIToFICstmrCdtTrf struct {
	GrpHdr      groupHeader                 `xml:"GrpHdr" json:"groupHeader" binding:"required"`
	CdtTrfTxInf []creditTransferTransaction `xml:"CdtTrfTxInf" json:"creditTransferTransaction" binding:"required"`
	SplmtryData []splmtryData               `xml:"SplmtryData" json:"supplementaryData"`
}

type groupHeader struct {
	MsgId             string   `xml:"MsgId" json:"messageId" binding:"required,max=35"`
	CreDtTm           string   `xml:"CreDtTm" json:"creationDateTime" binding:"required"`
	NbOfTxs           string   `xml:"NbOfTxs" json:"numberOfTxs" binding:"required,numeric,omitempty,min=1,max=15"`
	SttlmInf          sttlmInf `xml:"SttlmInf" json:"settlementInfo" binding:"required"`
	XpryDtTm          string   `xml:"XpryDtTm,omitempty" json:"expiryDateTime,omitempty"`
	BtchBookg         bool     `xml:"BtchBookg,omitempty" json:"batchBooking,omitempty"`
	CtrlSum           float64  `xml:"CtrlSum,omitempty" json:"controlSum,omitempty"`
	TtlIntrBkSttlmAmt amount   `xml:"TtlIntrBkSttlmAmt,omitempty" json:"totalSettlementAmount,omitzero"`
	IntrBkSttlmDt     string   `xml:"IntrBkSttlmDt,omitempty" json:"settlementDate,omitempty"`
	PmtTpInf          pmtTpInf `xml:"PmtTpInf,omitempty" json:"paymentTypeInfo,omitzero"`
	InstgAgt          agent    `xml:"InstgAgt,omitempty" json:"instructingAgent,omitzero"`
	InstdAgt          agent    `xml:"InstdAgt,omitempty" json:"instructedAgent,omitzero"`
}

type sttlmInf struct {
	SttlmMtd string `xml:"SttlmMtd" json:"settlementMethod"`
}

type pmtTpInf struct {
	InstrPrty string    `xml:"InstrPrty,omitempty" json:"instructionPriority,omitempty" binding:"omitempty,oneof=HIGH NORM"`
	ClrChanl  string    `xml:"ClrChanl,omitempty" json:"clearingChannel,omitempty" binding:"omitempty,oneof=RTGS RTNS MPNS BOOK"`
	SvcLvl    []svcLvl  `xml:"SvcLvl,omitempty" json:"serviceLevel,omitempty"`
	LclInstrm lclInstrm `xml:"LclInstrm,omitempty" json:"localInstrument,omitzero"`
	CtgyPurp  ctgyPurp  `xml:"CtgyPurp,omitempty" json:"categoryPurpose,omitzero"`
}

type svcLvl struct {
	Cd    string `xml:"Cd,omitempty" json:"code,omitempty" binding:"omitempty,min=1,max=4,required_without=Prtry"`
	Prtry string `xml:"Prtry,omitempty" json:"proprietary,omitempty" binding:"max=35,required_without=Cd"`
}

type lclInstrm struct {
	Cd    string `xml:"Cd,omitempty" json:"code,omitempty" binding:"omitempty,min=1,max=35,required_without=Prtry"`
	Prtry string `xml:"Prtry,omitempty" json:"proprietary,omitempty" binding:"max=35,required_without=Cd"`
}

type ctgyPurp struct {
	Cd    string `xml:"Cd,omitempty" json:"code,omitempty" binding:"omitempty,min=1,max=4,required_without=Prtry"`
	Prtry string `xml:"Prtry,omitempty" json:"proprietary,omitempty" binding:"max=35,required_without=Cd"`
}

type agent struct {
	FinInstnId finInstnId `xml:"FinInstnId" json:"financialInstitutionId" binding:"required"`
	BrnchId    brnchId    `xml:"BrnchId,omitempty" json:"branchId,omitzero"`
}

type finInstnId struct {
	BICFI       string       `xml:"BICFI,omitempty" json:"bicfi,omitempty" binding:"regexp=^[A-Z0-9]{4}[A-Z]{2}[A-Z0-9]{2}([A-Z0-9]{3})?$"`
	ClrSysMmbId clrSysMmbId  `xml:"ClrSysMmbId,omitempty" json:"clearingSystemMemberId,omitzero"`
	LEI         string       `xml:"LEI,omitempty" json:"lei,omitempty" binding:"regexp=^[A-Z0-9]{18}[0-9]{2}$"`
	Nm          string       `xml:"Nm,omitempty" json:"name,omitempty" binding:"max=140"`
	PstlAdr     pstlAdr      `xml:"PstlAdr,omitempty" json:"postalAddress,omitzero"`
	Othr        finInstnOthr `xml:"Othr,omitempty" json:"other,omitzero"`
}

type clrSysMmbId struct {
	ClrSysId clrSysId `xml:"ClrSysId,omitempty" json:"clearingSystemId,omitzero"`
	MmbId    string   `xml:"MmbId,omitempty" json:"memberId,omitempty" binding:"max=35"`
}

type clrSysId struct {
	Cd    string `xml:"Cd,omitempty" json:"code,omitempty" binding:"omitempty,min=1,max=5"`
	Prtry string `xml:"Prtry,omitempty" json:"proprietary,omitempty" binding:"max=35"`
}

type brnchId struct {
	Id      string  `xml:"Id,omitempty" json:"id,omitempty" binding:"max=35"`
	LEI     string  `xml:"LEI,omitempty" json:"lei,omitempty" binding:"regexp=^[A-Z0-9]{18}[0-9]{2}$"`
	Nm      string  `xml:"Nm,omitempty" json:"name,omitempty" binding:"max=140"`
	PstlAdr pstlAdr `xml:"PstlAdr,omitempty" json:"postalAddress,omitzero"`
}

type creditTransferTransaction struct {
	PmtId             pmtId             `xml:"PmtId" json:"paymentId" binding:"required"`
	PmtTpInf          pmtTpInf          `xml:"PmtTpInf,omitempty" json:"paymentTypeInfo,omitzero"`
	IntrBkSttlmAmt    amount            `xml:"IntrBkSttlmAmt" json:"settlementAmount" binding:"required"`
	IntrBkSttlmDt     string            `xml:"IntrBkSttlmDt" json:"settlementDate" binding:"required"`
	SttlmPrty         string            `xml:"SttlmPrty,omitempty" json:"settlementPriority,omitempty"`
	SttlmTmIndctn     sttlmTmIndctn     `xml:"SttlmTmIndctn,omitempty" json:"settlementTimeIndication,omitzero"`
	SttlmTmReq        sttlmTmReq        `xml:"SttlmTmReq,omitempty" json:"settlementTimeRequest,omitzero"`
	AddtlDtTm         addtlDtTm         `xml:"AddtlDtTm,omitempty" json:"additionalDateTime,omitzero"`
	InstdAmt          amount            `xml:"InstdAmt,omitempty" json:"instructedAmount,omitzero"`
	XchgRate          float64           `xml:"XchgRate,omitempty" json:"exchangeRate,omitempty"`
	AgrdRate          agrdRate          `xml:"AgrdRate,omitempty" json:"agreedRate,omitzero"`
	ChrgBr            string            `xml:"ChrgBr,omitempty" json:"chargeBearer,omitempty" binding:"omitempty,oneof=DEBT CRED SHAR SLEV"`
	ChrgsInf          []chrgsInf        `xml:"ChrgsInf,omitempty" json:"chargesInfo,omitempty"`
	MndtRltdInf       mndtRltdInf       `xml:"MndtRltdInf,omitempty" json:"mandateRelatedInfo,omitzero"`
	PmtSgntr          pmtSgntr          `xml:"PmtSgntr,omitempty" json:"paymentSignature,omitzero"`
	PrvsInstgAgt1     agent             `xml:"PrvsInstgAgt1,omitempty" json:"pia1,omitzero"`
	PrvsInstgAgt1Acct account           `xml:"PrvsInstgAgt1Acct,omitempty" json:"pia1Account,omitzero"`
	PrvsInstgAgt2     agent             `xml:"PrvsInstgAgt2,omitempty" json:"pia2,omitzero"`
	PrvsInstgAgt2Acct account           `xml:"PrvsInstgAgt2Acct,omitempty" json:"pia2Account,omitzero"`
	PrvsInstgAgt3     agent             `xml:"PrvsInstgAgt3,omitempty" json:"pia3,omitzero"`
	PrvsInstgAgt3Acct account           `xml:"PrvsInstgAgt3Acct,omitempty" json:"pia3Account,omitzero"`
	InstgAgt          agent             `xml:"InstgAgt,omitempty" json:"instructingAgent,omitzero"`
	InstdAgt          agent             `xml:"InstdAgt,omitempty" json:"instructedAgent,omitzero"`
	IntrmyAgt1        agent             `xml:"IntrmyAgt1,omitempty" json:"ia1,omitzero"`
	IntrmyAgt1Acct    account           `xml:"IntrmyAgt1Acct,omitempty" json:"ia1Account,omitzero"`
	IntrmyAgt2        agent             `xml:"IntrmyAgt2,omitempty" json:"ia2,omitzero"`
	IntrmyAgt2Acct    account           `xml:"IntrmyAgt2Acct,omitempty" json:"ia2Account,omitzero"`
	IntrmyAgt3        agent             `xml:"IntrmyAgt3,omitempty" json:"ia3,omitzero"`
	IntrmyAgt3Acct    account           `xml:"IntrmyAgt3Acct,omitempty" json:"ia3Account,omitzero"`
	UltmtDbtr         party             `xml:"UltmtDbtr,omitempty" json:"ultimateDebtor,omitzero"`
	InitgPty          party             `xml:"InitgPty,omitempty" json:"initiatingParty,omitzero"`
	Dbtr              party             `xml:"Dbtr" json:"debtor" binding:"required"`
	DbtrAcct          account           `xml:"DbtrAcct" json:"debtorAccount" binding:"required"`
	DbtrAgt           agent             `xml:"DbtrAgt,omitempty" json:"debtorAgent,omitzero"`
	DbtrAgtAcct       account           `xml:"DbtrAgtAcct,omitempty" json:"debtorAgentAccount,omitzero"`
	CdtrAgt           agent             `xml:"CdtrAgt,omitempty" json:"creditorAgent,omitzero"`
	CdtrAgtAcct       account           `xml:"CdtrAgtAcct,omitempty" json:"creditorAgentAccount,omitzero"`
	Cdtr              party             `xml:"Cdtr" json:"creditor" binding:"required"`
	CdtrAcct          account           `xml:"CdtrAcct" json:"creditorAccount" binding:"required"`
	UltmtCdtr         party             `xml:"UltmtCdtr,omitempty" json:"ultimateCreditor,omitzero"`
	InstrForCdtrAgt   []instrForCdtrAgt `xml:"InstrForCdtrAgt,omitempty" json:"instructionsForCreditorAgent,omitempty"`
	InstrForNxtAgt    []instrForNxtAgt  `xml:"InstrForNxtAgt,omitempty" json:"instructionsForNextAgent,omitempty"`
	Purp              purp              `xml:"Purp,omitempty" json:"purpose,omitzero"`
	RgltryRptg        []rgltryRptg      `xml:"RgltryRptg,omitempty" json:"regulatoryReporting,omitempty" binding:"max=10"`
	Tax               tax               `xml:"Tax,omitempty" json:"tax,omitzero"`
	RltdRmtInf        []rltdRmtInf      `xml:"RltdRmtInf,omitempty" json:"relatedRemittanceInfo,omitempty"`
	RmtInf            rmtInf            `xml:"RmtInf,omitempty" json:"remittanceInfo,omitzero"`
	SplmtryData       []splmtryData     `xml:"SplmtryData,omitempty" json:"supplementaryData,omitzero"`
}

type pmtId struct {
	InstrId    string `xml:"InstrId" json:"instructionId" binding:"max=35"`
	EndToEndId string `xml:"EndToEndId" json:"endToEndId" binding:"required,max=35"`
	TxId       string `xml:"TxId" json:"transactionId" binding:"max=35"`
	UETR       string `xml:"UETR" json:"uetr" binding:"uuid4"`
	ClrSysRef  string `xml:"ClrSysRef,omitempty" json:"clearingSystemReference,omitempty" binding:"max=35"`
}

type amount struct {
	Value    float64 `xml:",chardata" json:"amountValue" binding:"required"`
	Currency string  `xml:"Ccy,attr" json:"amountCurrency" binding:"required,regexp=^[A-Z]{3}$"`
}

type sttlmTmIndctn struct {
	DbtDtTm string `xml:"DbtDtTm,omitempty" json:"debitDateTime,omitempty"`
	CdtDtTm string `xml:"CdtDtTm,omitempty" json:"creditDateTime,omitempty"`
}

type sttlmTmReq struct {
	CLSTm  string `xml:"CLSTm,omitempty" json:"closeTime,omitempty"`
	TillTm string `xml:"TillTm,omitempty" json:"tillTime,omitempty"`
	FrTm   string `xml:"FrTm,omitempty" json:"fromTime,omitempty"`
	RjctTm string `xml:"RjctTm,omitempty" json:"rejectTime,omitempty"`
}

type addtlDtTm struct {
	AccptncDtTm     string `xml:"AccptncDtTm,omitempty" json:"acceptanceDateTime,omitempty"`
	PoolgAdjstmntDt string `xml:"PoolgAdjstmntDt,omitempty" json:"poolingAdjustmentDate,omitempty"`
	XpryDtTm        string `xml:"XpryDtTm,omitempty" json:"expiryDateTime,omitempty"`
}

type agrdRate struct {
	UnitCcy         string  `xml:"UnitCcy,omitempty" json:"unitCurrency,omitempty" binding:"regexp=^[A-Z]{3}$"`
	QtdCcy          string  `xml:"QtdCcy,omitempty" json:"quotedCurrency,omitempty" binding:"regexp=^[A-Z]{3}$"`
	PreAgrdXchgRate float64 `xml:"PreAgrdXchgRate" json:"preAgreedExchangeRate" binding:"required"`
	QtnDtTm         string  `xml:"QtnDtTm,omitempty" json:"quotationDateTime,omitempty"`
	QtId            string  `xml:"QtId,omitempty" json:"quoteId,omitempty" binding:"uuid4"`
	FXAgt           agent   `xml:"FXAgt,omitempty" json:"foreignExchangeAgent,omitzero"`
}

type chrgsInf struct {
	Amt amount  `xml:"Amt" json:"amount" binding:"required"`
	Agt agent   `xml:"Agt" json:"agent" binding:"required"`
	Tp  chrgsTp `xml:"Tp,omitempty" json:"type,omitzero"`
}

type chrgsTp struct {
	Cd    string       `xml:"Cd,omitempty" json:"code,omitempty"`
	Prtry chrgsTpPrtry `xml:"Prtry,omitempty" json:"proprietary,omitzero"`
}

type chrgsTpPrtry struct {
	Id   string `xml:"Id,omitempty" json:"id,omitempty"`
	Issr string `xml:"Issr,omitempty" json:"issuer,omitempty"`
}

type mndtRltdInf struct {
	MndtId       string         `xml:"MndtId,omitempty" json:"mandateId,omitempty" binding:"max=35"`
	Tp           mndtRltdInfTp  `xml:"Tp,omitempty" json:"type,omitzero"`
	DtOfSgntr    string         `xml:"DtOfSgntr,omitempty" json:"dateOfSignature,omitempty"`
	DtOfVrfctn   string         `xml:"DtOfVrfctn,omitempty" json:"dateOfVerification,omitempty"`
	ElctrncSgntr []byte         `xml:"ElctrncSgntr,omitempty" json:"electronicSignature,omitempty"`
	FrstPmtDt    string         `xml:"FrstPmtDt,omitempty" json:"firstPaymentDate,omitempty"`
	FnlPmtDt     string         `xml:"FnlPmtDt,omitempty" json:"finalPaymentDate,omitempty"`
	Frqcy        frqcy          `xml:"Frqcy,omitempty" json:"frequency,omitzero"`
	Rsn          mndtRltdInfRsn `xml:"Rsn,omitempty" json:"reason,omitzero"`
}

type mndtRltdInfTp struct {
	SvcLvl    []svcLvl              `xml:"SvcLvl,omitempty" json:"serviceLevel,omitempty"`
	LclInstrm lclInstrm             `xml:"LclInstrm,omitempty" json:"localInstrument,omitzero"`
	CtgyPurp  ctgyPurp              `xml:"CtgyPurp,omitempty" json:"categoryPurpose,omitzero"`
	Clssfctn  mndtRltdInfTpClssfctn `xml:"Clssfctn,omitempty" json:"classification,omitzero"`
}

type mndtRltdInfTpClssfctn struct {
	Cd    string `xml:"Cd,omitempty" json:"code,omitempty" binding:"required_without=Prtry,omitempty,oneof=FIXE USGB VARI"`
	Prtry string `xml:"Prtry,omitempty" json:"proprietary,omitempty" binding:"max=35,required_without=Cd"`
}

type frqcy struct {
	Tp     string   `xml:"Tp,omitempty" json:"type,omitempty" binding:"required_without_all=Prd PtInTm,omitempty,oneof=YEAR MNTH QURT MIAN WEEK DAIL ADHO INDA FRTN"`
	Prd    frqcyPrd `xml:"Prd,omitempty" json:"period,omitzero" binding:"required_without_all=Tp PtInTm"`
	PtInTm ptInTm   `xml:"PtInTm,omitempty" json:"pointInTime,omitzero" binding:"required_without_all=Tp Prd"`
}

type frqcyPrd struct {
	Tp        string  `xml:"Tp" json:"type" binding:"required,oneof=YEAR MNTH QURT MIAN WEEK DAIL ADHO INDA FRTN"`
	CntPerPrd float64 `xml:"CntPerPrd" json:"countPerPeriod" binding:"required"`
}

type ptInTm struct {
	Tp     string `xml:"Tp" json:"type" binding:"required,oneof=YEAR MNTH QURT MIAN WEEK DAIL ADHO INDA FRTN"`
	PtInTm string `xml:"PtInTm,omitempty" json:"pointInTime,omitempty" binding:"regexp=^[0-9]{2}$"`
}

type mndtRltdInfRsn struct {
	Cd    string `xml:"Cd,omitempty" json:"code,omitempty" binding:"omitempty,min=1,max=4,required_without=Prtry"`
	Prtry string `xml:"Prtry,omitempty" json:"proprietary,omitempty" binding:"max=35,required_without=Cd"`
}

type pmtSgntr struct {
	ILPV4 string `xml:"ILPV4,omitempty" json:"ilpv4,omitempty" binding:"hexadecimal,required_without=Sgntr"`
	Sgntr string `xml:"Sgntr,omitempty" json:"signature,omitempty" binding:"regexp=^([0-9a-fA-F]{2}){32}$required_without=ILPV4"`
}

type party struct {
	Nm        string   `xml:"Nm,omitempty" json:"name,omitempty"`
	PstlAdr   pstlAdr  `xml:"PstlAdr,omitempty" json:"postalAddress,omitzero"`
	Id        partyId  `xml:"Id,omitempty" json:"partyId,omitzero"`
	CtryOfRes string   `xml:"CtryOfRes,omitempty" json:"countryOfResidence,omitempty" binding:"regexp=^[A-Z]{2}$"`
	CtctDtls  ctctDtls `xml:"CtctDtls,omitempty" json:"contactDetails,omitzero"`
}

type finInstnOthr struct {
	Id      string          `xml:"Id,omitempty" json:"id,omitempty" binding:"max=35"`
	Issr    string          `xml:"Issr,omitempty" json:"issuer,omitempty" binding:"max=35"`
	SchmeNm finInstnSchmeNm `xml:"SchmeNm,omitempty" json:"schemeName,omitzero"`
}

type finInstnSchmeNm struct {
	Cd    string `xml:"Cd,omitempty" json:"code,omitempty" binding:"required_without=Prtry,max=4"`
	Prtry string `xml:"Prtry,omitempty" json:"proprietary,omitempty" binding:"required_without=Cd,max=35"`
}

type pstlAdr struct {
	AdrTp       adressType `xml:"AdrTp,omitempty" json:"addressType,omitzero"`
	CareOf      string     `xml:"CareOf,omitempty" json:"careOf,omitempty" binding:"max=140"`
	Dept        string     `xml:"Dept,omitempty" json:"department,omitempty" binding:"max=70"`
	SubDept     string     `xml:"SubDept,omitempty" json:"subDepartment,omitempty" binding:"max=70"`
	StrtNm      string     `xml:"StrtNm,omitempty" json:"streetName,omitempty" binding:"max=140"`
	BldgNb      string     `xml:"BldgNb,omitempty" json:"buildingNumber,omitempty" binding:"max=16"`
	BldgNm      string     `xml:"BldgNm,omitempty" json:"buildingName,omitempty" binding:"max=140"`
	Flr         string     `xml:"Flr,omitempty" json:"floor,omitempty" binding:"max=70"`
	UnitNb      string     `xml:"UnitNb,omitempty" json:"unitNumber,omitempty" binding:"max=16"`
	PstBx       string     `xml:"PstBx,omitempty" json:"postBox,omitempty" binding:"max=16"`
	Room        string     `xml:"Room,omitempty" json:"room,omitempty" binding:"max=70"`
	PstCd       string     `xml:"PstCd,omitempty" json:"postalCode,omitempty" binding:"max=16"`
	TwnNm       string     `xml:"TwnNm,omitempty" json:"townName,omitempty" binding:"max=140"`
	TwnLctnNm   string     `xml:"TwnLctnNm,omitempty" json:"townLocationName,omitempty" binding:"max=140"`
	DstrctNm    string     `xml:"DstrctNm,omitempty" json:"districtName,omitempty" binding:"max=140"`
	CtrySubDvsn string     `xml:"CtrySubDvsn,omitempty" json:"countrySubdivision,omitempty" binding:"max=35"`
	Ctry        string     `xml:"Ctry,omitempty" json:"country,omitempty" binding:"regexp=^[A-Z]{2}$"`
	AdrLine     string     `xml:"AdrLine,omitempty" json:"addressLine,omitempty" binding:"max=70"`
}

type adressType struct {
	Cd    string      `xml:"Cd,omitempty" json:"code,omitempty" binding:"required_without=Prtry,omitempty,oneof=ADDR PBOX HOME BIZZ MLTO DLVY"`
	Prtry *adrTpPrtry `xml:"Prtry,omitempty" json:"proprietary,omitzero" binding:"required_without=Cd"`
}

type adrTpPrtry struct {
	Id      string `xml:"Id" json:"id" binding:"required,regexp=^[a-zA-Z0-9]{4}$"`
	Issr    string `xml:"Issr" json:"issuer" binding:"required,max=35"`
	SchmeNm string `xml:"SchmeNm,omitempty" json:"schemeName,omitempty" binding:"max=35"`
}

type partyId struct {
	OrgId  orgId  `xml:"OrgId,omitempty" json:"orgId,omitzero" binding:"required_without=PrvtId"`
	PrvtId prvtId `xml:"PrvtId,omitempty" json:"privateId,omitzero" binding:"required_without=OrgId"`
}

type orgId struct {
	AnyBIC string      `xml:"AnyBIC,omitempty" json:"anyBic,omitempty"`
	LEI    string      `xml:"LEI,omitempty" json:"lei,omitempty" binding:"regexp=^[A-Z0-9]{18}[0-9]{2}$"`
	Othr   []OrgIdOthr `xml:"Othr,omitempty" json:"other,omitzero"`
}

type OrgIdOthr struct {
	Id      string     `xml:"Id,omitempty" json:"id,omitempty"`
	Issr    string     `xml:"Issr,omitempty" json:"issuer,omitempty"`
	SchmeNm orgSchmeNm `xml:"SchmeNm,omitempty" json:"schemeName,omitzero"`
}

type orgSchmeNm struct {
	Cd    string `xml:"Cd,omitempty" json:"code,omitempty"`
	Prtry string `xml:"Prtry,omitempty" json:"proprietary,omitempty" binding:"max=35"`
}

type prvtId struct {
	DtAndPlcOfBirth dateAndPlaceOfBirth `xml:"DtAndPlcOfBirth,omitempty" json:"dateAndPlaceOfBirth,omitzero"`
	Othr            []PrvtIdOthr        `xml:"Othr,omitempty" json:"other,omitzero"`
}

type PrvtIdOthr struct {
	Id      string        `xml:"Id,omitempty" json:"id,omitempty"`
	Issr    string        `xml:"Issr,omitempty" json:"issuer,omitempty"`
	SchmeNm personSchmeNm `xml:"SchmeNm,omitempty" json:"schemeName,omitzero"`
}

type personSchmeNm struct {
	Cd    string `xml:"Cd,omitempty" json:"code,omitempty"`
	Prtry string `xml:"Prtry,omitempty" json:"proprietary,omitempty" binding:"max=35"`
}

type dateAndPlaceOfBirth struct {
	BirthDt     string `xml:"BirthDt,omitempty" json:"birthDate,omitempty"`
	PrvcOfBirth string `xml:"PrvcOfBirth,omitempty" json:"provinceOfBirth,omitempty"`
	CityOfBirth string `xml:"CityOfBirth,omitempty" json:"cityOfBirth,omitempty"`
	CtryOfBirth string `xml:"CtryOfBirth,omitempty" json:"countryOfBirth,omitempty" binding:"regexp=^[A-Z]{2}$"`
}

type ctctDtls struct {
	NmPrfx    string        `xml:"NmPrfx,omitempty" json:"namePrefix,omitempty"`
	Nm        string        `xml:"Nm,omitempty" json:"name,omitempty"`
	PhneNb    string        `xml:"PhneNb,omitempty" json:"phoneNumber,omitempty" binding:"e164PhoneNumbers"`
	MobNb     string        `xml:"MobNb,omitempty" json:"mobileNumber,omitempty" binding:"e164PhoneNumbers"`
	FaxNb     string        `xml:"FaxNb,omitempty" json:"faxNumber,omitempty" binding:"e164PhoneNumbers"`
	URLAdr    string        `xml:"URLAdr,omitempty" json:"urlAddress,omitempty"`
	EmailAdr  string        `xml:"EmailAdr,omitempty" json:"emailAddress,omitempty" binding:"email"`
	EmailPurp string        `xml:"EmailPurp,omitempty" json:"emailPurpose,omitempty"`
	JobTitl   string        `xml:"JobTitl,omitempty" json:"jobTitle,omitempty"`
	Rspnsblty string        `xml:"Rspnsblty,omitempty" json:"responsibility,omitempty"`
	Dept      string        `xml:"Dept,omitempty" json:"department,omitempty"`
	Othr      []ContactOthr `xml:"Othr,omitempty" json:"other,omitzero"`
	PrefrdMtd string        `xml:"PrefrdMtd,omitempty" json:"preferredMethod,omitempty"`
}

type ContactOthr struct {
	Id      string `xml:"Id,omitempty" json:"id,omitempty"`
	ChanlTp string `xml:"ChanlTp,omitempty" json:"channelType,omitempty"`
}

type account struct {
	Id   accountId    `xml:"Id" json:"accountId"`
	Tp   accountType  `xml:"Tp,omitempty" json:"type,omitzero"`
	Ccy  string       `xml:"Ccy,omitempty" json:"currency,omitempty" binding:"regexp=^[A-Z]{3}$"`
	Nm   string       `xml:"Nm,omitempty" json:"name,omitempty"`
	Prxy accountProxy `xml:"Prxy,omitempty" json:"proxy,omitzero"`
}

type accountId struct {
	IBAN string      `xml:"IBAN,omitempty" json:"iban,omitempty" binding:"required_without=Othr"`
	Othr accountOthr `xml:"Othr,omitempty" json:"other,omitzero" binding:"required_without=IBAN"`
}

type accountOthr struct {
	Id      string         `xml:"Id,omitempty" json:"id,omitempty" binding:"max=34"`
	Issr    string         `xml:"Issr,omitempty" json:"issuer,omitempty" binding:"max=35"`
	SchmeNm accountSchmeNm `xml:"SchmeNm,omitempty" json:"schemeName,omitzero"`
}

type accountSchmeNm struct {
	Cd    string `xml:"Cd,omitempty" json:"code,omitempty" binding:"required_without=Prtry"`
	Prtry string `xml:"Prtry,omitempty" json:"proprietary,omitempty" binding:"required_without=Cd,max=35"`
}

type accountType struct {
	Cd    string `xml:"Cd,omitempty" json:"code,omitempty"`
	Prtry string `xml:"Prtry,omitempty" json:"proprietary,omitempty" binding:"max=35"`
}

type accountProxy struct {
	Tp string `xml:"Tp,omitempty" json:"type,omitempty"`
	Id string `xml:"Id,omitempty" json:"id,omitempty"`
}

type instrForCdtrAgt struct {
	Cd       string `xml:"Cd,omitempty" json:"code,omitempty" binding:"omitempty,min=1,max=4"`
	InstrInf string `xml:"InstrInf,omitempty" json:"instructionInfo,omitempty" binding:"max=140"`
}

type instrForNxtAgt struct {
	Cd       string `xml:"Cd,omitempty" json:"code,omitempty" binding:"omitempty,oneof=PHOA TELA"`
	InstrInf string `xml:"InstrInf,omitempty" json:"instructionInfo,omitempty" binding:"max=140"`
}

type purp struct {
	Cd    string `xml:"Cd,omitempty" json:"code,omitempty" binding:"omitempty,min=1,max=4,required_without=Prtry"`
	Prtry string `xml:"Prtry,omitempty" json:"proprietary,omitempty" binding:"required_without=Cd,max=35"`
}

type rgltryRptg struct {
	DbtCdtRptgInd string            `xml:"DbtCdtRptgInd,omitempty" json:"debitCreditReportingIndicator,omitempty" binding:"omitempty,oneof=CRED DEBT BOTH"`
	Authrty       rgltryRptgAuthrty `xml:"Authrty,omitempty" json:"authority,omitzero"`
	Dtls          []rgltryRptgDtls  `xml:"Dtls,omitempty" json:"details,omitempty"`
}

type rgltryRptgAuthrty struct {
	Nm   string `xml:"Nm,omitempty" json:"name,omitempty" binding:"max=140"`
	Ctry string `xml:"Ctry,omitempty" json:"country,omitempty" binding:"regexp=^[A-Z]{2}$"`
}

type rgltryRptgDtls struct {
	Tp   string   `xml:"Tp,omitempty" json:"type,omitempty" binding:"max=35"`
	Dt   string   `xml:"Dt,omitempty" json:"date,omitempty"`
	Ctry string   `xml:"Ctry,omitempty" json:"country,omitempty" binding:"regexp=^[A-Z]{2}$"`
	Cd   string   `xml:"Cd,omitempty" json:"code,omitempty" binding:"omitempty,min=1,max=10"`
	Amt  amount   `xml:"Amt,omitempty" json:"amount,omitzero"`
	Inf  []string `xml:"Inf,omitempty" json:"information,omitempty"`
}

type tax struct {
	Cdtr            taxParty1 `xml:"Cdtr,omitempty" json:"creditor,omitzero"`
	Dbtr            taxParty2 `xml:"Dbtr,omitempty" json:"debtor,omitzero"`
	UltmtDbtr       taxParty2 `xml:"UltmtDbtr,omitempty" json:"ultimateDebtor,omitzero"`
	AdmstnZone      string    `xml:"AdmstnZone,omitempty" json:"administrationZone,omitempty"`
	RefNb           string    `xml:"RefNb,omitempty" json:"referenceNumber,omitempty"`
	Mtd             string    `xml:"Mtd,omitempty" json:"method,omitempty"`
	TtlTaxblBaseAmt amount    `xml:"TtlTaxblBaseAmt,omitempty" json:"totalTaxableBaseAmount,omitzero"`
	TtlTaxAmt       amount    `xml:"TtlTaxAmt,omitempty" json:"totalTaxAmount,omitzero"`
	Dt              string    `xml:"Dt,omitempty" json:"date,omitempty"`
	SeqNb           int64     `xml:"SeqNb,omitempty" json:"sequenceNumber,omitempty"`
	Rcrd            []taxRcrd `xml:"Rcrd,omitempty" json:"records,omitempty"`
}

type taxParty1 struct {
	TaxId  string `xml:"TaxId,omitempty" json:"taxId,omitempty"`
	RegnId string `xml:"RegnId,omitempty" json:"registrationId,omitempty"`
	TaxTp  string `xml:"TaxTp,omitempty" json:"taxType,omitempty" binding:"max=35"`
}

type taxParty2 struct {
	TaxId   string          `xml:"TaxId,omitempty" json:"taxId,omitempty"`
	RegnId  string          `xml:"RegnId,omitempty" json:"registrationId,omitempty"`
	TaxTp   string          `xml:"TaxTp,omitempty" json:"taxType,omitempty"`
	Authstn taxPartyAuthstn `xml:"Authstn,omitempty" json:"authority,omitzero"`
}

type taxPartyAuthstn struct {
	Titl string `xml:"Titl,omitempty" json:"title,omitempty"`
	Nm   string `xml:"Nm,omitempty" json:"name,omitempty"`
}

type taxRcrd struct {
	Tp       string `xml:"Tp,omitempty" json:"type,omitempty"`
	Ctgy     string `xml:"Ctgy,omitempty" json:"category,omitempty"`
	CtgyDtls string `xml:"CtgyDtls,omitempty" json:"categoryDetails,omitempty"`
	DbtrSts  string `xml:"DbtrSts,omitempty" json:"debtorStatus,omitempty"`
	CertId   string `xml:"CertId,omitempty" json:"certificateId,omitempty"`
	FrmsCd   string `xml:"FrmsCd,omitempty" json:"formsCode,omitempty"`
	Prd      taxPrd `xml:"Prd,omitempty" json:"period,omitzero"`
	TaxAmt   taxAmt `xml:"TaxAmt,omitempty" json:"taxAmount,omitzero"`
	AddtlInf string `xml:"AddtlInf,omitempty" json:"additionalInfo,omitempty"`
}

type taxPrd struct {
	Yr     int16   `xml:"Yr,omitempty" json:"year,omitempty"`
	Tp     string  `xml:"Tp,omitempty" json:"type,omitempty"`
	FrToDt datePrd `xml:"FrToDt,omitempty" json:"fromToDate,omitzero"`
}

type datePrd struct {
	FrDt string `xml:"FrDt,omitempty" json:"fromDate,omitempty"`
	ToDt string `xml:"ToDt,omitempty" json:"toDate,omitempty"`
}

type taxAmt struct {
	Rate         float64    `xml:"Rate,omitempty" json:"rate,omitempty"`
	TaxblBaseAmt amount     `xml:"TaxblBaseAmt,omitempty" json:"taxableBaseAmount,omitzero"`
	TtlAmt       amount     `xml:"TtlAmt,omitempty" json:"totalAmount,omitzero"`
	Dtls         taxAmtDtls `xml:"Dtls,omitempty" json:"details,omitzero"`
}

type taxAmtDtls struct {
	Prd taxPrd `xml:"Prd,omitempty" json:"period,omitzero"`
	Amt amount `xml:"Amt,omitempty" json:"amount,omitzero"`
}

type rltdRmtInf struct {
	RmtId       string        `xml:"RmtId,omitempty" json:"remittanceId,omitempty"`
	RmtLctnDtls []rmtLctnDtls `xml:"RmtLctnDtls,omitempty" json:"remittanceLocationDetails,omitempty"`
}

type rmtLctnDtls struct {
	Mtd        string   `xml:"Mtd,omitempty" json:"method,omitempty"`
	ElctrncAdr string   `xml:"ElctrncAdr,omitempty" json:"electronicAddress,omitempty"`
	PstlAdr    nmAndAdr `xml:"PstlAdr,omitempty" json:"postalAddress,omitzero"`
}

type nmAndAdr struct {
	Nm  string  `xml:"Nm,omitempty" json:"name,omitempty"`
	Adr pstlAdr `xml:"Adr,omitempty" json:"address,omitzero"`
}

type rmtInf struct {
	Ustrd []string     `xml:"Ustrd,omitempty" json:"ustrdRemittanceInfo,omitempty"`
	Strd  []rmtInfStrd `xml:"Strd,omitempty" json:"strdRemittanceInfo,omitempty"`
}

type rmtInfStrd struct {
	RfrdDocInf  []rfrdDocInf   `xml:"RfrdDocInf,omitempty" json:"referredDocumentInfo,omitempty"`
	RfrdDocAmt  lineDtlsAmount `xml:"RfrdDocAmt,omitempty" json:"referredDocumentAmount,omitzero"`
	CdtrRefInf  cdtrRefInf     `xml:"CdtrRefInf,omitempty" json:"creditorReferenceInfo,omitzero"`
	Invcr       party          `xml:"Invcr,omitempty" json:"invoicer,omitzero"`
	Invcee      party          `xml:"Invcee,omitempty" json:"invoicee,omitzero"`
	TaxRmt      tax            `xml:"TaxRmt,omitempty" json:"taxRemittance,omitzero"`
	GrnshmtRmt  grnshmtRmt     `xml:"GrnshmtRmt,omitempty" json:"garnishmentRemittance,omitzero"`
	AddtlRmtInf []string       `xml:"AddtlRmtInf,omitempty" json:"additionalRemittanceInfo,omitempty"`
}

type rfrdDocInf struct {
	Tp       rfrdDocInfTp `xml:"Tp,omitempty" json:"type,omitzero"`
	Nb       string       `xml:"Nb,omitempty" json:"number,omitempty"`
	RltdDt   dtAndTp      `xml:"RltdDt,omitempty" json:"relatedDate,omitzero"`
	LineDtls []lineDtls   `xml:"LineDtls,omitempty" json:"lineDetails,omitempty"`
}

type rfrdDocInfTp struct {
	CdOrPrtry cdOrPrtry `xml:"CdOrPrtry,omitempty" json:"codeOrProprietary,omitzero"`
	Issr      string    `xml:"Issr,omitempty" json:"issuer,omitempty"`
}

type cdOrPrtry struct {
	Cd    string `xml:"Cd,omitempty" json:"code,omitempty"`
	Prtry string `xml:"Prtry,omitempty" json:"proprietary,omitempty" binding:"max=35"`
}

type dtAndTp struct {
	Tp string `xml:"Tp,omitempty" json:"type,omitempty"`
	Dt string `xml:"Dt,omitempty" json:"date,omitempty"`
}

type lineDtls struct {
	Id   []lineDtlsId   `xml:"Id,omitempty" json:"identifiers,omitempty"`
	Desc string         `xml:"Desc,omitempty" json:"description,omitempty"`
	Amt  lineDtlsAmount `xml:"Amt,omitempty" json:"amount,omitzero"`
}

type lineDtlsId struct {
	Tp     lineDtlsIdTp `xml:"Tp,omitempty" json:"type,omitzero"`
	Nb     string       `xml:"Nb,omitempty" json:"number,omitempty"`
	RltdDt string       `xml:"RltdDt,omitempty" json:"relatedDate,omitempty"`
}

type lineDtlsIdTp struct {
	CdOrPrtry cdOrPrtry `xml:"CdOrPrtry,omitempty" json:"codeOrProprietary,omitzero"`
	Issr      string    `xml:"Issr,omitempty" json:"issuer,omitempty"`
}

type lineDtlsAmount struct {
	RmtAmtAndTp       []lineDtlsRmtAmtAndTp       `xml:"RmtAmtAndTp,omitempty" json:"remittanceAmountAndType,omitempty"`
	AdjstmntAmtAndRsn []lineDtlsAdjstmntAmtAndRsn `xml:"AdjstmntAmtAndRsn,omitempty" json:"adjustmentAmountAndReason,omitempty"`
}

type lineDtlsRmtAmtAndTp struct {
	Tp  lineDtlsCdOrPrtry `xml:"Tp,omitempty" json:"type,omitzero"`
	Amt amount            `xml:"Amt,omitempty" json:"amount,omitzero"`
}

type lineDtlsCdOrPrtry struct {
	Cd    string `xml:"Cd,omitempty" json:"code,omitempty"`
	Prtry string `xml:"Prtry,omitempty" json:"proprietary,omitempty" binding:"max=35"`
}

type lineDtlsAdjstmntAmtAndRsn struct {
	Amt       amount `xml:"Amt,omitempty" json:"amount,omitzero"`
	CdtDbtInd string `xml:"CdtDbtInd,omitempty" json:"creditDebitIndicator,omitempty"`
	Rsn       string `xml:"Rsn,omitempty" json:"reason,omitempty"`
	AddtlInf  string `xml:"AddtlInf,omitempty" json:"additionalInfo,omitempty"`
}

type cdtrRefInf struct {
	Tp  cdtrRefInfTp `xml:"Tp,omitempty" json:"type,omitzero"`
	Ref string       `xml:"Ref,omitempty" json:"reference,omitempty"`
}

type cdtrRefInfTp struct {
	CdOrPrtry cdOrPrtry `xml:"CdOrPrtry,omitempty" json:"codeOrProprietary,omitzero"`
	Issr      string    `xml:"Issr,omitempty" json:"issuer,omitempty"`
}

type grnshmtRmt struct {
	Tp                grnshmtRmtTp `xml:"Tp,omitempty" json:"type,omitzero"`
	Grnshee           party        `xml:"Grnshee,omitempty" json:"garnishee,omitzero"`
	GrnshmtAdmstr     party        `xml:"GrnshmtAdmstr,omitempty" json:"garnishmentAdministrator,omitzero"`
	RefNb             string       `xml:"RefNb,omitempty" json:"referenceNumber,omitempty"`
	Dt                string       `xml:"Dt,omitempty" json:"date,omitempty"`
	RmtdAmt           amount       `xml:"RmtdAmt,omitempty" json:"remittedAmount,omitzero"`
	FmlyMdclInsrncInd bool         `xml:"FmlyMdclInsrncInd,omitempty" json:"familyMedicalInsuranceIndicator,omitempty"`
	MplyeeTermntnInd  bool         `xml:"MplyeeTermntnInd,omitempty" json:"employeeTerminationIndicator,omitempty"`
}

type grnshmtRmtTp struct {
	CdOrPrtry cdOrPrtry `xml:"CdOrPrtry,omitempty" json:"codeOrProprietary,omitzero"`
	Issr      string    `xml:"Issr,omitempty" json:"issuer,omitempty"`
}

type splmtryData struct {
	PlcAndNm string `xml:"PlcAndNm,omitempty" json:"placeAndName,omitempty" binding:"max=350"`
	Envlp    []byte `xml:"Envlp" json:"envelope" binding:"required"`
}
