package entities

type Credential struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}
