package model

type Message struct {
	EmailMessage string
	PasswordMessage string
	PolicyMessage string
	LoginMessage string
}

type MessageLoggedIn struct {
	AlreadyLoggedIn bool
}