package main

import (
	"fmt"
	"log"
	"net/http"

	"api/src/config"
	"api/src/router"
)

func main() {
	config.Load()

	fmt.Println("executando a API!!!")
	r := router.Generate()

	fmt.Printf("Escutando a porta %d...\n", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
