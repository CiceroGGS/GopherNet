package routes

import (
	"gophernet/src/controllers"
	"net/http"
)

var loginRoute = Routes{
	URI:                    "/login",
	Method:                 http.MethodPost,
	Function:               controllers.Login,
	RequiresAuthentication: false,
}
