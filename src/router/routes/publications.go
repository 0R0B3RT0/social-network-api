package routes

import (
	"api/src/controllers"
	"net/http"
)

var publicationRoutes = []Route{
	{
		URI:                "/publications",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CreatePublication,
		RequerAutenticacao: true,
	},
	{
		URI:                "/publications/{id}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.UpdatePublication,
		RequerAutenticacao: true,
	},
	{
		URI:                "/publications",
		Metodo:             http.MethodGet,
		Funcao:             controllers.FindPublications,
		RequerAutenticacao: true,
	},
	{
		URI:                "/publications/{id}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.FindPublication,
		RequerAutenticacao: true,
	},
	{
		URI:                "/publications/{id}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletePublication,
		RequerAutenticacao: true,
	},
}
