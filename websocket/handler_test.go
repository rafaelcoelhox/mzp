package websocket

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/websocket"
)

func TestWebSocketHandler(t *testing.T) {
	handler := NewWebSocketHandler()
	server := httptest.NewServer(http.HandlerFunc(handler.HandleConnection))
	defer server.Close()

	t.Run("Teste de conexão bem-sucedida", func(t *testing.T) {
		wsURL := "ws" + server.URL[4:]
		conn, resp, err := websocket.DefaultDialer.Dial(wsURL, nil)
		if err != nil {
			t.Fatalf("Erro ao conectar: %v", err)
		}
		defer conn.Close()

		if resp.StatusCode != http.StatusSwitchingProtocols {
			t.Errorf("Status code incorreto. Esperado %d, Recebido %d",
				http.StatusSwitchingProtocols, resp.StatusCode)
		}
	})

	t.Run("Teste de requisição não-WebSocket", func(t *testing.T) {
		resp, err := http.Get(server.URL)
		if err != nil {
			t.Fatalf("Erro na requisição HTTP: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusBadRequest {
			t.Errorf("Esperado BadRequest, recebido: %d", resp.StatusCode)
		}
	})
}
