package modelos

import (
	"DevBook/src/seguranca"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// Usuario representa um usuario utilizando a rede social
type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criadoEm,omitempty"`
}

// Preparar vai chamar os métodos para validar e formatar o usuário recebido
func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}
	if erro := usuario.formatar(etapa); erro != nil {
		return erro
	}
	return nil
}

func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("nome é obrigatório e não pode estar em branco")
	}
	if usuario.Nick == "" {
		return errors.New("nick é obrigatório e não pode estar em branco")
	}
	if usuario.Email == "" {
		return errors.New("email é obrigatório e não pode estar em branco")
	}
	if erro := checkmail.ValidateFormat(usuario.Email); erro != nil {
		return errors.New("email inserido está no formato inválido")
	}
	if usuario.Senha == "" && etapa == "cadastro" {
		return errors.New("senha é obrigatória e não pode estar em branco")
	}
	return nil
}

func (usuario *Usuario) formatar(etapa string) error {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)

	if etapa == "cadastro" {
		senhaComHash, erro := seguranca.Hash(usuario.Senha)
		if erro != nil {
			panic(erro)
		}
		usuario.Senha = string(senhaComHash)
	}
	return nil
}
