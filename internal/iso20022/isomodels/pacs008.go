package isomodels

import "encoding/xml"

type Pacs008 struct {
	XMLName           xml.Name           `xml:"Document"`
	Xmlns             string             `xml:"xmlns,attr"`
	SchemaLocation    string             `xml:"schemaLocation,attr"`
	XmlnsXsi          string             `xml:"xsi,attr"`
	FIToFICstmrCdtTrf *FIToFICstmrCdtTrf `xml:"FIToFICstmrCdtTrf" json:"fiToFiCustomerCreditTransfer" binding:"required"`
}

type FIToFICstmrCdtTrf struct {
	GrpHdr      *GrpHdrPacs008 `xml:"GrpHdr" json:"groupHeader" binding:"required"`
	CdtTrfTxInf []*CdtTrfTxInf `xml:"CdtTrfTxInf" json:"creditTransferTransaction" binding:"required,min=1,dive"`
	SplmtryData []*SplmtryData `xml:"SplmtryData,omitempty" json:"supplementaryData,omitempty" binding:"omitempty,dive"`
}

type GrpHdrPacs008 struct {
	MsgId             string           `xml:"MsgId" json:"messageId" binding:"required,max=35"`
	CreDtTm           string           `xml:"CreDtTm" json:"creationDateTime" binding:"required,isoDateTime"`
	NbOfTxs           string           `xml:"NbOfTxs" json:"numberOfTxs" binding:"required,numeric,omitempty,min=1,max=15"`
	SttlmInf          *SttlmInf        `xml:"SttlmInf" json:"settlementInfo" binding:"required"`
	XpryDtTm          string           `xml:"XpryDtTm,omitempty" json:"expiryDateTime,omitempty" binding:"omitempty,isoDateTime"`
	BtchBookg         bool             `xml:"BtchBookg,omitempty" json:"batchBooking,omitempty"`
	CtrlSum           float64          `xml:"CtrlSum,omitempty" json:"controlSum,omitempty"`
	TtlIntrBkSttlmAmt *Amount          `xml:"TtlIntrBkSttlmAmt,omitempty" json:"totalSettlementAmount,omitzero"`
	IntrBkSttlmDt     string           `xml:"IntrBkSttlmDt,omitempty" json:"settlementDate,omitempty" binding:"isoDate"`
	PmtTpInf          *PmtTpInfPacs008 `xml:"PmtTpInf,omitempty" json:"paymentTypeInfo,omitzero"`
	InstgAgt          *Agent           `xml:"InstgAgt,omitempty" json:"instructingAgent,omitzero"`
	InstdAgt          *Agent           `xml:"InstdAgt,omitempty" json:"instructedAgent,omitzero"`
}

type SttlmInf struct {
	SttlmMtd             string          `xml:"SttlmMtd" json:"settlementMethod"`
	SttlmAcct            *Account        `xml:"SttlmAcct" json:"settlementAccount"`
	ClrSys               *SttlmInfClrSys `xml:"ClrSys" json:"clearinSystem"`
	InstgRmbrsmntAgt     *Agent          `xml:"InstgRmbrsmntAgt" json:"InstructingReimbursementAgent"`
	InstgRmbrsmntAgtAcct *Account        `xml:"InstgRmbrsmntAgtAcct" json:"InstructingReimbursementAgentAccount"`
	InstdRmbrsmntAgt     *Agent          `xml:"InstdRmbrsmntAgt" json:"InstructedReimbursementAgent"`
	InstdRmbrsmntAgtAcct *Account        `xml:"InstdRmbrsmntAgtAcct" json:"InstructedReimbursementAgentAccount"`
	ThrdRmbrsmntAgt      *Agent          `xml:"ThrdRmbrsmntAgt" json:"ThirdReimbursementAgent"`
	ThrdRmbrsmntAgtAcct  *Account        `xml:"ThrdRmbrsmntAgtAcct" json:"ThirdReimbursementAgentAccount"`
}

type SttlmInfClrSys struct {
	Cd    string `xml:"Cd" json:"code"`
	Prtry string `xml:"Prtry" json:"proprietary"`
}

type PmtTpInfPacs008 struct {
	InstrPrty string     `xml:"InstrPrty,omitempty" json:"instructionPriority,omitempty" binding:"omitempty,oneof=HIGH NORM"`
	ClrChanl  string     `xml:"ClrChanl,omitempty" json:"clearingChannel,omitempty" binding:"omitempty,oneof=RTGS RTNS MPNS BOOK"`
	SvcLvl    []*SvcLvl  `xml:"SvcLvl,omitempty" json:"serviceLevel,omitempty" binding:"omitempty,dive"`
	LclInstrm *LclInstrm `xml:"LclInstrm,omitempty" json:"localInstrument,omitzero"`
	CtgyPurp  *CtgyPurp  `xml:"CtgyPurp,omitempty" json:"categoryPurpose,omitzero"`
}

type SvcLvl struct {
	Cd    string `xml:"Cd,omitempty" json:"code,omitempty" binding:"omitempty,min=1,max=4,required_without=Prtry"`
	Prtry string `xml:"Prtry,omitempty" json:"proprietary,omitempty" binding:"max=35,required_without=Cd"`
}

type LclInstrm struct {
	Cd    string `xml:"Cd,omitempty" json:"code,omitempty" binding:"omitempty,min=1,max=35,required_without=Prtry"`
	Prtry string `xml:"Prtry,omitempty" json:"proprietary,omitempty" binding:"max=35,required_without=Cd"`
}

type CtgyPurp struct {
	Cd    string `xml:"Cd,omitempty" json:"code,omitempty" binding:"omitempty,min=1,max=4,required_without=Prtry"`
	Prtry string `xml:"Prtry,omitempty" json:"proprietary,omitempty" binding:"max=35,required_without=Cd"`
}

type Agent struct {
	FinInstnId *FinInstnId `xml:"FinInstnId" json:"financialInstitutionId" binding:"required"`
	BrnchId    *BrnchId    `xml:"BrnchId,omitempty" json:"branchId,omitzero"`
}

