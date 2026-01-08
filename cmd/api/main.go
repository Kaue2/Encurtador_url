package main

import (
	"encurtador/internal/store"
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
	defer storage.Close()

	fmt.Println("Conexão estabelecida e tabelas criadas")

	// Definindo rota simples
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Encurtador funcionando"))
	})

	port := ":8080"
	fmt.Printf("Subindo server...\n")

	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}