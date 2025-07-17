package main

import (
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
		logger.Errorf("Fail to read config file: %v", err)
		os.Exit(1)
	}

	// Load TURN config
	publicIP := cfg.Section("turn").Key("public_ip").String()
	stunPort, err := cfg.Section("turn").Key("port").Int()
	if err != nil {
		stunPort = 3478
	}
	realm := cfg.Section("turn").Key("realm").String()

	turnConfig := turn.DefaultConfig()
	turnConfig.PublicIP = publicIP
	turnConfig.Port = stunPort
	turnConfig.Realm = realm

	turnServer := turn.NewTurnServer(turnConfig)

	// WebSocket signaling server
	signaler := signaler.NewSignaler(turnServer)
	wsServer := websocket.NewWebSocketServer(signaler.HandleNewWebSocket, signaler.HandleTurnServerCredentials)

	// Load WebSocket server config
	sslCert := cfg.Section("general").Key("cert").String()
	sslKey := cfg.Section("general").Key("key").String()
	bindAddress := cfg.Section("general").Key("bind").MustString("0.0.0.0")
	port, err := cfg.Section("general").Key("port").Int()
	if err != nil {
		port = 8080
	}
	htmlRoot := cfg.Section("general").Key("html_root").MustString("./static")

	serverConfig := websocket.DefaultConfig()
	serverConfig.Host = bindAddress
	serverConfig.Port = port
	serverConfig.CertFile = sslCert
	serverConfig.KeyFile = sslKey
	serverConfig.HTMLRoot = htmlRoot

	// Use Bind or BindWithoutTLS depending on SSL settings
	if sslCert != "" && sslKey != "" {
		wsServer.Bind(serverConfig)
	} else {
		wsServer.BindWithoutTLS(serverConfig)
	}
}