type FinInstnId struct {
	BICFI       string        `xml:"BICFI,omitempty" json:"bicfi,omitempty" binding:"regexp=^[A-Z0-9]{4}[A-Z]{2}[A-Z0-9]{2}([A-Z0-9]{3})?$"`
	ClrSysMmbId *ClrSysMmbId  `xml:"ClrSysMmbId,omitempty" json:"clearingSystemMemberId,omitzero"`
	LEI         string        `xml:"LEI,omitempty" json:"lei,omitempty" binding:"omitempty,regexp=^[A-Z0-9]{18}[0-9]{2}$"`
	Nm          string        `xml:"Nm,omitempty" json:"name,omitempty" binding:"max=140"`
	PstlAdr     *PstlAdr      `xml:"PstlAdr,omitempty" json:"postalAddress,omitzero"`
	Othr        *FinInstnOthr `xml:"Othr,omitempty" json:"other,omitzero"`
}

type ClrSysMmbId struct {
	ClrSysId *ClrSysId `xml:"ClrSysId,omitempty" json:"clearingSystemId,omitzero"`
	MmbId    string    `xml:"MmbId,omitempty" json:"memberId,omitempty" binding:"max=35"`
}

type ClrSysId struct {
	Cd    string `xml:"Cd,omitempty" json:"code,omitempty" binding:"omitempty,min=1,max=5"`
	Prtry string `xml:"Prtry,omitempty" json:"proprietary,omitempty" binding:"max=35"`
}

type BrnchId struct {
	Id      string   `xml:"Id,omitempty" json:"id,omitempty" binding:"max=35"`
	LEI     string   `xml:"LEI,omitempty" json:"lei,omitempty" binding:"regexp=^[A-Z0-9]{18}[0-9]{2}$"`
	Nm      string   `xml:"Nm,omitempty" json:"name,omitempty" binding:"max=140"`
	PstlAdr *PstlAdr `xml:"PstlAdr,omitempty" json:"postalAddress,omitzero"`
}

type CdtTrfTxInf struct {
	PmtId             *PmtId             `xml:"PmtId" json:"paymentId" binding:"required"`
	PmtTpInf          *PmtTpInfPacs008   `xml:"PmtTpInf,omitempty" json:"paymentTypeInfo,omitzero"`
	IntrBkSttlmAmt    *Amount            `xml:"IntrBkSttlmAmt" json:"settlementAmount" binding:"required"`
	IntrBkSttlmDt     string             `xml:"IntrBkSttlmDt" json:"settlementDate" binding:"required,isoDate"`
	SttlmPrty         string             `xml:"SttlmPrty,omitempty" json:"settlementPriority,omitempty"`
	SttlmTmIndctn     *SttlmTmIndctn     `xml:"SttlmTmIndctn,omitempty" json:"settlementTimeIndication,omitzero"`
	SttlmTmReq        *SttlmTmReq        `xml:"SttlmTmReq,omitempty" json:"settlementTimeRequest,omitzero"`
	AddtlDtTm         *AddtlDtTm         `xml:"AddtlDtTm,omitempty" json:"additionalDateTime,omitzero"`
	InstdAmt          *Amount            `xml:"InstdAmt,omitempty" json:"instructedAmount,omitzero"`
	XchgRate          float64            `xml:"XchgRate,omitempty" json:"exchangeRate,omitempty"`
	AgrdRate          *AgrdRate          `xml:"AgrdRate,omitempty" json:"agreedRate,omitzero"`
	ChrgBr            string             `xml:"ChrgBr,omitempty" json:"chargeBearer,omitempty" binding:"omitempty,oneof=DEBT CRED SHAR SLEV"`
	ChrgsInf          []*ChrgsInf        `xml:"ChrgsInf,omitempty" json:"chargesInfo,omitempty" binding:"dive"`
	MndtRltdInf       *CdtTrfMndt        `xml:"MndtRltdInf,omitempty" json:"mandateRelatedInfo,omitzero"`
	PmtSgntr          *PmtSgntr          `xml:"PmtSgntr,omitempty" json:"paymentSignature,omitzero"`
	PrvsInstgAgt1     *Agent             `xml:"PrvsInstgAgt1,omitempty" json:"pia1,omitzero"`
	PrvsInstgAgt1Acct *Account           `xml:"PrvsInstgAgt1Acct,omitempty" json:"pia1Account,omitzero"`
	PrvsInstgAgt2     *Agent             `xml:"PrvsInstgAgt2,omitempty" json:"pia2,omitzero"`
	PrvsInstgAgt2Acct *Account           `xml:"PrvsInstgAgt2Acct,omitempty" json:"pia2Account,omitzero"`
	PrvsInstgAgt3     *Agent             `xml:"PrvsInstgAgt3,omitempty" json:"pia3,omitzero"`
	PrvsInstgAgt3Acct *Account           `xml:"PrvsInstgAgt3Acct,omitempty" json:"pia3Account,omitzero"`
	InstgAgt          *Agent             `xml:"InstgAgt,omitempty" json:"instructingAgent,omitzero"`
	InstdAgt          *Agent             `xml:"InstdAgt,omitempty" json:"instructedAgent,omitzero"`
	IntrmyAgt1        *Agent             `xml:"IntrmyAgt1,omitempty" json:"ia1,omitzero"`
	IntrmyAgt1Acct    *Account           `xml:"IntrmyAgt1Acct,omitempty" json:"ia1Account,omitzero"`
	IntrmyAgt2        *Agent             `xml:"IntrmyAgt2,omitempty" json:"ia2,omitzero"`
	IntrmyAgt2Acct    *Account           `xml:"IntrmyAgt2Acct,omitempty" json:"ia2Account,omitzero"`
	IntrmyAgt3        *Agent             `xml:"IntrmyAgt3,omitempty" json:"ia3,omitzero"`
	IntrmyAgt3Acct    *Account           `xml:"IntrmyAgt3Acct,omitempty" json:"ia3Account,omitzero"`
	UltmtDbtr         *Party             `xml:"UltmtDbtr,omitempty" json:"ultimateDebtor,omitzero"`
	InitgPty          *Party             `xml:"InitgPty,omitempty" json:"initiatingParty,omitzero"`
	Dbtr              *Party             `xml:"Dbtr" json:"debtor"`
	DbtrAcct          *Account           `xml:"DbtrAcct" json:"debtorAccount"`
	DbtrAgt           *Agent             `xml:"DbtrAgt,omitempty" json:"debtorAgent,omitzero"`
	DbtrAgtAcct       *Account           `xml:"DbtrAgtAcct,omitempty" json:"debtorAgentAccount,omitzero"`
	CdtrAgt           *Agent             `xml:"CdtrAgt,omitempty" json:"creditorAgent,omitzero"`
	CdtrAgtAcct       *Account           `xml:"CdtrAgtAcct,omitempty" json:"creditorAgentAccount,omitzero"`
	Cdtr              *Party             `xml:"Cdtr" json:"creditor"`
	CdtrAcct          *Account           `xml:"CdtrAcct" json:"creditorAccount"`
	UltmtCdtr         *Party             `xml:"UltmtCdtr,omitempty" json:"ultimateCreditor,omitzero"`
	InstrForCdtrAgt   []*InstrForCdtrAgt `xml:"InstrForCdtrAgt,omitempty" json:"instructionsForCreditorAgent,omitempty" binding:"dive"`
	InstrForNxtAgt    []*InstrForNxtAgt  `xml:"InstrForNxtAgt,omitempty" json:"instructionsForNextAgent,omitempty" binding:"dive"`
	Purp              *Purp              `xml:"Purp,omitempty" json:"purpose,omitzero"`
	RgltryRptg        []*RgltryRptg      `xml:"RgltryRptg,omitempty" json:"regulatoryReporting,omitempty" binding:"max=10,dive"`
	Tax               *Tax               `xml:"Tax,omitempty" json:"tax,omitzero"`
	RltdRmtInf        []*RltdRmtInf      `xml:"RltdRmtInf,omitempty" json:"relatedRemittanceInfo,omitempty" binding:"dive"`
	RmtInf            *RmtInf            `xml:"RmtInf,omitempty" json:"remittanceInfo,omitzero"`
	SplmtryData       []*SplmtryData     `xml:"SplmtryData,omitempty" json:"supplementaryData,omitzero" binding:"omitempty,dive"`
}

