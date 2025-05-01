package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type User struct {
    Name     string `json:"name"`
    Email    string `json:"email"`
}

type UserRequest struct {
    Name     string `json:"name"`
    Email    string `json:"email"`
    Password string `json:"password"` // Usado apenas para receber dados da requisição
}

func CreateUser(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var userReq UserRequest
        err := json.NewDecoder(r.Body).Decode(&userReq)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        // Inserir no banco (ajuste os nomes das colunas conforme seu banco)
        _, err = db.Exec(
            "INSERT INTO users (name, email, password_hash) VALUES ($1, $2, $3)",
            userReq.Name,
            userReq.Email,
            userReq.Password,
        )
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(map[string]string{"message": "Usuário criado com sucesso"})
    }
}

func GetUsers(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        rows, err := db.Query("SELECT name, email FROM users")
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        defer rows.Close()

        var users []User
        for rows.Next() {
            var u User
            err := rows.Scan(&u.Name, &u.Email)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            users = append(users, u)
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(users)
    }
}

func LoginSession(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var userReq UserRequest
        err := json.NewDecoder(r.Body).Decode(&userReq)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        // Aqui você deve implementar a lógica de autenticação
        // Exemplo: verificar se o email e senha estão corretos

        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(map[string]string{"message": "Login bem-sucedido"})
    }
}