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

	if err := godotenv.Load(); err != nil {
		fmt.Println("Arquivo .env não encontrado; usando variáveis de ambiente")
	}

	ctx := context.Background()

	dbConn, err := database.ConnectDB(ctx)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer dbConn.Close(ctx)

	// 3. Montando a Clean Architecture: Infra -> Usecase -> Handler
	ideogramRepo := postgres.NewPostgresIdeogramRepository(dbConn)
	translateUC := usecase.NewTranslateIdeogramUseCase(ideogramRepo)
	ideogramHandler := handlers.NewIdeogramHandler(translateUC)

	// 4. Configura as Rotas usando o net/http moderno (Go 1.22+)
	mux := http.NewServeMux()

	// Rota de tradução estruturada
	mux.HandleFunc("GET /translate", ideogramHandler.TranslateHandler)

	createIdeogramUseCase := usecase.NewCreateIdeogramUseCase(ideogramRepo)
	createIdeogramHandler := handlers.NewCreateIdeogramHandler(createIdeogramUseCase)
	mux.Handle("POST /ideograms", createIdeogramHandler)

	fmt.Printf("Servidor HTTP rodando na porta %s...\n", serverPort)
	if err := http.ListenAndServe(":"+serverPort, mux); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v\n", err)
	}
}
