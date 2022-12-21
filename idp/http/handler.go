package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/knwoop/fedcm-example/idp/fedcm"
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
	cookie, errs := r.Cookie("SID")
	if errs != nil || cookie.Value == "" {
		return
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
	cookie, errs := r.Cookie("SID")
	if errs != nil || cookie.Value != "token" {
		return
	}

	fmt.Fprintf(w, "%s", "hello")
}

type metadataHandler struct {
}

func (h *metadataHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", "hello")
}

type assertionHandler struct {
}

func (h *assertionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", "hello")
}
