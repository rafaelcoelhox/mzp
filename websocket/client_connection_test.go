package websocket

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

func TestClientConnection(t *testing.T) {
	// Configurar servidor de teste
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := websocket.Upgrade(w, r, nil, 1024, 1024)
		if err != nil {
			t.Fatalf("Erro ao atualizar conexão: %v", err)
		}
		client := NewClientConnection(conn)
		go client.HandleMessages()
	}))
	defer server.Close()

	// Conectar ao servidor WebSocket
	wsURL := "ws" + server.URL[4:]
	conn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		t.Fatalf("Erro ao conectar: %v", err)
	}
	defer conn.Close()

	t.Run("Teste de mensagem válida", func(t *testing.T) {
		testMessage := []byte("Olá, teste!")
		expectedPrefix := []byte("Sua mensagem: Olá, teste!. Recebida em: ")

		// Enviar mensagem
		if err := conn.WriteMessage(websocket.TextMessage, testMessage); err != nil {
			t.Fatalf("Erro ao escrever mensagem: %v", err)
		}

		// Ler resposta
		_, resp, err := conn.ReadMessage()
		if err != nil {
			t.Fatalf("Erro ao ler resposta: %v", err)
		}

		// Verificar formato da resposta
		if !bytes.HasPrefix(resp, expectedPrefix) {
			t.Errorf("Resposta inesperada:\nEsperado prefixo: %s\nRecebido: %s",
				expectedPrefix, resp)
		}

		// Verificar formato da data
		dateStr := resp[len(expectedPrefix):]
		if _, err := time.Parse(time.RFC3339, string(dateStr)); err != nil {
			t.Errorf("Formato de data inválido: %v", err)
		}
	})

	t.Run("Teste de fechamento de conexão", func(t *testing.T) {
		// Enviar mensagem de fechamento
		err := conn.WriteMessage(
			websocket.CloseMessage,
			websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""),
		)
		if err != nil {
			t.Fatalf("Erro ao enviar fechamento: %v", err)
		}

		// Verificar fechamento
		_, _, err = conn.ReadMessage()
		if err == nil {
			t.Error("Esperado erro de conexão fechada, mas nenhum ocorreu")
		}
	})
}

func TestProcessMessage(t *testing.T) {
	mockConn := &websocket.Conn{} // Conexão mockada
	client := NewClientConnection(mockConn)
	testMessage := []byte("Teste")

	result := client.processMessage(testMessage)
	expected := []byte("Sua mensagem: Teste. Recebida em: ")

	if !bytes.HasPrefix(result, expected) {
		t.Errorf("Saída inesperada:\nEsperado: %s\nRecebido: %s", expected, result)
	}
}
