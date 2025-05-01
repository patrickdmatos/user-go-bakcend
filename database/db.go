package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	// Validação das variáveis de ambiente
	requiredEnvVars := []string{"DB_HOST", "DB_PORT", "DB_USER", "DB_PASSWORD", "DB_NAME"}
	for _, envVar := range requiredEnvVars {
		if os.Getenv(envVar) == "" {
			log.Fatalf("Variável de ambiente %s não definida", envVar)
		}
	}

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	// Configuração do pool de conexões
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Erro ao abrir conexão com o banco:", err)
	}

	// Testa a conexão com timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	if err := db.PingContext(ctx); err != nil {
		log.Fatal("Erro ao conectar ao banco (timeout):", err)
	}

	// Configurações do pool de conexões
	db.SetMaxOpenConns(25)          // Número máximo de conexões abertas
	db.SetMaxIdleConns(25)          // Número máximo de conexões ociosas
	db.SetConnMaxLifetime(5 * time.Minute) // Tempo máximo de vida da conexão

	log.Println("Conectado ao banco com sucesso")
	return db
}