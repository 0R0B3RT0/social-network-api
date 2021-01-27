package routes

import (
	"net/http"

	"api/src/controllers"
)

var userRoutes = []Rota{
	{
		URI:                "/users",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CreateUser,
		RequerAutenticacao: false,
	},
	{
		URI:                "/users/{id}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.UpdateUser,
		RequerAutenticacao: false,
	},
	{
		URI:                "/users",
		Metodo:             http.MethodGet,
		Funcao:             controllers.FindUsers,
		RequerAutenticacao: false,
	},
	{
		URI:                "/users/{id}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.FindUser,
		RequerAutenticacao: false,
	},
	{
		URI:                "/users/{id}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeleteUser,
		RequerAutenticacao: false,
	},
}
