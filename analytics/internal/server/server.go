package server

import (
	"analytics/internal/config"
	"analytics/internal/storage"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

type Server struct {
	cfg     *config.Server
	storage storage.Storage
}

func (s *Server) CreateServer() func() {
	router := chi.NewRouter()

	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ping"))
	})

	router.Get("/metrics", s.MetricsHandler)

	corsHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		router.ServeHTTP(w, r)
	})

	srv := &http.Server{
		Addr:         s.cfg.Addr,
		Handler:      corsHandler,
		ReadTimeout:  s.cfg.Timeout,
		WriteTimeout: s.cfg.Timeout,
		IdleTimeout:  s.cfg.IdleTimeout,
	}

	go func() {
		fmt.Printf("HTTP server starting: %s\n", s.cfg.Addr)

		if err := srv.ListenAndServe(); err != nil {
			fmt.Printf("error: %s\n", err)
			return
		}
	}()

	return func() {
		_ = srv.Close()
	}
}

func NewRestAPI(cfg *config.Server, storage storage.Storage) *Server {
	return &Server{
		cfg:     cfg,
		storage: storage,
	}
}
