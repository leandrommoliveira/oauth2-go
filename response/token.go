package response

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	Scopes       string `json:"scopes"`
	ExpiresIn    string `json:"expires_in"`
}
