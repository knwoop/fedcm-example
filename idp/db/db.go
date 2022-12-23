package db

import (
	"context"
	"errors"
	"sync"

	"github.com/knwoop/fedcm-example/idp/user"
)

var (
	ErrNotFound = errors.New("not found")
)

type DB struct {
	users map[string]*user.User
	lock  sync.RWMutex
}

func NewDB() *DB {
	return &DB{users: map[string]*user.User{
		"ebd46852-4df1-4732-909c-6f7a63bac241": {
			ID:       "ebd46852-4df1-4732-909c-6f7a63bac241",
			Username: "kenwoo",
			Name:     "Kenta Takahashi",
			Email:    "knwoop@example.com",
			Picture:  "https://avatars.githubusercontent.com/u/13586089?s=200",
		},
		"1d4dff77-4fac-4476-bc57-38beec692d02": {
			ID:       "1d4dff77-4fac-4476-bc57-38beec692d02",
			Username: "kenwoo-work",
			Name:     "Kenta Takahashi(work)",
			Email:    "knwoopwork@example.com",
			Picture:  "https://avatars.githubusercontent.com/u/13586089?s=200",
		},
	}}
}

func (d *DB) GetUserByID(ctx context.Context, id string) (*user.User, error) {
	d.lock.RLock()
	defer d.lock.RUnlock()

	user, ok := d.users[id]
	if !ok {
		return nil, ErrNotFound
	}

	return user, nil
}

func (d *DB) GetUserByUserName(ctx context.Context, username string) (*user.User, error) {
	d.lock.RLock()
	defer d.lock.RUnlock()

	for _, u := range d.users {
		if u.Username == username {
			return u, nil
		}
	}

	return nil, ErrNotFound
}

func (d *DB) GetAllUsers(ctx context.Context) ([]*user.User, error) {
	d.lock.RLock()
	defer d.lock.RUnlock()

	users := make([]*user.User, len(d.users))
	i := 0
	for _, u := range d.users {
		users[i] = u
		i++
	}

	return users, nil
}

func (d *DB) PutUser(ctx context.Context, u *user.User) error {
	d.lock.Lock()
	d.users[u.ID] = u
	d.lock.Unlock()

	return nil
}
