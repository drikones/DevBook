package main

import (
	"DevBook/src/config"
	"DevBook/src/router"
	"fmt"
	"log"
	"net/http"
)

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
func main() {
	config.Carregar()
	fmt.Println(config.Porta)
	fmt.Println(config.StringConexaoBanco)
	fmt.Println("Rodando API")
	r := router.Gerar()
	log.Fatal(http.ListenAndServe(":8000", r))
}
