package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/knwoop/fedcm-example/idp/user"
)

var (
	_ http.Handler = (*getUserHandler)(nil)
	_ http.Handler = (*listUserHandler)(nil)
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
