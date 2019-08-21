package main

// User - User struct to hold the User data
type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// JWT -
type JWT struct {
	Token string `json:"token"`
}

// Error - To assign error message to be sent to the client
type Error struct {
	Message string `json:"message"`
}

// To return new user
func newUser(id int, email string, password string) User {
	user := User{ID: id, Email: email, Password: password}
	return user
}
