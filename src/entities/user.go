package entities

import (
	"gotaskapp/src/security"
	"time"
)

type User struct {
	Id        uint64    `json:"id,omitempty"`
	Firstname string    `json:"firstname,omitempty"`
	Lastname  string    `json:"lastname,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"-"`
	CreateAt  time.Time `json:"-"`
}

func (u *User) PasswordToHash() error {
	hash, err := security.PasswordToHash(u.Password)

	if err != nil {
		return err
	}

	u.Password = string(hash)

	return nil
}
