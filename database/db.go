package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func Connect(host, port, user, password, dbname string) *sql.DB {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Erro ao abrir conexão com o banco:", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	if err := db.PingContext(ctx); err != nil {
		log.Fatal("Erro ao conectar ao banco (timeout):", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	log.Println("Conectado ao banco com sucesso")
	return db
}