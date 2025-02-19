package main

import (
	"log"
	"net/http"

	"github.com/rafaelcoelhox/mzp/websocket"
)

func main() {
	wsHandler := websocket.NewWebSocketHandler()

	http.HandleFunc("/", wsHandler.HandleConnection)

	log.Println("Servidor iniciado na porta :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Erro ao iniciar servidor: %v", err)
	}
}