type PmtId struct {
	InstrId    string `xml:"InstrId" json:"instructionId" binding:"max=35"`
	EndToEndId string `xml:"EndToEndId" json:"endToEndId" binding:"required,max=35"`
	TxId       string `xml:"TxId" json:"transactionId" binding:"max=35"`
	UETR       string `xml:"UETR" json:"uetr" binding:"uuid4"`
	ClrSysRef  string `xml:"ClrSysRef,omitempty" json:"clearingSystemReference,omitempty" binding:"max=35"`
}

type Amount struct {
	Value    float64 `xml:",chardata" json:"amountValue" binding:"required"`
	Currency string  `xml:"Ccy,attr" json:"amountCurrency" binding:"required,regexp=^[A-Z]{3}$"`
}

type SttlmTmIndctn struct {
	DbtDtTm string `xml:"DbtDtTm,omitempty" json:"debitDateTime,omitempty" binding:"isoDateTime"`
	CdtDtTm string `xml:"CdtDtTm,omitempty" json:"creditDateTime,omitempty" binding:"isoDateTime"`
}

type SttlmTmReq struct {
	CLSTm  string `xml:"CLSTm,omitempty" json:"closeTime,omitempty"`
	TillTm string `xml:"TillTm,omitempty" json:"tillTime,omitempty"`
	FrTm   string `xml:"FrTm,omitempty" json:"fromTime,omitempty"`
	RjctTm string `xml:"RjctTm,omitempty" json:"rejectTime,omitempty"`
}

type AddtlDtTm struct {
	AccptncDtTm     string `xml:"AccptncDtTm,omitempty" json:"acceptanceDateTime,omitempty" binding:"isoDateTime"`
	PoolgAdjstmntDt string `xml:"PoolgAdjstmntDt,omitempty" json:"poolingAdjustmentDate,omitempty" binding:"isoDate"`
	XpryDtTm        string `xml:"XpryDtTm,omitempty" json:"expiryDateTime,omitempty" binding:"isoDateTime"`
}

type AgrdRate struct {
	UnitCcy         string  `xml:"UnitCcy,omitempty" json:"unitCurrency,omitempty" binding:"regexp=^[A-Z]{3}$"`
	QtdCcy          string  `xml:"QtdCcy,omitempty" json:"quotedCurrency,omitempty" binding:"regexp=^[A-Z]{3}$"`
	PreAgrdXchgRate float64 `xml:"PreAgrdXchgRate" json:"preAgreedExchangeRate" binding:"required"`
	QtnDtTm         string  `xml:"QtnDtTm,omitempty" json:"quotationDateTime,omitempty" binding:"isoDateTime"`
	QtId            string  `xml:"QtId,omitempty" json:"quoteId,omitempty" binding:"uuid4"`
	FXAgt           Agent   `xml:"FXAgt,omitempty" json:"foreignExchangeAgent,omitzero"`
}

type ChrgsInf struct {
	Amt *Amount  `xml:"Amt" json:"amount" binding:"required"`
	Agt *Agent   `xml:"Agt" json:"agent" binding:"required"`
	Tp  *ChrgsTp `xml:"Tp,omitempty" json:"type,omitzero"`
}

type ChrgsTp struct {
	Cd    string        `xml:"Cd,omitempty" json:"code,omitempty"`
	Prtry *ChrgsTpPrtry `xml:"Prtry,omitempty" json:"proprietary,omitzero"`
}

type ChrgsTpPrtry struct {
	Id   string `xml:"Id,omitempty" json:"id,omitempty"`
	Issr string `xml:"Issr,omitempty" json:"issuer,omitempty"`
}

