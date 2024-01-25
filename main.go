package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

type Config struct {
	listenAddr string
}

type Server struct {
	*Config
}

func NewServer(cfg *Config) *Server {
	return &Server{
		Config: cfg,
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if path == "/test" {
		slog.Info("Test link accessed")
		_, err := w.Write([]byte(fmt.Sprintf("You got this response from %s", s.Config.listenAddr)))
		if err != nil {
			slog.Error(err.Error())
		}
	}
}

func main() {
	config := Config{
		listenAddr: ":3000",
	}
	s := NewServer(&config)
	err := http.ListenAndServe(s.Config.listenAddr, s)
	if err != nil {
		os.Exit(1)
	}

}
