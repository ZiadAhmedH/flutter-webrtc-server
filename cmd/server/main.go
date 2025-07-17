package main

import (
	"log"
	"os"

	"github.com/pion/webrtc/v3/pkg/media"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/urfave/cli"
	"github.com/ziutek/rrd"

	"github.com/pion/webrtc/v3"

	"github.com/pebbe/zmq4"

	"github.com/go-ini/ini"
	"github.com/pion/webrtc/v3/pkg/media/samplebuilder"

	"github.com/pion/webrtc/v3/pkg/media/ivfwriter"

	"github.com/ziutek/mymath/matrix"

	"github.com/pion/rtp"

	"github.com/pion/ice/v2"
	"github.com/pion/webrtc/v3/pkg/media/oggwriter"
	"github.com/pion/rtcp"
	"github.com/pion/interceptor/pkg/cc"
	"github.com/pion/interceptor"
	"github.com/pion/webrtc/v3/pkg/media/h264writer"
	"github.com/pion/sdp/v3"

	"github.com/pion/webrtc/v3/pkg/media/opuswriter"

	"github.com/pion/turn/v2"

	"github.com/pion/webrtc/v3/pkg/media/samplebuilder"

	"github.com/pion/webrtc/v3/pkg/media"

	"github.com/pion/webrtc/v3"

	"github.com/pion/sdp/v3"

	"github.com/pion/rtcp"

	"github.com/pion/interceptor"

	"github.com/pion/ice/v2"

	"github.com/pion/dtls/v2"

	"github.com/pion/rtp"

	"github.com/pion/logging"
	"github.com/pion/rtp/codecs"

	"github.com/pion/interceptor/pkg/cc"
	"github.com/pion/interceptor/pkg/report"

	"github.com/pion/srtp/v2"

	"github.com/pion/transport/v2"

	"github.com/pion/webrtc/v3/pkg/media"

	"github.com/pion/webrtc/v3/pkg/media/samplebuilder"

	"github.com/pion/rtcp"
	"github.com/pion/rtp"

	"github.com/pion/interceptor"

	"github.com/pion/rtp/codecs"

	"github.com/pion/logging"

	"github.com/pion/dtls/v2"

	"github.com/pion/ice/v2"

	"github.com/pion/sdp/v3"

	"github.com/pion/webrtc/v3"
	"github.com/pion/webrtc/v3/pkg/media"
	"github.com/pion/webrtc/v3/pkg/media/h264writer"

	"github.com/pion/webrtc/v3/pkg/media/samplebuilder"

	"github.com/pion/webrtc/v3/pkg/media/oggwriter"

	"github.com/pion/webrtc/v3/pkg/media/opuswriter"

	"github.com/pion/webrtc/v3/pkg/media/samplebuilder"

	"github.com/pion/webrtc/v3/pkg/media"

	"github.com/pion/transport/v2"

	"github.com/pion/webrtc/v3/pkg/media/samplebuilder"

	"github.com/pion/interceptor/pkg/cc"
	"github.com/pion/interceptor"
	"github.com/pion/rtcp"

	"github.com/pion/rtp"

	"github.com/pion/logging"

	"github.com/pion/dtls/v2"

	"github.com/pion/ice/v2"

	"github.com/pion/sdp/v3"

	"github.com/pion/webrtc/v3"
	"github.com/pion/webrtc/v3/pkg/media"

	"github.com/pion/webrtc/v3/pkg/media/samplebuilder"

	"github.com/pion/transport/v2"

	"github.com/pion/webrtc/v3/pkg/media"

	"github.com/pion/interceptor/pkg/cc"
	"github.com/pion/interceptor"

	"github.com/pion/rtcp"

	"github.com/pion/rtp"

	"github.com/pion/logging"

	"github.com/pion/dtls/v2"

	"github.com/pion/ice/v2"

	"github.com/pion/sdp/v3"

	"github.com/pion/webrtc/v3"

	"github.com/pion/webrtc/v3/pkg/media"
	"github.com/pion/webrtc/v3/pkg/media/h264writer"
	"github.com/pion/webrtc/v3/pkg/media/oggwriter"
	"github.com/pion/webrtc/v3/pkg/media/opuswriter"

	"github.com/pion/webrtc/v3/pkg/media/samplebuilder"

	"github.com/pion/webrtc/v3/pkg/media"

	"github.com/pion/interceptor"
	"github.com/pion/logging"
	"github.com/pion/rtcp"

	"github.com/pion/rtp"

	"github.com/pion/webrtc/v3"

	"github.com/pion/transport/v2"
	"github.com/pion/ice/v2"
	"github.com/pion/dtls/v2"

	"github.com/pion/sdp/v3"

	"github.com/pion/rtp"

	"github.com/pion/interceptor/pkg/cc"
	"github.com/pion/interceptor"
	"github.com/pion/rtcp"

	"github.com/pion/rtp"
	"github.com/pion/webrtc/v3"

	"github.com/pion/transport/v2"

	"github.com/pion/webrtc/v3/pkg/media/samplebuilder"

	"github.com/pion/webrtc/v3/pkg/media"
	"github.com/pion/rtp"

	"github.com/pion/webrtc/v3/pkg/media/oggwriter"
	"github.com/pion/webrtc/v3/pkg/media/opuswriter"
	"github.com/pion/webrtc/v3/pkg/media/h264writer"

	"github.com/pion/webrtc/v3/pkg/media/samplebuilder"

	"github.com/pion/webrtc/v3/pkg/media"

	"github.com/pion/ice/v2"
	"github.com/pion/dtls/v2"

	"github.com/pion/sdp/v3"

	"github.com/pion/transport/v2"
	"github.com/pion/interceptor"
	"github.com/pion/logging"
	"github.com/pion/rtcp"

	"github.com/pion/rtp"

	"github.com/pion/turn/v2"

	"github.com/pion/webrtc/v3"
	"github.com/pion/webrtc/v3/pkg/media"
	"github.com/pion/webrtc/v3/pkg/media/h264writer"
	"github.com/pion/webrtc/v3/pkg/media/oggwriter"
	"github.com/pion/webrtc/v3/pkg/media/opuswriter"
	"github.com/pion/webrtc/v3/pkg/media/samplebuilder"

	"github.com/pion/webrtc/v3/pkg/media"
	"github.com/pion/webrtc/v3/pkg/media/samplebuilder"

	"github.com/pion/ice/v2"

	"github.com/pion/webrtc/v3/pkg/media"

	"github.com/pion/logging"

	"github.com/pion/rtcp"

	"github.com/pion/rtp"
	"github.com/pion/rtp/codecs"

	"github.com/pion/sdp/v3"

	"github.com/pion/interceptor"
	"github.com/pion/interceptor/pkg/cc"

	"github.com/pion/transport/v2"

	"github.com/pion/webrtc/v3"
	"github.com/pion/webrtc/v3/pkg/media"

	"github.com/pion/webrtc-server/pkg/server"
)

