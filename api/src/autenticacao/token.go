package autenticacao

import (
	"DevBook/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// CriarToken cria um token JWT para um usuário
func CriarToken(usuarioID uint64) (string, error) {
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissoes["usuarioId"] = usuarioID

	//secret
	secret := config.SecretKey
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)
	tokenString, erro := token.SignedString(secret)
	if erro != nil {
		return "", erro
	}
	return tokenString, nil
}

// ValidarToken verifica se um token passado na requisição é válido
func ValidarToken(r *http.Request) error {
	tokenString, erro := extrairToken(r)
	if erro != nil {
		return erro
	}
	token, erro := jwt.Parse(tokenString, retornarChaveVerificacao)
	if erro != nil {
		return erro
	}
	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return errors.New("token inválido")
	}
	return nil
}

func extrairToken(r *http.Request) (string, error) {
	token := r.Header.Get("Authorization")
	if token == "" {
		return "", errors.New("token não informado")
	}
	if token[:7] != "Bearer " {
		return "", errors.New("token mal formatado")
	}
	return token[7:], nil
}

func retornarChaveVerificacao(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("método de assinatura inesperado: ", token.Header["alg"])
	}
	return config.SecretKey, nil
}

// ExtrairUsuarioId extrai o id do usuário do token
func ExtrairUsuarioId(r *http.Request) (uint64, error) {
	tokenString, erro := extrairToken(r)
	if erro != nil {
		return 0, erro
	}
	token, erro := jwt.Parse(tokenString, retornarChaveVerificacao)
	if erro != nil {
		return 0, erro
	}
	if permissoes, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		usuarioID, erro := strconv.ParseUint(fmt.Sprintf("%.0f", permissoes["usuarioId"]), 10, 64)
		if erro != nil {
			return 0, erro
		}
		return usuarioID, nil
	}
	return 0, errors.New("token inválido")
}
