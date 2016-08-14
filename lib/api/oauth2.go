package api

type OAuth2 struct {
	ClientID    string
	Secret      string
	AuthUrl     string
	TokenUrl    string
	RedirectUrl string
	Scope       string
}
