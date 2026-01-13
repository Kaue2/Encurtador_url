package api

import "net/http"

func (h *Handler) RegisterRoutes() *http.ServeMux { 
	mux := http.NewServeMux()

	mux.HandleFunc("/encurtar", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			return
		}

		h.Create(w, r)
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			return 
		}

		h.Redirect(w, r)
	})

	return mux
}