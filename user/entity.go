package user

import "time"

// struct dibawah ini mewakili data yg ada di database kita
// ini diibaratkan sebagai model



type User struct {
	// fields here
	ID             int
	Name           string
	Occupation     string
	Email          string
	PasswordHash   string
	AvatarFileName string
	Role           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}