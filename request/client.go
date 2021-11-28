package request

type Client struct {
	Name        string `json:"name"`
	RedirectUri string `json:"redirect_uri"`
	Type        string `json:"type"`
}
