package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func ConnectDB() *pgx.Conn {
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		connStr = "postgres://postgres:postgres@localhost:5433/mandarindb?sslmode=disable"
	}

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v\n", err)
	}

	fmt.Println("Conexão com o PostgreSQL estabelecida com sucesso!")
	return conn
}