type CdtTrfMndt struct {
	MndtId       string         `xml:"MndtId,omitempty" json:"mandateId,omitempty" binding:"max=35"`
	Tp           *CdtTrfMndtTp  `xml:"Tp,omitempty" json:"type,omitzero"`
	DtOfSgntr    string         `xml:"DtOfSgntr,omitempty" json:"dateOfSignature,omitempty" binding:"isoDate"`
	DtOfVrfctn   string         `xml:"DtOfVrfctn,omitempty" json:"dateOfVerification,omitempty" binding:"isoDateTime"`
	ElctrncSgntr []byte         `xml:"ElctrncSgntr,omitempty" json:"electronicSignature,omitempty"`
	FrstPmtDt    string         `xml:"FrstPmtDt,omitempty" json:"firstPaymentDate,omitempty" binding:"isoDate"`
	FnlPmtDt     string         `xml:"FnlPmtDt,omitempty" json:"finalPaymentDate,omitempty" binding:"isoDate"`
	Frqcy        *Frqcy         `xml:"Frqcy,omitempty" json:"frequency,omitzero"`
	Rsn          *CdtTrfMndtRsn `xml:"Rsn,omitempty" json:"reason,omitzero"`
}

type CdtTrfMndtTp struct {
	SvcLvl    []*SvcLvl             `xml:"SvcLvl,omitempty" json:"serviceLevel,omitempty" binding:"dive"`
	LclInstrm *LclInstrm            `xml:"LclInstrm,omitempty" json:"localInstrument,omitzero"`
	CtgyPurp  *CtgyPurp             `xml:"CtgyPurp,omitempty" json:"categoryPurpose,omitzero"`
	Clssfctn  *CdtTrfMndtTpClssfctn `xml:"Clssfctn,omitempty" json:"classification,omitzero"`
}

type CdtTrfMndtTpClssfctn struct {
	Cd    string `xml:"Cd,omitempty" json:"code,omitempty" binding:"required_without=Prtry,omitempty,oneof=FIXE USGB VARI"`
	Prtry string `xml:"Prtry,omitempty" json:"proprietary,omitempty" binding:"max=35,required_without=Cd"`
}

type Frqcy struct {
	Tp     string    `xml:"Tp,omitempty" json:"type,omitempty" binding:"omitempty,required_without_all=Prd PtInTm,omitempty,oneof=YEAR MNTH QURT MIAN WEEK DAIL ADHO INDA FRTN"`
	Prd    *FrqcyPrd `xml:"Prd,omitempty" json:"period,omitzero" binding:"omitempty,required_without_all=Tp PtInTm"`
	PtInTm *PtInTm   `xml:"PtInTm,omitempty" json:"pointInTime,omitzero" binding:"omitempty,required_without_all=Tp Prd"`
}

type FrqcyPrd struct {
	Tp        string  `xml:"Tp" json:"type" binding:"required,oneof=YEAR MNTH QURT MIAN WEEK DAIL ADHO INDA FRTN"`
	CntPerPrd float64 `xml:"CntPerPrd" json:"countPerPeriod" binding:"required"`
}

type PtInTm struct {
	Tp     string `xml:"Tp" json:"type" binding:"required,oneof=YEAR MNTH QURT MIAN WEEK DAIL ADHO INDA FRTN"`
	PtInTm string `xml:"PtInTm,omitempty" json:"pointInTime,omitempty" binding:"omitempty,regexp=^[0-9]{2}$"`
}

type CdtTrfMndtRsn struct {
	Cd    string `xml:"Cd,omitempty" json:"code,omitempty" binding:"omitempty,min=1,max=4,required_without=Prtry"`
	Prtry string `xml:"Prtry,omitempty" json:"proprietary,omitempty" binding:"max=35,required_without=Cd"`
}

type PmtSgntr struct {
	ILPV4 string `xml:"ILPV4,omitempty" json:"ilpv4,omitempty" binding:"hexadecimal,required_without=Sgntr"`
	Sgntr string `xml:"Sgntr,omitempty" json:"signature,omitempty" binding:"omitempty,required_without=ILPV4,regexp=^([0-9a-fA-F]{2}){32}$"`
}

type Party struct {
	Nm        string    `xml:"Nm,omitempty" json:"name,omitempty" binding:"required"`
	PstlAdr   *PstlAdr  `xml:"PstlAdr,omitempty" json:"postalAddress,omitzero"`
	Id        *PartyId  `xml:"Id,omitempty" json:"partyId,omitzero"`
	CtryOfRes string    `xml:"CtryOfRes,omitempty" json:"countryOfResidence,omitempty" binding:"omitempty,regexp=^[A-Z]{2}$"`
	CtctDtls  *CtctDtls `xml:"CtctDtls,omitempty" json:"contactDetails,omitzero"`
}

type FinInstnOthr struct {
	Id      string           `xml:"Id,omitempty" json:"id,omitempty" binding:"max=35"`
	Issr    string           `xml:"Issr,omitempty" json:"issuer,omitempty" binding:"max=35"`
	SchmeNm *FinInstnSchmeNm `xml:"SchmeNm,omitempty" json:"schemeName,omitzero"`
}

type FinInstnSchmeNm struct {
	Cd    string `xml:"Cd,omitempty" json:"code,omitempty" binding:"required_without=Prtry,max=4"`
	Prtry string `xml:"Prtry,omitempty" json:"proprietary,omitempty" binding:"required_without=Cd,max=35"`
}

