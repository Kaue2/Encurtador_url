package main

import (
	"encurtador/internal/api"
	"encurtador/internal/store"
	"encurtador/internal/cache"
	"fmt"
	"net/http"
	"os"
	"log"

	"github.com/joho/godotenv"
)

func main()  {
	if err := godotenv.Load(); err != nil {
		log.Println("Aviso: não foi encontrado um arquivo .env")
	}

	// Construindo a string de conexão
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	db := "encurtador"
	host := "localhost"

	connStr := fmt.Sprintf("postgres://%s:%s@%s:5432/%s", user, password, host, db)
	fmt.Printf("String de conexão: %s\n", connStr)

	// Criando um storage
	storage, err := store.NewStore(connStr)
	if err != nil {
		log.Fatalf("Erro: falha ao conectar no banco: %v", err)
	}
	defer storage.Close() // agenda a liberação desse objeto na memória
	fmt.Println("Postgres conectado")

	redisClient, err := cache.NewCache("localhost:6379")
	if err != nil {
		log.Fatalf("Erro: falha ao conectar com o redis: %v", err)
	}
	defer redisClient.Close()
	fmt.Println("Redis conectado")
	
	handler := api.NewHandler(storage, redisClient)

	fmt.Println("Conexão estabelecida e tabelas criadas")

	router := handler.RegisterRoutes()

	port := ":8080"
	fmt.Printf("Subindo server...\n")

	server := &http.Server {
		Addr:  port,
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}