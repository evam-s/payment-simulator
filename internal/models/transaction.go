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
	OrgIdLei                 string          `json:"orgIdLei,omitempty" binding:"regexp=^[A-Z0-9]{18}[0-9]{2}$"`
	OrgIdOther               []OrgIdOthr     `json:"orgIdOther,omitzero"`
	PrivateIdBirthDate       string          `json:"birthDate,omitempty"`
	PrivateIdProvinceOfBirth string          `json:"provinceOfBirth,omitempty"`
	PrivateIdCityOfBirth     string          `json:"cityOfBirth,omitempty"`
	PrivateIdCountryOfBirth  string          `json:"countryOfBirth,omitempty" binding:"regexp=^[A-Z]{2}$"`
	PrivateIdOther           []PrivateIdOthr `json:"other,omitzero"`
	CountryOfResidence       string          `json:"countryOfResidence,omitempty" binding:"regexp=^[A-Z]{2}$"`
	ContactDetails           CtctDtls        `json:"contactDetails,omitzero"`
}

type PstlAdr struct {
	AddressTypeCode                  string `json:"addressTypeCode,omitempty"`
	AddressTypeProprietaryId         string `json:"addressTypeProprietaryId,omitempty"`
	AddressTypeProprietaryIssuer     string `json:"addressTypeProprietaryIssuer,omitempty"`
	AddressTypeProprietarySchemeName string `json:"addressTypeProprietarySchemeName,omitempty"`
	CareOf                           string `json:"careOf,omitempty" binding:"max=140"`
	Department                       string `json:"department,omitempty" binding:"max=70"`
	SubDepartment                    string `json:"subDepartment,omitempty" binding:"max=70"`
	StreetName                       string `json:"streetName,omitempty" binding:"max=140"`
	BuildingNumber                   string `json:"buildingNumber,omitempty" binding:"max=16"`
	BuildingName                     string `json:"buildingName,omitempty" binding:"max=140"`
	Floor                            string `json:"floor,omitempty" binding:"max=70"`
	UnitNumber                       string `json:"unitNumber,omitempty" binding:"max=16"`
	PostBox                          string `json:"postBox,omitempty" binding:"max=16"`
	Room                             string `json:"room,omitempty" binding:"max=70"`
	PostalCode                       string `json:"postalCode,omitempty" binding:"max=16"`
	TownName                         string `json:"townName,omitempty" binding:"max=140"`
	TownLocationName                 string `json:"townLocationName,omitempty" binding:"max=140"`
	DistrictName                     string `json:"districtName,omitempty" binding:"max=140"`
	CountrySubdivision               string `json:"countrySubdivision,omitempty" binding:"max=35"`
	Country                          string `json:"country,omitempty" binding:"regexp=^[A-Z]{2}$"`
	AddressLine                      string `json:"addressLine,omitempty" binding:"max=70"`
}

type OrgIdOthr struct {
	Id                    string `json:"id,omitempty"`
	Issuer                string `json:"issuer,omitempty"`
	SchemeNameCode        string `json:"schemeNameCode,omitempty"`
	SchemeNameProprietary string `json:"schemeNameProprietary,omitempty" binding:"max=35"`
}

type PrivateIdOthr struct {
	Id                    string `json:"id,omitempty"`
	Issuer                string `json:"issuer,omitempty"`
	SchemeNameCode        string `json:"schemeNameCode,omitempty"`
	SchemeNameProprietary string `json:"schemeNameProprietary,omitempty" binding:"max=35"`
}

type CtctDtls struct {
	NamePrefix      string        `json:"namePrefix,omitempty"`
	Name            string        `json:"name,omitempty"`
	PhoneNumber     string        `json:"phoneNumber,omitempty" binding:"regexp=^\\+[0-9]{1,3}-[0-9()+\\-]{1,30}$"`
	MobileNumber    string        `json:"mobileNumber,omitempty" binding:"regexp=^\\+[0-9]{1,3}-[0-9()+\\-]{1,30}$"`
	FaxNumber       string        `json:"faxNumber,omitempty" binding:"regexp=^\\+[0-9]{1,3}-[0-9()+\\-]{1,30}$"`
	UrlAddress      string        `json:"urlAddress,omitempty"`
	EmailAddress    string        `json:"emailAddress,omitempty" binding:"email"`
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
	Currency                   string `json:"currency,omitempty" binding:"regexp=^[A-Z]{3}$"`
	Name                       string `json:"name,omitempty"`
	ProxyId                    string `json:"proxyId,omitempty"`
	ProxyType                  string `json:"proxyType,omitempty"`
}

type Agent struct {
	FinancialInstitutionId FinInstnId `json:"financialInstitutionId" binding:"required"`
	BranchId               BrnchId    `json:"branchId,omitzero"`
}

type FinInstnId struct {
	Bicfi                       string  `json:"bicfi,omitempty" binding:"regexp=^[A-Z0-9]{4}[A-Z]{2}[A-Z0-9]{2}([A-Z0-9]{3})?$"`
	ClearingSystemIdCode        string  `json:"clearingSystemIdCode,omitzero" binding:"min=1,max=5"`
	ClearingSystemIdProprietary string  `json:"clearingSystemIdProprietary,omitzero" binding:"max=35"`
	MemberId                    string  `json:"memberId,omitempty" binding:"max=35"`
	Lei                         string  `json:"lei,omitempty" binding:"regexp=^[A-Z0-9]{18}[0-9]{2}$"`
	Name                        string  `json:"name,omitempty" binding:"max=140"`
	PostalAddress               PstlAdr `json:"postalAddress,omitzero"`
	OtherId                     string  `json:"otherId,omitempty" binding:"max=35"`
	OtherIssuer                 string  `json:"otherIssuer,omitempty" binding:"max=35"`
	OtherSchemeNameCode         string  `json:"otherSchemeNameCode,omitzero" binding:"min=1,max=4"`
	OtherSchemeNameProprietary  string  `json:"otherSchemeNameProprietary,omitzero" binding:"max=35"`
}

type BrnchId struct {
	Id            string  `json:"id,omitempty" binding:"max=35"`
	Lei           string  `json:"lei,omitempty" binding:"regexp=^[A-Z0-9]{18}[0-9]{2}$"`
	Name          string  `json:"name,omitempty" binding:"max=140"`
	PostalAddress PstlAdr `json:"postalAddress,omitzero"`
}

type Charges struct {
	Amount Amount  `json:"amount,omitzero" binding:"required"`
	Agent  Agent   `json:"agent,omitzero" binding:"required"`
	Type   ChrgsTp `json:"type,omitzero"`
}

type Amount struct {
	Value    float64 `json:"value" binding:"required"`
	Currency string  `json:"currency" binding:"required,regexp=^[A-Z]{3}$"`
}

type ChrgsTp struct {
	Code              string `json:"code,omitempty"`
	ProprietaryId     string `json:"proprietaryId,omitzero"`
	ProprietaryIssuer string `json:"proprietaryIssuer,omitzero"`
}
