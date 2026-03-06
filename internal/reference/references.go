package reference

type Account struct {
	Id          string          `json:"id" bson:"_id"`
	Name        string          `json:"name"`
	PhoneNumber string          `json:"phoneNumber"`
	Email       string          `json:"email"`
	Currency    string          `json:"currency"`
	Balance     int             `json:"balance"`
	Active      bool            `json:"active"`
	Address     *AccountAddress `json:"address"`
}

type AccountAddress struct {
	Type               string `json:"type"`
	CareOf             string `json:"careOf"`
	Department         string `json:"department"`
	SubDepartment      string `json:"subDepartment"`
	StreetName         string `json:"streetName"`
	BuildingNumber     string `json:"buildingNumber"`
	BuildingName       string `json:"buildingName"`
	Floor              string `json:"floor"`
	UnitNumber         string `json:"unitNumber"`
	PostBox            string `json:"postBox"`
	Room               string `json:"room"`
	PostalCode         string `json:"postalCode"`
	TownName           string `json:"townName"`
	TownLocationName   string `json:"townLocationName"`
	DistrictName       string `json:"districtName"`
	CountrySubdivision string `json:"countrySubdivision"`
	Country            string `json:"country"`
	AddressLine        string `json:"addressLine"`
}
