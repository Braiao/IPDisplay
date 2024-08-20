package main

import (
	"fmt"
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Início")

	aplicacao := app.Gerar()
	if erro := aplicacao.Run(os.Args); erro != nil { //faz os comandos da linha de comando serem reconhecidos
		log.Fatal(erro)
	}

}
