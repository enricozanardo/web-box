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

type EmergencyMessage struct {
	EmergencyAddressMessage string
	EmergencyNumberMessage string
	EergencyPostalCodeMessage string
	EmergencyPlaceMessage string
	EmergencySuccessMessage string
	EmergencyErrorMessage string
}