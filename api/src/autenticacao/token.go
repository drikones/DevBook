package autenticacao

import (
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
	secret := []byte("secret")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)
	tokenString, erro := token.SignedString(secret)
	if erro != nil {
		return "", erro
	}
	return tokenString, nil
}
