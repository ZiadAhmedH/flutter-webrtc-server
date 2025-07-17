package main

import (
	"log"
	"net/http"
	"os"

	"github.com/flutter-webrtc/flutter-webrtc-server/pkg/logger"
	"github.com/flutter-webrtc/flutter-webrtc-server/pkg/websocket"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Create a new SFU instance
	s := sfu.NewSFU(sfu.Config{})

	// Create WebSocket server
	wsServer := websocket.NewWebSocketServer(s)

	// Define HTTP mux
	mux := http.NewServeMux()

	// Register WebSocket handler
	mux.HandleFunc("/ws", wsServer.ServeWebSocket)

	// Optionally add pprof or health check endpoints
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	logger.Infof("WebSocket server listening on :%s", port)
	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		logger.Errorf("server error: %v", err)
	}
}
