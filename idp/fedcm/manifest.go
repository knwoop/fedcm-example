package fedcm

type Manifest struct {
	AccountsEndpoint       string    `json:"accounts_endpoint"`
	ClientMetadataEndpoint string    `json:"client_metadata_endpoint"`
	IDAssertionEndpoint    string    `json:"id_assertion_endpoint"`
	Branding               *Branding `json:"branding,omitempty"`
}

type Branding struct {
	BackgroundColor string  `json:"background_color,omitempty"`
	Color           string  `json:"color,omitempty"`
	Icons           []*Icon `json:"icons,omitempty"`
}

type Icon struct {
	Url  string `json:"url"`
	Size int    `json:"size,omitempty"`
}
