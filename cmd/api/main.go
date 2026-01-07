package main

import (
	"fmt"
	"net/http"
)

func main()  {
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