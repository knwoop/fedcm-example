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

type getMeHandler struct {
}

func (h *getMeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", "hello")
}

type getWellKnownFile struct {
}

func (h *getWellKnownFile) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

	if err := json.NewEncoder(w).Encode(&m); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

type accountsHandler struct {
}

func (h *accountsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
