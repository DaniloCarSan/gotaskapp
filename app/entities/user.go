package entities

import (
	"gotaskapp/app/security"
	"strings"
	"time"
)

type User struct {
	Id        uint64    `json:"id,omitempty"`
	Firstname string    `json:"firstname,omitempty"`
	Lastname  string    `json:"lastname,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"-"`
	Verified  string    `json:"verified"`
	CreateAt  time.Time `json:"create_at,omitempty"`
}

// Convert password to hash
func (u *User) PasswordToHash() error {
	hash, err := security.PasswordToHash(u.Password)

	if err != nil {
		return err
	}

	u.Password = string(hash)

	return nil
}

// Verify if email is verified
func (u *User) IsEmailVerified() bool {
	return strings.EqualFold(u.Verified, "Y")
}
