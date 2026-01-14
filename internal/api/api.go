package api

import (
	"context"
	"encoding/json"
	"encurtador/internal/cache"
	"encurtador/internal/shortener"
	"encurtador/internal/store"
	"log"
	"net/http"
	"time"
)

type RequestBody struct {
	URL string `json:"url"`
}

type ResponseBody struct {
	ShortUrl string `json:"short_url"`
	Code string `json:"code"` 
}

type Handler struct {
	store *store.Store
	cache *cache.Cache
}

func NewHandler(s *store.Store) *Handler {
	return &Handler{store: s}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var body RequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Erro: JSON inválido", http.StatusBadRequest)
		return
	}

	var code string
	var err error

	for i := 0; i < 6; i++ {
		code = shortener.Generate(6)
		log.Printf("Código gerado: %s\n", code)
		_, err = h.store.Save(body.URL, code)
		if err == nil {
			break
		}

		log.Printf("Colisão detectada ou erro de banco: %v.", err)
	}

	if err != nil {
		http.Error(w, "Erro: falha ao criar URL encurtada", http.StatusInternalServerError)
		return
	}

	response := ResponseBody {
		Code: code,
		ShortUrl: "http://localhost:8080/" + code,
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Path[1:]

	urlCache, err := h.cache.Get(r.Context(), code)
	if err == nil {
		http.Redirect(w, r, urlCache, http.StatusFound)
		return 
	}

	log.Println("Cache MISS: " + code)

	url, err := h.store.Get(code)
	if err != nil {
		http.Error(w, "Erro: URL não encontrada", http.StatusNotFound)
		return
	} 

	go func() {
		_ = h.cache.Save(context.Background(), code, url, 24*time.Hour)
	}()
	
	http.Redirect(w, r, url, http.StatusFound)
}