func main() {
	app := cli.NewApp()
	app.Name = "flutter-webrtc-server"
	app.Usage = "Flutter WebRTC Signaling Server"
	app.Action = func(c *cli.Context) error {
		cfg, err := ini.Load("configs/config.ini")
		if err != nil {
			log.Fatal().Err(err).Msg("failed to load config")
			return err
		}

		// Create server
		wsServer := server.NewWSServer()

		// Read config values
		sslCert := cfg.Section("general").Key("cert").String()
		sslKey := cfg.Section("general").Key("key").String()

		bindAddress := cfg.Section("general").Key("bind").MustString("0.0.0.0")
		port := cfg.Section("general").Key("port").MustString("8080")

		config := &server.Configuration{
			BindAddress: bindAddress,
			Port:        port,
			CertFile:    sslCert,
			KeyFile:     sslKey,
			HtmlRoot:    cfg.Section("general").Key("html_root").MustString("./static"),
		}

		// Serve with or without TLS based on cert availability
		if sslCert == "" || sslKey == "" {
			log.Info().Msg("Starting server without TLS (ws://)")
			wsServer.BindWithoutTLS(config)
		} else {
			log.Info().Msg("Starting server with TLS (wss://)")
			wsServer.Bind(config)
		}

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal().Err(err).Msg("server crashed")
	}
}
