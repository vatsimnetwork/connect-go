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
	Cid int `json:"cid"`
	Personal Personal `json:"personal"`
	Vatsim vatsim `json:"vatsim"`
	OAuth OAuth `json:"oauth"`
}

// Data represents user personal details
type Personal struct {
	Name_first string `json:"name_first"`
	Name_last string `json:"name_last"`
	Name_full string `json:"name_full"`
}

// Data represents user vatsim details
type vatsim struct {
	Rating Rating `json:"rating"`
	PilotRating Pilot `json:"pilotrating"`
	Division Division `json:"division"`
	Region Region `json:"region"`
	SubDivision SubDivision `json:"subdivision"`
}

// Data represents user rating details
type Rating struct {
	Id int `json:"id"`
	Long string `json:"long"`
	Short string `json:"short"`
}

// Data represents user pilot rating details
type Pilot struct {
	Id int `json:"id"`
}

// Data represents user division details
type Division struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

// Data represents user region details
type Region struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

// Data represents user subdivision details
type SubDivision struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

// Data represents oatuh details
type OAuth struct {
	TokenValid string `json:"token_valid"`
}