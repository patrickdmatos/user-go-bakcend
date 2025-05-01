package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/patrickdmatos/user-go-bakcend/database"
)

func main() {
	// Carrega .env (opcional)
	if err := godotenv.Load(); err != nil {
		log.Println("Arquivo .env não encontrado. Usando variáveis de ambiente.")
	}

	// Conexão com o banco (com tratamento de erro)
	db := database.Connect(
	os.Getenv("DB_HOST"),
	os.Getenv("DB_PORT"),
	os.Getenv("DB_USER"),
	os.Getenv("DB_PASSWORD"),
	os.Getenv("DB_NAME"),
)
    defer db.Close()

	// Configuração do router
	r := mux.NewRouter()

	// Middlewares
	r.Use(loggingMiddleware)
	r.Use(mux.CORSMethodMiddleware(r)) // CORS básico

	// Configuração do servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if _, err := strconv.Atoi(port); err != nil {
		log.Fatal("PORT deve ser um número")
	}

	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Servidor rodando na porta", port)
	log.Fatal(srv.ListenAndServe())
}

// Middleware de logging
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s %s", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}