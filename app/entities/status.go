package entities

type Status struct {
	Id     uint64 `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	UserId uint64 `json:"user_id,omitempty"`
}
