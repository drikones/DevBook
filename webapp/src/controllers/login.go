package controllers

import (
	"net/http"
	"webapp/src/respostas"
)

// FazerLogin repassa o e-mail e a senha do usuário para a API e devolve a resposta
func FazerLogin(w http.ResponseWriter, r *http.Request) {
	response, erro := http.Post(urlAPI+"/login", "application/json", r.Body)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: "Não foi possível conectar à API"})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	respostas.JSON(w, response.StatusCode, nil)
}
