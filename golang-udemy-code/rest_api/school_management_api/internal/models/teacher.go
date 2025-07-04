package models

type Teacher struct {
	ID        int    `json:"id,omitempty" db:"id,omitempty"`
	FirstName string `json:"first_name,omitempty" db:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty" db:"last_name,omitempty"`
	Class     string `json:"class,omitempty" db:"email,omitempty"`
	Subject   string `json:"subject,omitempty" db:"class,omitempty"`
	Email     string `json:"email,omitempty" db:"subject,omitempty"`
}
