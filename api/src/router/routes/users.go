package routes

import (
	"gophernet/src/controllers"
	"net/http"
)

var usersRoutes = []Routes{
	{
		URI:                    "/usuarios",
		Method:                 http.MethodPost,
		Function:               controllers.CreateUser,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/usuarios",
		Method:                 http.MethodGet,
		Function:               controllers.SearchUsers,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/usuarios/{id}",
		Method:                 http.MethodGet,
		Function:               controllers.FindUserByID,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/usuarios/{id}",
		Method:                 http.MethodPut,
		Function:               controllers.UpdateUser,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/usuarios/{id}",
		Method:                 http.MethodDelete,
		Function:               controllers.RemoveUser,
		RequiresAuthentication: true,
	},
}
