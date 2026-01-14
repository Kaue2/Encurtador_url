# --- Estágio 1: Builder ---
FROM golang:1.23-alpine AS builder

# Instala git
RUN apk add --no-cache git

WORKDIR /app

# Copia os arquivos de dependência primeiro
COPY go.mod go.sum ./
RUN go mod download

# Copia o código fonte
COPY . .

# Compila o binário
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api

# --- Estágio 2: Runner ---
FROM alpine:latest

WORKDIR /root/

# Instala certificados CA
RUN apk --no-cache add ca-certificates

# Copia APENAS o binário do estágio anterior
COPY --from=builder /app/main .

# Copia o .env
# COPY .env . 

# Expõe a porta
EXPOSE 8080

# Comando para rodar
CMD ["./main"]