type PstlAdr struct {
	AdrTp       *AdrTp `xml:"AdrTp,omitempty" json:"addressType,omitzero"`
	CareOf      string `xml:"CareOf,omitempty" json:"careOf,omitempty" binding:"max=140"`
	Dept        string `xml:"Dept,omitempty" json:"department,omitempty" binding:"max=70"`
	SubDept     string `xml:"SubDept,omitempty" json:"subDepartment,omitempty" binding:"max=70"`
	StrtNm      string `xml:"StrtNm,omitempty" json:"streetName,omitempty" binding:"max=140"`
	BldgNb      string `xml:"BldgNb,omitempty" json:"buildingNumber,omitempty" binding:"max=16"`
	BldgNm      string `xml:"BldgNm,omitempty" json:"buildingName,omitempty" binding:"max=140"`
	Flr         string `xml:"Flr,omitempty" json:"floor,omitempty" binding:"max=70"`
	UnitNb      string `xml:"UnitNb,omitempty" json:"unitNumber,omitempty" binding:"max=16"`
	PstBx       string `xml:"PstBx,omitempty" json:"postBox,omitempty" binding:"max=16"`
	Room        string `xml:"Room,omitempty" json:"room,omitempty" binding:"max=70"`
	PstCd       string `xml:"PstCd,omitempty" json:"postalCode,omitempty" binding:"max=16"`
	TwnNm       string `xml:"TwnNm,omitempty" json:"townName,omitempty" binding:"max=140"`
	TwnLctnNm   string `xml:"TwnLctnNm,omitempty" json:"townLocationName,omitempty" binding:"max=140"`
	DstrctNm    string `xml:"DstrctNm,omitempty" json:"districtName,omitempty" binding:"max=140"`
	CtrySubDvsn string `xml:"CtrySubDvsn,omitempty" json:"countrySubdivision,omitempty" binding:"max=35"`
	Ctry        string `xml:"Ctry,omitempty" json:"country,omitempty" binding:"omitempty,regexp=^[A-Z]{2}$"`
	AdrLine     string `xml:"AdrLine,omitempty" json:"addressLine,omitempty" binding:"max=70"`
}

type AdrTp struct {
	Cd    string      `xml:"Cd,omitempty" json:"code,omitempty" binding:"omitempty,required_without=Prtry,oneof=ADDR PBOX HOME BIZZ MLTO DLVY"`
	Prtry *AdrTpPrtry `xml:"Prtry,omitempty" json:"proprietary,omitzero" binding:"omitempty,required_without=Cd"`
}

type AdrTpPrtry struct {
	Id      string `xml:"Id" json:"id" binding:"required,regexp=^[a-zA-Z0-9]{4}$"`
	Issr    string `xml:"Issr" json:"issuer" binding:"required,max=35"`
	SchmeNm string `xml:"SchmeNm,omitempty" json:"schemeName,omitempty" binding:"max=35"`
}

type PartyId struct {
	OrgId  *OrgId  `xml:"OrgId,omitempty" json:"orgId,omitzero" binding:"required_without=PrvtId"`
	PrvtId *PrvtId `xml:"PrvtId,omitempty" json:"privateId,omitzero" binding:"required_without=OrgId"`
}

type OrgId struct {
	AnyBIC string       `xml:"AnyBIC,omitempty" json:"anyBic,omitempty"`
	LEI    string       `xml:"LEI,omitempty" json:"lei,omitempty" binding:"omitempty,regexp=^[A-Z0-9]{18}[0-9]{2}$"`
	Othr   []*OrgIdOthr `xml:"Othr,omitempty" json:"other,omitzero" binding:"dive"`
}

type OrgIdOthr struct {
	Id      string      `xml:"Id,omitempty" json:"id,omitempty"`
	Issr    string      `xml:"Issr,omitempty" json:"issuer,omitempty"`
	SchmeNm *OrgSchmeNm `xml:"SchmeNm,omitempty" json:"schemeName,omitzero"`
}

type OrgSchmeNm struct {
	Cd    string `xml:"Cd,omitempty" json:"code,omitempty"`
	Prtry string `xml:"Prtry,omitempty" json:"proprietary,omitempty" binding:"max=35"`
}

type PrvtId struct {
	DtAndPlcOfBirth *DateAndPlaceOfBirth `xml:"DtAndPlcOfBirth,omitempty" json:"dateAndPlaceOfBirth,omitzero" binding:"omitempty"`
	Othr            []*PrvtIdOthr        `xml:"Othr,omitempty" json:"other,omitzero" binding:"dive"`
}

type DateAndPlaceOfBirth struct {
	BirthDt     string `xml:"BirthDt" json:"birthDate" binding:"required,isoDate"`
	PrvcOfBirth string `xml:"PrvcOfBirth,omitempty" json:"provinceOfBirth,omitempty"`
	CityOfBirth string `xml:"CityOfBirth,omitempty" json:"cityOfBirth,omitempty"`
	CtryOfBirth string `xml:"CtryOfBirth,omitempty" json:"countryOfBirth,omitempty" binding:"omitempty,regexp=^[A-Z]{2}$"`
}

type PrvtIdOthr struct {
	Id      string         `xml:"Id,omitempty" json:"id,omitempty"`
	Issr    string         `xml:"Issr,omitempty" json:"issuer,omitempty"`
	SchmeNm *PersonSchmeNm `xml:"SchmeNm,omitempty" json:"schemeName,omitzero"`
}

type PersonSchmeNm struct {
	Cd    string `xml:"Cd,omitempty" json:"code,omitempty"`
	Prtry string `xml:"Prtry,omitempty" json:"proprietary,omitempty" binding:"max=35"`
}

