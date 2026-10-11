package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func ConnectDB() *pgx.Conn {

	_ = godotenv.Load()

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", user, password, host, port, dbname, sslmode)

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v\n", err)
	}

	fmt.Println("Conexão com o PostgreSQL estabelecida com sucesso!")
	return conn
}
