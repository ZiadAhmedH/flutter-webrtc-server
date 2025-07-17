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

    signaler := signaler.NewSignaler(turnServer)
    wsServer := websocket.NewWebSocketServer(signaler.HandleNewWebSocket, signaler.HandleTurnServerCredentials)

    cfgGen := cfg.Section("general")
    sslCert := cfgGen.Key("cert").String()
    sslKey := cfgGen.Key("key").String()
    bindAddress := cfgGen.Key("bind").MustString("0.0.0.0")
    port, err := cfgGen.Key("port").Int()
    if err != nil {
        port = 8080
    }
    htmlRoot := cfgGen.Key("html_root").MustString("./static")

    serverConfig := websocket.DefaultConfig()
    serverConfig.Host = bindAddress
    serverConfig.Port = port
    serverConfig.CertFile = sslCert
    serverConfig.KeyFile = sslKey
    serverConfig.HTMLRoot = htmlRoot

    // Only one Bind method: it runs TLS if cert/key provided, else HTTP
    wsServer.Bind(serverConfig)
}
