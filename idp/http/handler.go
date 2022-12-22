package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/knwoop/fedcm-example/idp/config"
	"github.com/knwoop/fedcm-example/idp/fedcm"
	"github.com/knwoop/fedcm-example/idp/token"
	"github.com/knwoop/fedcm-example/idp/user"
)

var (
	_ http.Handler = (*listUserHandler)(nil)
	_ http.Handler = (*getUserHandler)(nil)
	_ http.Handler = (*getMeHandler)(nil)

	// fedCM endpoints below
	_ http.Handler = (*getWellKnownFile)(nil)
	_ http.Handler = (*accountsHandler)(nil)
	_ http.Handler = (*metadataHandler)(nil)
	_ http.Handler = (*assertionHandler)(nil)
)

type getUserHandler struct {
}

func (h *getUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := &user.User{
		ID:       1,
		Username: "user1",
		Email:    "user1@examole.com",
	}

	if err := json.NewEncoder(w).Encode(&user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

type listUserHandler struct {
}

func (h *listUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	users := []*user.User{
		{
			ID:       1,
			Username: "user1",
			Email:    "user1@examole.com",
		},
		{
			ID:       2,
			Username: "user2",
			Email:    "user2@examole.com",
		},
		{
			ID:       3,
			Username: "user3",
			Email:    "user3@examole.com",
		},
	}

	if err := json.NewEncoder(w).Encode(&users); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

type signinHandler struct {
}

func (h *signinHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login!!")
	cookie := &http.Cookie{
		Name:     "SID",
		Value:    "token",
		MaxAge:   0,
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
	user := &user.User{
		ID:       1,
		Username: "user1",
		Email:    "user1@examole.com",
	}

	if err := json.NewEncoder(w).Encode(&user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

type getMeHandler struct {
}

func (h *getMeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("SID")
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			http.Error(w, "cookie not found", http.StatusUnauthorized)
			return
		default:
			log.Println(err)
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}
	}

	fmt.Println("meeeeeeeeeee", cookie.Value)
}

type getWellKnownFile struct {
}

func (h *getWellKnownFile) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	wi := &fedcm.WebIdentity{
		ProviderURLs: []string{
			"http://localhost:8080/config.json",
		},
	}

	w.Header().Set("content-type", "application/json")

	if err := json.NewEncoder(w).Encode(&wi); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

type getConfigFile struct {
}

func (h *getConfigFile) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m := &fedcm.Manifest{
		AccountsEndpoint:       "/accounts",
		ClientMetadataEndpoint: "/metadata",
		IDAssertionEndpoint:    "/assertion",
		Branding: &fedcm.Branding{
			BackgroundColor: "green",
			Color:           "0xFFEEAA",
			Icons: []*fedcm.Icon{
				{
					Url:  "https://idp.example/icon.ico",
					Size: 25,
				},
			},
		},
	}

	w.Header().Set("content-type", "application/json")

	if err := json.NewEncoder(w).Encode(&m); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

type accountsHandler struct {
}

func (h *accountsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("SID")
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			http.Error(w, "cookie not found", http.StatusUnauthorized)
			return
		default:
			log.Println(err)
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}
	}

	accounts := &fedcm.IdentityProviderAccounts{
		Accounts: []*fedcm.IdentityProviderAccount{
			{
				ID:        "knwoop",
				GivenName: "kenwoo",
				Name:      "Kenta Takahashi",
				Email:     "knwoop@gmail.com",
				Picture:   "https://avatars.githubusercontent.com/u/13586089?s=200",
			},
		},
	}
	w.Header().Set("content-type", "application/json")
	if err := json.NewEncoder(w).Encode(&accounts); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

type metadataHandler struct {
}

func (h *metadataHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("meta!!")
	m := &fedcm.IdentityProviderClientMetadata{
		PrivacyPolicyURL:  config.PrivacyPolicyURL,
		TermsOfServiceURL: config.TermsOfServiceURL,
	}

	w.Header().Set("content-type", "application/json")
	if err := json.NewEncoder(w).Encode(&m); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

type assertionHandler struct {
}

func (h *assertionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for k, v := range r.Form {
		fmt.Printf("%v : %v\n", k, v)
	}
	accountID := r.Form.Get("account_id")
	clientID := r.Form.Get("client_id")
	nonce := r.Form.Get("nonce")
	d := r.Form.Get("disclosure_text_shown")
	disclosureTextShown, err := strconv.ParseBool(d)
	if err != nil {
		http.Error(w, "invalid bool", http.StatusBadRequest)
		return
	}

	assertion := &fedcm.IdentityAssertion{
		AccountID:           accountID,
		ClientID:            clientID,
		Nonce:               nonce,
		DisclosureTextShown: disclosureTextShown,
	}

	_ = assertion

	token, err := token.GenereateIDToken("idp", "knwoop")
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to generate token: %s", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	if err := json.NewEncoder(w).Encode(&struct {
		Tokne string `json:"token"`
	}{
		Tokne: string(token),
	}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
