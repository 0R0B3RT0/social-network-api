package controllers

import "net/http"

// CreateUser insere um usuário no banco de dados
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando Usuário!"))
}

// UpdateUser atualiza um usuário no banco
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizar Usuário!"))
}

// FindUser busta todos os utuários do banco
func FindUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando Usuários!"))
}

// FindUsers busca um usuário no banco
func FindUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando Usuário!"))
}

// DeleteUser remove um usuário do bando
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando Usuário!"))
}
