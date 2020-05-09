package Handle

// An Access represents access token which we get from AccessToken function
type Access struct {
	Token string `json:"access_token"`
}

// An User represents response from Connect which we get in GetData function
type User struct {
	Data Data `json:"data"`
}

// Data represents user details return from Connect in GetData function
type Data struct {
	CID int `json:"cid"`
	Personal Personal `json:"personal"`
	Vatsim Vatsim `json:"vatsim"`
	OAuth OAuth `json:"oauth"`
}

// Personal represents user personal details
type Personal struct {
	NameFirst string `json:"name_first"`
	NameLast string `json:"name_last"`
	NameFull string `json:"name_full"`
}

// Vatsim represents user vatsim details
type Vatsim struct {
	Rating Rating `json:"rating"`
	PilotRating Pilot `json:"pilotrating"`
	Division Division `json:"division"`
	Region Region `json:"region"`
	Subdivision Subdivision `json:"subdivision"`
}

// Rating represents user controller rating details
type Rating struct {
	ID int `json:"id"`
	Long string `json:"long"`
	Short string `json:"short"`
}

// Pilot represents user pilot rating details
type Pilot struct {
	ID int `json:"id"`
}

// Division represents user division details
type Division struct {
	ID string `json:"id"`
	Name string `json:"name"`
}

// Region represents user region details
type Region struct {
	ID string `json:"id"`
	Name string `json:"name"`
}

// Subdivision represents user subdivision details
type Subdivision struct {
	ID string `json:"id"`
	Name string `json:"name"`
}

// OAuth represents oatuh details
type OAuth struct {
	TokenValid string `json:"token_valid"`
}