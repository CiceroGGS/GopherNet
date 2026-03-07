package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var usersRoutes = []Route{
	{
		URI:               "/criar-usuario",
		Method:            http.MethodGet,
		Function:          controllers.LoadUserRegistrationPage,
		reqAuthentication: false,
	},
	{
		URI:               "/usuarios",
		Method:            http.MethodPost,
		Function:          controllers.CreateUser,
		reqAuthentication: false,
	},
}
