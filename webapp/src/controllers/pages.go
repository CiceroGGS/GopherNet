package controllers

import (
	"net/http"
	"webapp/src/utils"
)

// LoadLoginScreen carrega a tela de login
func LoadLoginScreen(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "login.html", nil)
}

// LoadUserRegistrationPage carrega a tela de cadastro de usuario
func LoadUserRegistrationPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "register.html", nil)
}
