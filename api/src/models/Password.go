package models

// Password representa o corpo da requisicao para atuailizar a senha de um usuario
type Password struct {
	CurrentPassword string `json:"senha_atual"`
	NewPassword     string `json:"senha_nova"`
}
