package models

import "time"

type PaymentOrder struct {
	Id                          string               `json:"id" bson:"_id"`
	Status                      string               `json:"status,omitempty"`
	EntityId                    string               `json:"entityId"`
	Debtor                      *Party               `json:"debtor"`
	UltimateDebtor              *Party               `json:"ultimateDebtor"`
	DebtorAcct                  *Account             `json:"debtorAcct"`
	DebtorAgent                 *Agent               `json:"debtorAgent"`
	DebtorAgentAcct             *Account             `json:"debtorAgentAcct"`
	Creditor                    *Party               `json:"creditor"`
	UltimateCreditor            *Party               `json:"ultimateCreditor"`
	CreditorAcct                *Account             `json:"creditorAcct"`
	CreditorAgent               *Agent               `json:"creditorAgent"`
	CreditorAgentAcct           *Account             `json:"creditorAgentAcct"`
	SettlementAmount            *Amount              `json:"settlementAmount"`
	InstructionId               string               `json:"instructionId,omitempty"`
	SettlementAcct              *Account             `json:"settlementAcct"`
	TransactionId               string               `json:"transactionId,omitempty"`
	EndToEndId                  string               `json:"endToEndId,omitempty"`
	UETR                        string               `json:"uetr,omitempty"`
	ClearingSystemReference     string               `json:"clearingSystemReference,omitempty"`
	MessageId                   string               `json:"messageId,omitempty"`
	MessageNameId               string               `json:"messageNameId,omitempty"`
	CreationDateTime            string               `json:"creationDateTime,omitempty"`
	NumberOfTxs                 string               `json:"numberOfTxs,omitempty"`
	SettlementMethod            string               `json:"settlementMethod,omitempty"`
	ControlSum                  float64              `json:"controlSum,omitempty"`
	SettlementDate              string               `json:"settlementDate,omitempty"`
	SettlementPriority          string               `json:"settlementPriority,omitempty"`
	UnstructuredRemittanceInfo  []string             `json:"unstructuredRemittanceInfo,omitempty"`
	PurposeCode                 string               `json:"purposeCode,omitempty"`
	PurposeProprietary          string               `json:"purposeProprietary,omitempty"`
	ChargeBearer                string               `json:"chargeBearer,omitempty"`
	Charges                     []*Charges           `json:"charges,omitempty"`
	Errors                      []string             `json:"errors,omitempty"`
	TxnAcceptanceDateTime       string               `json:"txnAcceptanceDateTime,omitempty"`
	ProcessingDateTime          string               `json:"processingDateTime,omitempty"`
	EffectiveSettlementDateTime string               `json:"effectiveSettlementDateTime,omitempty"`
	PaymentTypeInfo             *PaymentTypeInfo     `json:"paymentTypeInfo"`
	HeaderPaymentTypeInfo       *PaymentTypeInfo     `json:"headerPaymentTypeInfo"`
	SupplementaryData           []*SupplementaryData `json:"supplementaryData,omitempty"`
	HeaderSupplementaryData     []*SupplementaryData `json:"headerSupplementaryData,omitempty"`
	CreatedOn                   time.Time
	// StructuredRemittanceInfo    []StructuredRemittanceInfo `json:"structuredRemittanceInfo,omitempty"`
	// ReferenceNumber             string                     `json:"referenceNumber,omitempty"`
	// TotalTaxAmount              Amount                     `json:"totalTaxAmount,omitempty"`
}

type Party struct {
	Name                     string           `json:"name,omitempty"`
	PostalAddress            *PostalAddress   `json:"postalAddress,omitzero"`
	OrgIdAnyBic              string           `json:"orgIdAnyBic,omitempty"`
	OrgIdLei                 string           `json:"orgIdLei,omitempty"`
	OrgIdOther               []*OrgIdOthr     `json:"orgIdOther,omitzero"`
	PrivateIdBirthDate       string           `json:"birthDate,omitempty"`
	PrivateIdProvinceOfBirth string           `json:"provinceOfBirth,omitempty"`
	PrivateIdCityOfBirth     string           `json:"cityOfBirth,omitempty"`
	PrivateIdCountryOfBirth  string           `json:"countryOfBirth,omitempty"`
	PrivateIdOther           []*PrivateIdOthr `json:"other,omitzero"`
	CountryOfResidence       string           `json:"countryOfResidence,omitempty"`
	ContactDetails           *ContactDetails  `json:"contactDetails,omitzero"`
	IsFinancialInstitution   bool             `json:"isFinancialInstitution,omitzero"`
}

