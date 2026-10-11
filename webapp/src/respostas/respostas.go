package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

// ErroAPI representa a resposta de erro retornada pela API
type ErroAPI struct {
	Erro string `json:"erro"`
}

// JSON retorna uma resposta em formato JSON para a requisição
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if dados == nil {
		return
	}
	if erro := json.NewEncoder(w).Encode(dados); erro != nil {
		log.Println(erro)
	}
}

// TratarStatusCodeDeErro trata as respostas da API com status code 400 ou superior,
// repassando o status e a mensagem de erro para o navegador
func TratarStatusCodeDeErro(w http.ResponseWriter, r *http.Response) {
	var erro ErroAPI
	if err := json.NewDecoder(r.Body).Decode(&erro); err != nil || erro.Erro == "" {
		erro.Erro = http.StatusText(r.StatusCode)
	}
	JSON(w, r.StatusCode, erro)
}
