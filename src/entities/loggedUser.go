package entities

type LoggedUser struct {
	User
	Token string `json:"token"`
}
