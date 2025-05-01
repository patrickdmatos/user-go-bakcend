# Etapa base com Go
FROM golang:latest

# Diretório de trabalho dentro do container
WORKDIR /app

# Instala o Air (ferramenta de live reload)
RUN go install github.com/air-verse/air@latest

# Copia os arquivos go.mod e go.sum e baixa dependências
COPY go.mod go.sum ./
RUN go mod download

# Copia o restante do código para dentro do container
COPY . .

# Expõe a porta padrão do servidor Go
EXPOSE 8080

# Comando para rodar o servidor com reload automático
CMD ["air"]
