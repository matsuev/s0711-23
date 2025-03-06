package main

// UserModel ...
type UserModel struct {
	ID       int    `json:"id,omitempty"`
	Username string `json:"username" validate:"required,ascii"`
}

// ProfileModel ...
type ProfileModel struct {
	ID          int
	GivenName   string `json:"given_name,omitempty" validate:"omitempty,alphaunicode"`
	FamilyName  string `json:"family_name,omitempty" validate:"omitempty,alphaunicode"`
	Email       string
	BirthDate   string
	PhoneNumber string
}
