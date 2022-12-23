package token

import (
	"crypto/ed25519"
	"crypto/rand"
	"fmt"
	"time"

	"github.com/knwoop/fedcm-example/idp/config"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/lestrrat-go/jwx/v2/jwt/openid"
)

var (
	publicKey  jwk.Key
	privateKey ed25519.PrivateKey
)

func init() {
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}

	privateKey = priv

	publicKey, err = jwk.FromRaw(pub)
	if err != nil {
		panic(fmt.Sprintf("failed to create JWK key: %s", err))
	}
}

func GenereateIDToken(iss, sub string) ([]byte, error) {
	const aLongLongTimeAgo = 233431200

	t := openid.New()
	t.Set(jwt.IssuerKey, iss)
	t.Set(jwt.SubjectKey, sub)
	t.Set(jwt.AudienceKey, config.SampleClinetID)
	t.Set(jwt.ExpirationKey, time.Now().Add(time.Hour).Unix())
	t.Set(jwt.IssuedAtKey, time.Unix(aLongLongTimeAgo, 0))

	signed, err := jwt.Sign(t, jwt.WithKey(jwa.SignatureAlgorithm("EdDSA"), privateKey))
	if err != nil {
		return nil, fmt.Errorf("failed to sign token: %w", err)
	}
	return signed, nil
}

func VerifyIDToken(idtoken string) (string, error) {
	t, err := jwt.Parse([]byte(idtoken), jwt.WithKey(jwa.SignatureAlgorithm("EdDSA"), publicKey))
	if err != nil {
		return "", err
	}

	return t.Subject(), nil
}
