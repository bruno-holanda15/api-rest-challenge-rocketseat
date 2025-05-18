package entities

type User struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Biography string `json:"biography,omitempty"`
}
