package main

import (
	"fmt"

	"github.com/hrqcds/go-rest-api/database"
	"github.com/hrqcds/go-rest-api/routes"
)

func main() {

	database.DB_Conn()
	fmt.Println("Iniciando o servidor na porta 8000")
	routes.HandleRequest()
}
