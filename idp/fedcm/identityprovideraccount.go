package fedcm

type IdentityProviderAccounts struct {
	Accounts []*IdentityProviderAccount `json:"accounts"`
}

type IdentityProviderAccount struct {
	ID              string   `json:"id"`
	GivenName       string   `json:"given_name"`
	Name            string   `json:"name"`
	Email           string   `json:"email"`
	Picture         string   `json:"picture"`
	ApprovedClients []string `json:"approved_clients"`
}
