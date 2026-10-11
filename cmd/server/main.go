package main

import (
	"context"
	"fmt"
	"log"
	"mandarin-app/internal/database"
)

func main() {
	fmt.Println("Iniciando o servidor da aplicação Mandarim...")

	conn := database.ConnectDB()
	defer conn.Close(context.Background())

	ctx := context.Background()

	query := `
	CREATE TABLE IF NOT EXISTS ideograms (
		id SERIAL PRIMARY KEY,
		character VARCHAR(10) NOT NULL,
		pinyin VARCHAR(50) NOT NULL,
		meaning TEXT NOT NULL,
		example_sentence TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err := conn.Exec(ctx, query)
	if err != nil {
		log.Fatalf("Erro ao criar a tabela ideograms: %v\n", err)
	}

	fmt.Println("Tabela 'ideograms' verificada/criada com sucesso!")
}
