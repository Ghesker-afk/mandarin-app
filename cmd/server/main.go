package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"mandarin-app/internal/database"
	"mandarin-app/internal/handlers"
	"mandarin-app/internal/repository/postgres"
	"mandarin-app/internal/usecase"
)

func main() {
	_ = godotenv.Load()
	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = "8081"
	}

	fmt.Println("Iniciando o servidor da aplicação Mandarim...")

	// 1. Conecta ao banco de dados (retorna o *pgxpool.Pool)
	dbPool := database.ConnectDB()
	defer dbPool.Close(context.Background())

	ctx := context.Background()

	// 3. Montando a Clean Architecture: Infra -> Usecase -> Handler
	ideogramRepo := postgres.NewPostgresIdeogramRepository(dbPool)
	translateUC := usecase.NewTranslateIdeogramUseCase(ideogramRepo)
	ideogramHandler := handlers.NewIdeogramHandler(translateUC)

	// 4. Configura as Rotas usando o net/http moderno (Go 1.22+)
	mux := http.NewServeMux()

	// Rota de tradução estruturada
	mux.HandleFunc("GET /translate", ideogramHandler.TranslateHandler)

	fmt.Printf("Servidor HTTP rodando na porta %s...\n", serverPort)
	if err := http.ListenAndServe(":"+serverPort, mux); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v\n", err)
	}
}
