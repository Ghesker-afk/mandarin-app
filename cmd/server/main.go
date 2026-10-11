package main

import (
	"context"
	"fmt"
	"log"
	"mandarin-app/internal/database"
	"mandarin-app/internal/handlers"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = "8080"
	}

	fmt.Println("Iniciando o servidor da aplicação Mandarim...")

	// Conecta ao banco de dados e executa migrações
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

	// Configura os Handlers e Rotas
	ideogramHandler := handlers.NewIdeogramHandler(conn)
	mux := http.NewServeMux()

	mux.HandleFunc("/ideograms", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			ideogramHandler.CreateIdeogram(w, r)
		case http.MethodGet:
			ideogramHandler.ListIdeograms(w, r)
		default:
			http.Error(w, "Método não suportado", http.StatusMethodNotAllowed)
		}
	})

	fmt.Printf("Servidor HTTP rodando na porta %s...\n", serverPort)
	if err := http.ListenAndServe(":"+serverPort, mux); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v\n", err)
	}
}
