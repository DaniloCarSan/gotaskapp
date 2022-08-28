package entities

type Task struct {
	ID          uint64 `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
	StatusID    uint64 `json:"status_id,omitempty"`
	UserID      uint64 `json:"user_id,omitempty"`
}