type CtctDtls struct {
	NmPrfx    string         `xml:"NmPrfx,omitempty" json:"namePrefix,omitempty" binding:"omitempty,oneof=DOCT MADM MISS MIST MIKS"`
	Nm        string         `xml:"Nm,omitempty" json:"name,omitempty" binding:"max=140"`
	PhneNb    string         `xml:"PhneNb,omitempty" json:"phoneNumber,omitempty" binding:"e164PhoneNumbers"`
	MobNb     string         `xml:"MobNb,omitempty" json:"mobileNumber,omitempty" binding:"e164PhoneNumbers"`
	FaxNb     string         `xml:"FaxNb,omitempty" json:"faxNumber,omitempty" binding:"e164PhoneNumbers"`
	URLAdr    string         `xml:"URLAdr,omitempty" json:"urlAddress,omitempty" binding:"max=2048"`
	EmailAdr  string         `xml:"EmailAdr,omitempty" json:"emailAddress,omitempty" binding:"email"`
	EmailPurp string         `xml:"EmailPurp,omitempty" json:"emailPurpose,omitempty" binding:"max=35"`
	JobTitl   string         `xml:"JobTitl,omitempty" json:"jobTitle,omitempty" binding:"max=35"`
	Rspnsblty string         `xml:"Rspnsblty,omitempty" json:"responsibility,omitempty" binding:"max=35"`
	Dept      string         `xml:"Dept,omitempty" json:"department,omitempty" binding:"max=70"`
	Othr      []*ContactOthr `xml:"Othr,omitempty" json:"other,omitzero" binding:"dive"`
	PrefrdMtd string         `xml:"PrefrdMtd,omitempty" json:"preferredMethod,omitempty" binding:"omitempty,oneof=MAIL FAXX LETT CELL ONLI PHON"`
}

type ContactOthr struct {
	Id      string `xml:"Id,omitempty" json:"id,omitempty"`
	ChanlTp string `xml:"ChanlTp,omitempty" json:"channelType,omitempty"`
}

type Account struct {
	Id   *AccountId    `xml:"Id" json:"accountId"`
	Tp   *AccountType  `xml:"Tp,omitempty" json:"type,omitzero"`
	Ccy  string        `xml:"Ccy,omitempty" json:"currency,omitempty" binding:"regexp=^[A-Z]{3}$"`
	Nm   string        `xml:"Nm,omitempty" json:"name,omitempty"`
	Prxy *AccountProxy `xml:"Prxy,omitempty" json:"proxy,omitzero"`
}

type AccountId struct {
	IBAN string       `xml:"IBAN,omitempty" json:"iban,omitempty" binding:"omitempty,required_without=Othr"`
	Othr *AccountOthr `xml:"Othr,omitempty" json:"other,omitzero" binding:"omitempty,required_without=IBAN"`
}

type AccountOthr struct {
	Id      string          `xml:"Id,omitempty" json:"id,omitempty" binding:"required,max=34"`
	Issr    string          `xml:"Issr,omitempty" json:"issuer,omitempty" binding:"max=35"`
	SchmeNm *AccountSchmeNm `xml:"SchmeNm,omitempty" json:"schemeName,omitzero" binding:"omitempty"`
}

type AccountSchmeNm struct {
	Cd    string `xml:"Cd,omitempty" json:"code,omitempty" binding:"required_without=Prtry,max=4"`
	Prtry string `xml:"Prtry,omitempty" json:"proprietary,omitempty" binding:"required_without=Cd,max=35"`
}

type AccountType struct {
	Cd    string `xml:"Cd,omitempty" json:"code,omitempty"`
	Prtry string `xml:"Prtry,omitempty" json:"proprietary,omitempty" binding:"max=35"`
}

type AccountProxy struct {
	Tp string `xml:"Tp,omitempty" json:"type,omitempty"`
	Id string `xml:"Id,omitempty" json:"id,omitempty"`
}

type InstrForCdtrAgt struct {
	Cd       string `xml:"Cd,omitempty" json:"code,omitempty" binding:"omitempty,min=1,max=4"`
	InstrInf string `xml:"InstrInf,omitempty" json:"instructionInfo,omitempty" binding:"max=140"`
}

type InstrForNxtAgt struct {
	Cd       string `xml:"Cd,omitempty" json:"code,omitempty" binding:"omitempty,oneof=PHOA TELA"`
	InstrInf string `xml:"InstrInf,omitempty" json:"instructionInfo,omitempty" binding:"max=140"`
}

type Purp struct {
	Cd    string `xml:"Cd,omitempty" json:"code,omitempty" binding:"omitempty,min=1,max=4,required_without=Prtry"`
	Prtry string `xml:"Prtry,omitempty" json:"proprietary,omitempty" binding:"required_without=Cd,max=35"`
}

type RgltryRptg struct {
	DbtCdtRptgInd string             `xml:"DbtCdtRptgInd,omitempty" json:"debitCreditReportingIndicator,omitempty" binding:"omitempty,oneof=CRED DEBT BOTH"`
	Authrty       *RgltryRptgAuthrty `xml:"Authrty,omitempty" json:"authority,omitzero"`
	Dtls          []*RgltryRptgDtls  `xml:"Dtls,omitempty" json:"details,omitempty" binding:"dive"`
}

type RgltryRptgAuthrty struct {
	Nm   string `xml:"Nm,omitempty" json:"name,omitempty" binding:"max=140"`
	Ctry string `xml:"Ctry,omitempty" json:"country,omitempty" binding:"regexp=^[A-Z]{2}$"`
}

type RgltryRptgDtls struct {
	Tp   string   `xml:"Tp,omitempty" json:"type,omitempty" binding:"max=35"`
	Dt   string   `xml:"Dt,omitempty" json:"date,omitempty" binding:"isoDate"`
	Ctry string   `xml:"Ctry,omitempty" json:"country,omitempty" binding:"regexp=^[A-Z]{2}$"`
	Cd   string   `xml:"Cd,omitempty" json:"code,omitempty" binding:"omitempty,min=1,max=10"`
	Amt  *Amount  `xml:"Amt,omitempty" json:"amount,omitzero"`
	Inf  []string `xml:"Inf,omitempty" json:"information,omitempty"`
}

