package main

import (
	"context"
	"fmt"
	"mandarin-app/internal/database"
)

func main() {
	fmt.Println("Iniciando o servidor da aplicação Mandarim...")

	conn := database.ConnectDB()
	defer conn.Close(context.Background())

	fmt.Println("Aplicação pronta e rodando com Go e PostgreSQL!")
}
