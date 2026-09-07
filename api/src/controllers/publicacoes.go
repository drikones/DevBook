package controllers

import (
	"DevBook/src/autenticacao"
	"DevBook/src/banco"
	"DevBook/src/modelos"
	"DevBook/src/repositorios"
	"DevBook/src/respostas"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// CriarPublicacao insere uma publicação no banco de dados.
func CriarPublicacao(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var publicacao modelos.Publicacao
	if erro = json.Unmarshal(corpoRequest, &publicacao); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	autorID, erro := autenticacao.ExtrairUsuarioId(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}
	publicacao.AutorID = autorID

	if erro = publicacao.Preparar(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePublicacoes(db)
	publicacaoID, erro := repositorio.Criar(publicacao)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	publicacao.ID = publicacaoID
	respostas.JSON(w, http.StatusCreated, publicacao)

}

func BuscarPublicacoes(writer http.ResponseWriter, request *http.Request) {

}

func BuscarPublicacao(writer http.ResponseWriter, request *http.Request) {

}

func AtualizarPublicacao(writer http.ResponseWriter, request *http.Request) {

}

func DeletarPublicacao(writer http.ResponseWriter, request *http.Request) {

}
