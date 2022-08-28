package entities

type Status struct {
	ID     uint64 `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	UserID uint64 `json:"user_id,omitempty"`
}
