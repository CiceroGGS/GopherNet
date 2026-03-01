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
		URI:                    "/usuarios/{userId}",
		Method:                 http.MethodGet,
		Function:               controllers.FindUserByID,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/usuarios/{userId}",
		Method:                 http.MethodPut,
		Function:               controllers.UpdateUser,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/usuarios/{userId}",
		Method:                 http.MethodDelete,
		Function:               controllers.RemoveUser,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/usuarios/{userId}/seguir",
		Method:                 http.MethodPost,
		Function:               controllers.FollowUser,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/usuarios/{userId}/parar-de-seguir",
		Method:                 http.MethodPost,
		Function:               controllers.UnfollowUser,
		RequiresAuthentication: true,
	},
}
