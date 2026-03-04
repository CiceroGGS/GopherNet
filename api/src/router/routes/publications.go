package routes

import (
	"gophernet/src/controllers"
	"net/http"
)

var publicationsRoutes = []Routes{
	{
		URI:                    "/publicacoes",
		Method:                 http.MethodPost,
		Function:               controllers.CreatePublication,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/publicacoes",
		Method:                 http.MethodGet,
		Function:               controllers.SearchPublications,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/publicacoes/{publicacaoId}",
		Method:                 http.MethodGet,
		Function:               controllers.FindPublicationByID,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/publicacoes/{publicacaoId}",
		Method:                 http.MethodPut,
		Function:               controllers.UpdatePublication,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/publicacoes/{publicacaoId}",
		Method:                 http.MethodDelete,
		Function:               controllers.DeletePublication,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/usuarios/{usuarioId}/publicacoes",
		Method:                 http.MethodGet,
		Function:               controllers.SearchPublicationsByUser,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/publicacoes/{publicacaoId}/curtir",
		Method:                 http.MethodPost,
		Function:               controllers.LikePublication,
		RequiresAuthentication: true,
	},

	{
		URI:                    "/publicacoes/{publicacaoId}/tirar-curtida",
		Method:                 http.MethodPost,
		Function:               controllers.DislikePublication,
		RequiresAuthentication: true,
	},
}
