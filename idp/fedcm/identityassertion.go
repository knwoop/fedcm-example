package fedcm

type IdentityAssertion struct {
	AccountID           string
	ClientID            string
	Nonce               string
	DisclosureTextShown bool
}