type Tax struct {
	Cdtr            *TaxParty1 `xml:"Cdtr,omitempty" json:"creditor,omitzero"`
	Dbtr            *TaxParty2 `xml:"Dbtr,omitempty" json:"debtor,omitzero"`
	UltmtDbtr       *TaxParty2 `xml:"UltmtDbtr,omitempty" json:"ultimateDebtor,omitzero"`
	AdmstnZone      string     `xml:"AdmstnZone,omitempty" json:"administrationZone,omitempty"`
	RefNb           string     `xml:"RefNb,omitempty" json:"referenceNumber,omitempty"`
	Mtd             string     `xml:"Mtd,omitempty" json:"method,omitempty"`
	TtlTaxblBaseAmt *Amount    `xml:"TtlTaxblBaseAmt,omitempty" json:"totalTaxableBaseAmount,omitzero"`
	TtlTaxAmt       *Amount    `xml:"TtlTaxAmt,omitempty" json:"totalTaxAmount,omitzero"`
	Dt              *string    `xml:"Dt,omitempty" json:"date,omitempty" binding:"isoDate"`
	SeqNb           int64      `xml:"SeqNb,omitempty" json:"sequenceNumber,omitempty"`
	Rcrd            []*TaxRcrd `xml:"Rcrd,omitempty" json:"records,omitempty" binding:"dive"`
}

type TaxParty1 struct {
	TaxId  string `xml:"TaxId,omitempty" json:"taxId,omitempty"`
	RegnId string `xml:"RegnId,omitempty" json:"registrationId,omitempty"`
	TaxTp  string `xml:"TaxTp,omitempty" json:"taxType,omitempty" binding:"max=35"`
}

type TaxParty2 struct {
	TaxId   string           `xml:"TaxId,omitempty" json:"taxId,omitempty"`
	RegnId  string           `xml:"RegnId,omitempty" json:"registrationId,omitempty"`
	TaxTp   string           `xml:"TaxTp,omitempty" json:"taxType,omitempty"`
	Authstn *TaxPartyAuthstn `xml:"Authstn,omitempty" json:"authority,omitzero"`
}

type TaxPartyAuthstn struct {
	Titl string `xml:"Titl,omitempty" json:"title,omitempty"`
	Nm   string `xml:"Nm,omitempty" json:"name,omitempty"`
}

type TaxRcrd struct {
	Tp       string  `xml:"Tp,omitempty" json:"type,omitempty"`
	Ctgy     string  `xml:"Ctgy,omitempty" json:"category,omitempty"`
	CtgyDtls string  `xml:"CtgyDtls,omitempty" json:"categoryDetails,omitempty"`
	DbtrSts  string  `xml:"DbtrSts,omitempty" json:"debtorStatus,omitempty"`
	CertId   string  `xml:"CertId,omitempty" json:"certificateId,omitempty"`
	FrmsCd   string  `xml:"FrmsCd,omitempty" json:"formsCode,omitempty"`
	Prd      *TaxPrd `xml:"Prd,omitempty" json:"period,omitzero"`
	TaxAmt   *TaxAmt `xml:"TaxAmt,omitempty" json:"taxAmount,omitzero"`
	AddtlInf string  `xml:"AddtlInf,omitempty" json:"additionalInfo,omitempty"`
}

type TaxPrd struct {
	Yr     int16    `xml:"Yr,omitempty" json:"year,omitempty"`
	Tp     string   `xml:"Tp,omitempty" json:"type,omitempty"`
	FrToDt *DatePrd `xml:"FrToDt,omitempty" json:"fromToDate,omitzero"`
}

type DatePrd struct {
	FrDt string `xml:"FrDt,omitempty" json:"fromDate,omitempty" binding:"isoDate"`
	ToDt string `xml:"ToDt,omitempty" json:"toDate,omitempty" binding:"isoDate"`
}

type TaxAmt struct {
	Rate         float64     `xml:"Rate,omitempty" json:"rate,omitempty"`
	TaxblBaseAmt *Amount     `xml:"TaxblBaseAmt,omitempty" json:"taxableBaseAmount,omitzero"`
	TtlAmt       *Amount     `xml:"TtlAmt,omitempty" json:"totalAmount,omitzero"`
	Dtls         *TaxAmtDtls `xml:"Dtls,omitempty" json:"details,omitzero"`
}

type TaxAmtDtls struct {
	Prd *TaxPrd `xml:"Prd,omitempty" json:"period,omitzero"`
	Amt *Amount `xml:"Amt,omitempty" json:"Amount,omitzero"`
}

type RltdRmtInf struct {
	RmtId       string         `xml:"RmtId,omitempty" json:"remittanceId,omitempty"`
	RmtLctnDtls []*RmtLctnDtls `xml:"RmtLctnDtls,omitempty" json:"remittanceLocationDetails,omitempty" binding:"dive"`
}

type RmtLctnDtls struct {
	Mtd        string    `xml:"Mtd,omitempty" json:"method,omitempty"`
	ElctrncAdr string    `xml:"ElctrncAdr,omitempty" json:"electronicAddress,omitempty"`
	PstlAdr    *NmAndAdr `xml:"PstlAdr,omitempty" json:"postalAddress,omitzero"`
}

type NmAndAdr struct {
	Nm  string   `xml:"Nm,omitempty" json:"name,omitempty"`
	Adr *PstlAdr `xml:"Adr,omitempty" json:"address,omitzero"`
}

type RmtInf struct {
	Ustrd []string      `xml:"Ustrd,omitempty" json:"ustrdRemittanceInfo,omitempty"`
	Strd  []*RmtInfStrd `xml:"Strd,omitempty" json:"strdRemittanceInfo,omitempty" binding:"dive"`
}

type RmtInfStrd struct {
	RfrdDocInf  []*RfrdDocInf   `xml:"RfrdDocInf,omitempty" json:"referredDocumentInfo,omitempty" binding:"dive"`
	RfrdDocAmt  *LineDtlsAmount `xml:"RfrdDocAmt,omitempty" json:"referredDocumentAmount,omitzero"`
	CdtrRefInf  *CdtrRefInf     `xml:"CdtrRefInf,omitempty" json:"creditorReferenceInfo,omitzero"`
	Invcr       *Party          `xml:"Invcr,omitempty" json:"invoicer,omitzero"`
	Invcee      *Party          `xml:"Invcee,omitempty" json:"invoicee,omitzero"`
	TaxRmt      *Tax            `xml:"TaxRmt,omitempty" json:"taxRemittance,omitzero"`
	GrnshmtRmt  *GrnshmtRmt     `xml:"GrnshmtRmt,omitempty" json:"garnishmentRemittance,omitzero"`
	AddtlRmtInf []string        `xml:"AddtlRmtInf,omitempty" json:"additionalRemittanceInfo,omitempty"`
}

