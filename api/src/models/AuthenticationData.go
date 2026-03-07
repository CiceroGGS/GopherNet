package models

// AuthenticationData contem os dados do ID e do token do usuario autenticado
type AuthenticationData struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}