type PostalAddress struct {
	AddressTypeCode                  string `json:"addressTypeCode,omitempty"`
	AddressTypeProprietaryId         string `json:"addressTypeProprietaryId,omitempty"`
	AddressTypeProprietaryIssuer     string `json:"addressTypeProprietaryIssuer,omitempty"`
	AddressTypeProprietarySchemeName string `json:"addressTypeProprietarySchemeName,omitempty"`
	CareOf                           string `json:"careOf,omitempty"`
	Department                       string `json:"department,omitempty"`
	SubDepartment                    string `json:"subDepartment,omitempty"`
	StreetName                       string `json:"streetName,omitempty"`
	BuildingNumber                   string `json:"buildingNumber,omitempty"`
	BuildingName                     string `json:"buildingName,omitempty"`
	Floor                            string `json:"floor,omitempty"`
	UnitNumber                       string `json:"unitNumber,omitempty"`
	PostBox                          string `json:"postBox,omitempty"`
	Room                             string `json:"room,omitempty"`
	PostalCode                       string `json:"postalCode,omitempty"`
	TownName                         string `json:"townName,omitempty"`
	TownLocationName                 string `json:"townLocationName,omitempty"`
	DistrictName                     string `json:"districtName,omitempty"`
	CountrySubdivision               string `json:"countrySubdivision,omitempty"`
	Country                          string `json:"country,omitempty"`
	AddressLine                      string `json:"addressLine,omitempty"`
}

type OrgIdOthr struct {
	Id                    string `json:"id,omitempty"`
	Issuer                string `json:"issuer,omitempty"`
	SchemeNameCode        string `json:"schemeNameCode,omitempty"`
	SchemeNameProprietary string `json:"schemeNameProprietary,omitempty"`
}

type PrivateIdOthr struct {
	Id                    string `json:"id,omitempty"`
	Issuer                string `json:"issuer,omitempty"`
	SchemeNameCode        string `json:"schemeNameCode,omitempty"`
	SchemeNameProprietary string `json:"schemeNameProprietary,omitempty"`
}

type ContactDetails struct {
	NamePrefix      string         `json:"namePrefix,omitempty"`
	Name            string         `json:"name,omitempty"`
	PhoneNumber     string         `json:"phoneNumber,omitempty"`
	MobileNumber    string         `json:"mobileNumber,omitempty"`
	FaxNumber       string         `json:"faxNumber,omitempty"`
	UrlAddress      string         `json:"urlAddress,omitempty"`
	EmailAddress    string         `json:"emailAddress,omitempty"`
	EmailPurpose    string         `json:"emailPurpose,omitempty"`
	JobTitle        string         `json:"jobTitle,omitempty"`
	Responsibility  string         `json:"responsibility,omitempty"`
	Department      string         `json:"department,omitempty"`
	PreferredMethod string         `json:"preferredMethod,omitempty"`
	Other           []*ContactOthr `json:"other,omitzero"`
}

type ContactOthr struct {
	Id          string `json:"id,omitempty"`
	ChannelType string `json:"channelType,omitempty"`
}

type Account struct {
	Iban                       string `json:"iban,omitempty"`
	OtherId                    string `json:"otherId,omitempty"`
	OtherIssuer                string `json:"otherIssuer,omitempty"`
	OtherSchemeNameCode        string `json:"otherSchemeNameCode,omitempty"`
	OtherSchemeNameProprietary string `json:"otherSchemeNameProprietary,omitempty"`
	TypeCode                   string `json:"typeCode,omitempty"`
	TypeProprietary            string `json:"typeProprietary,omitempty"`
	Currency                   string `json:"currency,omitempty"`
	Name                       string `json:"name,omitempty"`
	ProxyId                    string `json:"proxyId,omitempty"`
	ProxyType                  string `json:"proxyType,omitempty"`
}

