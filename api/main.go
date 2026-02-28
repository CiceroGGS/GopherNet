package main

import (
	"fmt"
	"gophernet/src/config"
	"gophernet/src/router"
	"log"
	"net/http"
)

func main() {
	config.Load()
	fmt.Printf("API rodando na porta: %d\n", config.Port)

	r := router.Generate()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
