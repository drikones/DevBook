package middlewares

import (
	"DevBook/src/autenticacao"
	"DevBook/src/respostas"
	"log"
	"net/http"
)

// Logger é um middleware que registra as requisições
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.URL.Path, r.RemoteAddr)
		next(w, r)
	}
}

// Autenticar é um middleware que verifica se o usuário está autenticado
func Autenticar(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if erro := autenticacao.ValidarToken(r); erro != nil {
			respostas.Erro(w, http.StatusUnauthorized, erro)
			return
		}
		next(w, r)
	}
}
