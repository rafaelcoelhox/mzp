package main

import (
	"log"
	"net/http"

	"github.com/rafaelcoelhox/mzp/config"
	"github.com/rafaelcoelhox/mzp/websocket"
)

func main() {
	cfg, err := config.Load("config/config.toml")
	if err != nil {
		log.Fatalf("Erro ao carregar configuração: %v", err)
	}

	wsHandler := websocket.NewWebSocketHandler()

	http.HandleFunc("/", wsHandler.HandleConnection)

	log.Println("Servidor iniciado na porta :8080")
	if err := http.ListenAndServe(cfg.Server.Port, nil); err != nil {
		log.Fatalf("Erro ao iniciar servidor: %v", err)
	}
}
