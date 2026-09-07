package modelos

import (
	"errors"
	"strings"
	"time"
)

// Publicacao representa uma publicação feita por um usuário.
type Publicacao struct {
	ID        int64     `json:"id,omitempty"`
	Titulo    string    `json:"titulo,omitempty"`
	Conteudo  string    `json:"conteudo,omitempty"`
	AutorID   uint64    `json:"autorId,omitempty"`
	AutorNick uint64    `json:"autorNick,omitempty"`
	Curtidas  uint64    `json:"curtidas"`
	CriadaEm  time.Time `json:"criadaEm,omitempty"`
}

// Prepara faz as validações iniciais da publicação.
func (publicacao *Publicacao) Preparar() error {
	if erro := publicacao.validar(); erro != nil {
		return erro
	}
	publicacao.formatar()
	return nil
}

// validar se os campos obrigatórios estão preenchidos.
func (publicacao *Publicacao) validar() error {
	if publicacao.Titulo == "" {
		return errors.New("titulo é obrigatório e não pode estar em branco")
	}
	if publicacao.Conteudo == "" {
		return errors.New("conteudo é obrigatório e não pode estar em branco")
	}
	return nil
}

// formatar faz o trim dos campos e remove espaços desnecessários.
func (publicacao *Publicacao) formatar() {
	publicacao.Titulo = strings.TrimSpace(publicacao.Titulo)
	publicacao.Conteudo = strings.TrimSpace(publicacao.Conteudo)
}
