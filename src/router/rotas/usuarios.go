package rotas

import (
	"net/http"

	"api/src/controllers"
)

var userRoutes = []Rota{
	{
		URI:                "/usuarios",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CreateUser,
		RequerAutenticacao: false,
	},
	{
		URI:                "/usuarios/{id}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.UpdateUser,
		RequerAutenticacao: false,
	},
	{
		URI:                "/usuarios",
		Metodo:             http.MethodGet,
		Funcao:             controllers.FindUser,
		RequerAutenticacao: false,
	},
	{
		URI:                "/usuarios/{id}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.FindUsers,
		RequerAutenticacao: false,
	},
	{
		URI:                "/usuarios/{id}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeleteUser,
		RequerAutenticacao: false,
	},
}
