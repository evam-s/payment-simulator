package models

type PaymentOrder struct {
	Debtor                  Party     `json:"debtor"`
	DebtorAcct              Account   `json:"debtorAcct"`
	Creditor                Party     `json:"creditor"`
	CreditorAcct            Account   `json:"creditorAcct"`
	SettlementAmount        Amount    `json:"settlementAmount"`
	InstructionId           string    `json:"instructionId,omitempty"`
	TransactionId           string    `json:"transactionId,omitempty"`
	EndToEndId              string    `json:"endToEndId,omitempty"`
	UETR                    string    `json:"uetr,omitempty"`
	ClearingSystemReference string    `json:"clearingSystemReference,omitempty"`
	MsgId                   string    `json:"msgId,omitempty"`
	CreationDateTime        string    `json:"creationDateTime,omitempty"`
	NumberOfTxs             string    `json:"numberOfTxs,omitempty"`
	SettlementMethod        string    `json:"settlementMethod,omitempty"`
	ControlSum              float64   `json:"controlSum,omitempty"`
	SettlementDate          string    `json:"settlementDate,omitempty"`
	SettlementPriority      string    `json:"settlementPriority,omitempty"`
	RemittanceInfo          []string  `json:"remittanceInfo,omitempty"`
	Purpose                 string    `json:"purpose,omitempty"`
	ChargeBearer            string    `json:"chargeBearer,omitempty"`
	Charges                 []Charges `json:"charges,omitempty"`
	Status                  string    `json:"status,omitempty"`
	EntityId                string    `json:"entityId,omitempty"`
	Errors                  []string  `json:"errors,omitempty"`
	// ReferenceNumber         string    `json:"referenceNumber,omitempty"`
	// TotalTaxAmount          Amount    `json:"totalTaxAmount,omitempty"`
}

type Party struct {
	Name                     string          `json:"name,omitempty"`
	PostalAddress            PstlAdr         `json:"postalAddress,omitzero"`
	OrgIdAnyBic              string          `json:"orgIdAnyBic,omitempty"`
	OrgIdLei                 string          `json:"orgIdLei,omitempty"`
	OrgIdOther               []OrgIdOthr     `json:"orgIdOther,omitzero"`
	PrivateIdBirthDate       string          `json:"birthDate,omitempty"`
	PrivateIdProvinceOfBirth string          `json:"provinceOfBirth,omitempty"`
	PrivateIdCityOfBirth     string          `json:"cityOfBirth,omitempty"`
	PrivateIdCountryOfBirth  string          `json:"countryOfBirth,omitempty"`
	PrivateIdOther           []PrivateIdOthr `json:"other,omitzero"`
	CountryOfResidence       string          `json:"countryOfResidence,omitempty"`
	ContactDetails           CtctDtls        `json:"contactDetails,omitzero"`
}

type PstlAdr struct {
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

type CtctDtls struct {
	NamePrefix      string        `json:"namePrefix,omitempty"`
	Name            string        `json:"name,omitempty"`
	PhoneNumber     string        `json:"phoneNumber,omitempty"`
	MobileNumber    string        `json:"mobileNumber,omitempty"`
	FaxNumber       string        `json:"faxNumber,omitempty"`
	UrlAddress      string        `json:"urlAddress,omitempty"`
	EmailAddress    string        `json:"emailAddress,omitempty"`
	EmailPurpose    string        `json:"emailPurpose,omitempty"`
	JobTitle        string        `json:"jobTitle,omitempty"`
	Responsibility  string        `json:"responsibility,omitempty"`
	Department      string        `json:"department,omitempty"`
	Other           []ContactOthr `json:"other,omitzero"`
	PreferredMethod string        `json:"preferredMethod,omitempty"`
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
	FinancialInstitutionId FinInstnId `json:"financialInstitutionId"`
	BranchId               BrnchId    `json:"branchId,omitzero"`
}

type FinInstnId struct {
	Bicfi                       string  `json:"bicfi,omitempty"`
	ClearingSystemIdCode        string  `json:"clearingSystemIdCode,omitzero"`
	ClearingSystemIdProprietary string  `json:"clearingSystemIdProprietary,omitzero"`
	MemberId                    string  `json:"memberId,omitempty"`
	Lei                         string  `json:"lei,omitempty"`
	Name                        string  `json:"name,omitempty"`
	PostalAddress               PstlAdr `json:"postalAddress,omitzero"`
	OtherId                     string  `json:"otherId,omitempty"`
	OtherIssuer                 string  `json:"otherIssuer,omitempty"`
	OtherSchemeNameCode         string  `json:"otherSchemeNameCode,omitzero"`
	OtherSchemeNameProprietary  string  `json:"otherSchemeNameProprietary,omitzero"`
}

type BrnchId struct {
	Id            string  `json:"id,omitempty"`
	Lei           string  `json:"lei,omitempty"`
	Name          string  `json:"name,omitempty"`
	PostalAddress PstlAdr `json:"postalAddress,omitzero"`
}

type Charges struct {
	Amount Amount  `json:"amount,omitzero"`
	Agent  Agent   `json:"agent,omitzero"`
	Type   ChrgsTp `json:"type,omitzero"`
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
