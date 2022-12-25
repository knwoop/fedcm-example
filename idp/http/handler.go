package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/knwoop/fedcm-example/idp/config"
	"github.com/knwoop/fedcm-example/idp/fedcm"
	"github.com/knwoop/fedcm-example/idp/token"
	"github.com/knwoop/fedcm-example/idp/user"
)

func (s *Server) Signin(w http.ResponseWriter, r *http.Request) {
	var req user.User
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := s.db.GetUserByUserName(r.Context(), req.Username)
	if err != nil {
		u = &req
		u.ID = uuid.NewString()
		u.Username = "no " + time.Now().Format("150405")
		u.Name = "foo"
		u.Email = fmt.Sprintf("%s@example.com", u.Username)
		_ = s.db.PutUser(r.Context(), u)
	}

	cookie := &http.Cookie{
		Name:     "SID",
		Value:    u.ID,
		MaxAge:   0,
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
	if err := json.NewEncoder(w).Encode(&u); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s *Server) SigninWithIDToken(w http.ResponseWriter, r *http.Request) {
	type request struct {
		IDToken string `json:"id_token"`
	}
	var req request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	sub, err := token.VerifyIDToken(req.IDToken)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	u, err := s.db.GetUserByID(r.Context(), sub)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cookie := &http.Cookie{
		Name:     "SID",
		Value:    u.ID,
		MaxAge:   0,
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
	if err := json.NewEncoder(w).Encode(&u); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s *Server) getMeHandler(w http.ResponseWriter, r *http.Request) {
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

	u, err := s.db.GetUserByID(r.Context(), cookie.Value)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(w).Encode(&u); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s *Server) GetWellKnownFileHandler(w http.ResponseWriter, r *http.Request) {
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

func (s *Server) GetConfigFileHandler(w http.ResponseWriter, r *http.Request) {
	m := &fedcm.Manifest{
		AccountsEndpoint:       "/fedcm/accounts_endpoint",
		ClientMetadataEndpoint: "/fedcm/client_metadata_endpoint",
		IDAssertionEndpoint:    "/fedcm/id_assertion_endpoint",
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

func (s *Server) AccountsHandler(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("SID")
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			http.Error(w, "cookie not found", http.StatusUnauthorized)
			return
		default:

			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}
	}

	us, err := s.db.GetAllUsers(r.Context())
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
	}

	accounts := make([]*fedcm.IdentityProviderAccount, len(us))
	for i, u := range us {
		accounts[i] = &fedcm.IdentityProviderAccount{
			ID:        u.ID,
			GivenName: u.Username,
			Name:      u.Name,
			Email:     u.Email,
			Picture:   u.Picture,
		}
	}

	ipAccounts := &fedcm.IdentityProviderAccounts{
		Accounts: accounts,
	}
	w.Header().Set("content-type", "application/json")
	if err := json.NewEncoder(w).Encode(&ipAccounts); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s *Server) MetadataHandler(w http.ResponseWriter, r *http.Request) {
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

func (s *Server) AssertionHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
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

	u, _ := s.db.GetUserByID(r.Context(), assertion.AccountID)

	token, err := token.GenereateIDToken("idp", u.ID)
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