type RfrdDocInf struct {
	Tp       *RfrdDocInfTp `xml:"Tp,omitempty" json:"type,omitzero"`
	Nb       string        `xml:"Nb,omitempty" json:"number,omitempty"`
	RltdDt   *DtAndTp      `xml:"RltdDt,omitempty" json:"relatedDate,omitzero"`
	LineDtls []*LineDtls   `xml:"LineDtls,omitempty" json:"lineDetails,omitempty" binding:"dive"`
}

type RfrdDocInfTp struct {
	CdOrPrtry *CdOrPrtry `xml:"CdOrPrtry,omitempty" json:"codeOrProprietary,omitzero"`
	Issr      string     `xml:"Issr,omitempty" json:"issuer,omitempty"`
}

type CdOrPrtry struct {
	Cd    string `xml:"Cd,omitempty" json:"code,omitempty"`
	Prtry string `xml:"Prtry,omitempty" json:"proprietary,omitempty" binding:"max=35"`
}

type DtAndTp struct {
	Tp string `xml:"Tp,omitempty" json:"type,omitempty"`
	Dt string `xml:"Dt,omitempty" json:"date,omitempty" binding:"isoDate"`
}

type LineDtls struct {
	Id   []*LineDtlsId   `xml:"Id,omitempty" json:"identifiers,omitempty" binding:"dive"`
	Desc string          `xml:"Desc,omitempty" json:"description,omitempty"`
	Amt  *LineDtlsAmount `xml:"Amt,omitempty" json:"Amount,omitzero"`
}

type LineDtlsId struct {
	Tp     *LineDtlsIdTp `xml:"Tp,omitempty" json:"type,omitzero"`
	Nb     string        `xml:"Nb,omitempty" json:"number,omitempty"`
	RltdDt string        `xml:"RltdDt,omitempty" json:"relatedDate,omitempty" binding:"isoDate"`
}

type LineDtlsIdTp struct {
	CdOrPrtry *CdOrPrtry `xml:"CdOrPrtry,omitempty" json:"codeOrProprietary,omitzero"`
	Issr      string     `xml:"Issr,omitempty" json:"issuer,omitempty"`
}

type LineDtlsAmount struct {
	RmtAmtAndTp       []*LineDtlsRmtAmtAndTp       `xml:"RmtAmtAndTp,omitempty" json:"remittanceAmountAndType,omitempty" binding:"dive"`
	AdjstmntAmtAndRsn []*LineDtlsAdjstmntAmtAndRsn `xml:"AdjstmntAmtAndRsn,omitempty" json:"adjustmentAmountAndReason,omitempty" binding:"dive"`
}

type LineDtlsRmtAmtAndTp struct {
	Tp  *LineDtlsCdOrPrtry `xml:"Tp,omitempty" json:"type,omitzero"`
	Amt Amount             `xml:"Amt,omitempty" json:"Amount,omitzero"`
}

type LineDtlsCdOrPrtry struct {
	Cd    string `xml:"Cd,omitempty" json:"code,omitempty"`
	Prtry string `xml:"Prtry,omitempty" json:"proprietary,omitempty" binding:"max=35"`
}

type LineDtlsAdjstmntAmtAndRsn struct {
	Amt       *Amount `xml:"Amt,omitempty" json:"Amount,omitzero"`
	CdtDbtInd string  `xml:"CdtDbtInd,omitempty" json:"creditDebitIndicator,omitempty"`
	Rsn       string  `xml:"Rsn,omitempty" json:"reason,omitempty"`
	AddtlInf  string  `xml:"AddtlInf,omitempty" json:"additionalInfo,omitempty"`
}

type CdtrRefInf struct {
	Tp  *CdtrRefInfTp `xml:"Tp,omitempty" json:"type,omitzero"`
	Ref string        `xml:"Ref,omitempty" json:"reference,omitempty"`
}

type CdtrRefInfTp struct {
	CdOrPrtry *CdOrPrtry `xml:"CdOrPrtry,omitempty" json:"codeOrProprietary,omitzero"`
	Issr      string     `xml:"Issr,omitempty" json:"issuer,omitempty"`
}

type GrnshmtRmt struct {
	Tp                *GrnshmtRmtTp `xml:"Tp,omitempty" json:"type,omitzero"`
	Grnshee           *Party        `xml:"Grnshee,omitempty" json:"garnishee,omitzero"`
	GrnshmtAdmstr     *Party        `xml:"GrnshmtAdmstr,omitempty" json:"garnishmentAdministrator,omitzero"`
	RefNb             string        `xml:"RefNb,omitempty" json:"referenceNumber,omitempty"`
	Dt                string        `xml:"Dt,omitempty" json:"date,omitempty"`
	RmtdAmt           *Amount       `xml:"RmtdAmt,omitempty" json:"remittedAmount,omitzero"`
	FmlyMdclInsrncInd bool          `xml:"FmlyMdclInsrncInd,omitempty" json:"familyMedicalInsuranceIndicator,omitempty"`
	MplyeeTermntnInd  bool          `xml:"MplyeeTermntnInd,omitempty" json:"employeeTerminationIndicator,omitempty"`
}

type GrnshmtRmtTp struct {
	CdOrPrtry *CdOrPrtry `xml:"CdOrPrtry,omitempty" json:"codeOrProprietary,omitzero"`
	Issr      string     `xml:"Issr,omitempty" json:"issuer,omitempty"`
}

type SplmtryData struct {
	PlcAndNm string `xml:"PlcAndNm,omitempty" json:"placeAndName,omitempty" binding:"max=350"`
	Envlp    *Envlp `xml:"Envlp" json:"envelope" binding:"required"`
}

type Envlp struct {
	Data string `xml:",innerxml" json:"data" binding:"required"`
}
