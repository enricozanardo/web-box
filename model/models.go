package model

type Message struct {
	EmailMessage string
	PasswordMessage string
	PolicyMessage string
	LoginMessage string
	Allowed bool
}

type MessageLoggedIn struct {
	AlreadyLoggedIn bool
}