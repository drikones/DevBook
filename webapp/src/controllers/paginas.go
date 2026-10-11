package controllers

import (
	"net/http"
	"webapp/src/utils"
)

// CarregarTelaLogin renderiza a tela de login
func CarregarTelaLogin(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "login.html", nil)
}

// CarregarPaginaDeCadastro renderiza a tela de cadastro de usuário
func CarregarPaginaDeCadastro(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "cadastro.html", nil)
}
