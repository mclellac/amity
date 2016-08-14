package api

type User struct {
	Id         int
	GoogleID   int    `json:"id"`
	Name       string `json:"name"`
	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
	Link       string `json:"link"`
	Picture    string `json:"picture"`
	Gender     string `json:"gender"`
	Locale     string `json:"locale"`
}