type Agent struct {
	FiiBicfi                       string         `json:"fiiBicfi,omitempty"`
	FiiClearingSystemIdCode        string         `json:"fiiClearingSystemIdCode,omitzero"`
	FiiClearingSystemIdProprietary string         `json:"fiiClearingSystemIdProprietary,omitzero"`
	FiiMemberId                    string         `json:"fiiMemberId,omitempty"`
	FiiLei                         string         `json:"fiiLei,omitempty"`
	FiiName                        string         `json:"fiiName,omitempty"`
	FiiPostalAddress               *PostalAddress `json:"fiiPostalAddress,omitzero"`
	FiiOtherId                     string         `json:"fiiOtherId,omitempty"`
	FiiOtherIssuer                 string         `json:"fiiOtherIssuer,omitempty"`
	FiiOtherSchemeNameCode         string         `json:"fiiOtherSchemeNameCode,omitzero"`
	FiiOtherSchemeNameProprietary  string         `json:"fiiOtherSchemeNameProprietary,omitzero"`
	BiId                           string         `json:"biId,omitempty"`
	BiLei                          string         `json:"biLei,omitempty"`
	BiName                         string         `json:"biName,omitempty"`
	BiPostalAddress                *PostalAddress `json:"biPostalAddress,omitzero"`
}

type FinInstnId struct {
	Bicfi                       string         `json:"bicfi,omitempty"`
	ClearingSystemIdCode        string         `json:"clearingSystemIdCode,omitzero"`
	ClearingSystemIdProprietary string         `json:"clearingSystemIdProprietary,omitzero"`
	MemberId                    string         `json:"memberId,omitempty"`
	Lei                         string         `json:"lei,omitempty"`
	Name                        string         `json:"name,omitempty"`
	PostalAddress               *PostalAddress `json:"postalAddress,omitzero"`
	OtherId                     string         `json:"otherId,omitempty"`
	OtherIssuer                 string         `json:"otherIssuer,omitempty"`
	OtherSchemeNameCode         string         `json:"otherSchemeNameCode,omitzero"`
	OtherSchemeNameProprietary  string         `json:"otherSchemeNameProprietary,omitzero"`
}

type BrnchId struct {
	Id            string         `json:"id,omitempty"`
	Lei           string         `json:"lei,omitempty"`
	Name          string         `json:"name,omitempty"`
	PostalAddress *PostalAddress `json:"postalAddress,omitzero"`
}

type Charges struct {
	Amount *Amount  `json:"amount,omitzero"`
	Agent  *Agent   `json:"agent,omitzero"`
	Type   *ChrgsTp `json:"type,omitzero"`
}

type Amount struct {
	Value    float64 `json:"value"`
	Currency string  `json:"currency"`
}

type ChrgsTp struct {
	Code              string `json:"code,omitempty"`
	ProprietaryId     string `json:"proprietaryId,omitzero"`
	ProprietaryIssuer string `json:"proprietaryIssuer,omitzero"`
}

type PaymentTypeInfo struct {
	InstructionPriority        string               `json:"instructionPriority,omitempty"`
	ClearingChannel            string               `json:"clearingChannel,omitempty"`
	ServiceLevel               []*CodeOrProprietary `json:"serviceLevel,omitempty"`
	LocalInstrumentCode        string               `json:"localInstrumentCode,omitempty"`
	LocalInstrumentProprietary string               `json:"localInstrumentProprietary,omitempty"`
	CategoryPurposeCode        string               `json:"categoryPurposeCode,omitempty"`
	CategoryPurposeProprietary string               `json:"categoryPurposeProprietary,omitempty"`
}

type CodeOrProprietary struct {
	Code        string `json:"code,omitempty"`
	Proprietary string `json:"proprietary,omitempty"`
}

type SupplementaryData struct {
	PlaceAndName string `json:"placeAndName,omitempty"`
	Envelope     string `json:"envelope,omitempty"`
}
