package model

// User struct
type User struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Contact string `json:"contact"`
}

// Users map to store user data
var Users = map[string]User{
	"123": {
		ID:      "123",
		Name:    "Najeeb Ullah Shah",
		Contact: "03401243464",
	},
	"456": {
		ID:      "456",
		Name:    "Samreen Hassan",
		Contact: "03401243456",
	},
	"789": {
		ID:      "789",
		Name:    "Junaid King",
		Contact: "03401274264",
	},
}
