package main

import (
	"log"
	"net/http"
	"os"

	"github.com/flutter-webrtc/flutter-webrtc-server/pkg/logger"
	"github.com/flutter-webrtc/flutter-webrtc-server/pkg/signaler"
	"github.com/flutter-webrtc/flutter-webrtc-server/pkg/turn"
	"github.com/flutter-webrtc/flutter-webrtc-server/pkg/websocket"
	"gopkg.in/ini.v1"
)

func main() {
	cfg, err := ini.Load("configs/config.ini")
	if err != nil {
		logger.Errorf("failed to load config: %v", err)
		os.Exit(1)
	}

	// TURN server config (optional to include)
	turnCfg := turn.DefaultConfig()
	if ip := cfg.Section("turn").Key("public_ip").String(); ip != "" {
		turnCfg.PublicIP = ip
	}
	turnCfg.Port = cfg.Section("turn").Key("port").MustInt(3478)
	turnCfg.Realm = cfg.Section("turn").Key("realm").MustString("flutterwebrtc")
	turnServer := turn.NewTurnServer(turnCfg)

	// Signaling
	sig := signaler.NewSignaler(turnServer)
	ws := websocket.NewWebSocketServer(sig.HandleNewWebSocket, sig.HandleTurnServerCredentials)

	// WebSocket Handler
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		log.Println("New WS /ws request from", r.RemoteAddr)
		ws.ServeHTTP(w, r)
	})

	// Optional health check
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("OK"))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = cfg.Section("general").Key("port").MustString("8080")
	}

	logger.Infof("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		logger.Errorf("server failed: %v", err)
	}
}
