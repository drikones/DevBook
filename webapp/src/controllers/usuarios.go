package controllers

import (
	"io"
	"net/http"
)

const urlAPI = "http://localhost:8000"

// CriarUsuario repassa o cadastro de usuário para a API e devolve a resposta
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	resposta, erro := http.Post(urlAPI+"/usuarios", "application/json", r.Body)
	if erro != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"erro":"Não foi possível conectar à API"}`))
		return
	}
	defer resposta.Body.Close()

	w.Header().Set("Content-Type", resposta.Header.Get("Content-Type"))
	w.WriteHeader(resposta.StatusCode)
	io.Copy(w, resposta.Body)
}
