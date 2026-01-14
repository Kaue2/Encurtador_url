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

	dbHost := os.Getenv("DB_HOST")
    if dbHost == "" {
        dbHost = "localhost" // Fallback para rodar fora do docker
    }
    
    dbPort := os.Getenv("DB_PORT")
    if dbPort == "" {
        dbPort = "5432"
    }
    
    // Redis Host também
    redisHost := os.Getenv("REDIS_HOST")
    if redisHost == "" {
        redisHost = "localhost:6379"
    }

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
        os.Getenv("POSTGRES_USER"),
        os.Getenv("POSTGRES_PASSWORD"),
        dbHost,
        dbPort,
        os.Getenv("POSTGRES_DB"),
    )
	fmt.Printf("String de conexão: %s\n", connStr)

	// Criando um storage
	storage, err := store.NewStore(connStr)
	if err != nil {
		log.Fatalf("Erro: falha ao conectar no banco: %v", err)
	}
	defer storage.Close() // agenda a liberação desse objeto na memória
	fmt.Println("Postgres conectado")

	redisClient, err := cache.NewCache(redisHost)
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