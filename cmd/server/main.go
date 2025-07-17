package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // يسمح بأي origin
	},
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("فشل الترقية إلى WebSocket:", err)
		return
	}
	defer conn.Close()

	log.Println("✅ اتصال جديد عبر WebSocket")

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("❌ خطأ في قراءة الرسالة:", err)
			break
		}

		log.Printf("📨 استلمنا: %s", message)

		err = conn.WriteMessage(messageType, message)
		if err != nil {
			log.Println("❌ خطأ في إرسال الرسالة:", err)
			break
		}
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback افتراضي
	}

	http.HandleFunc("/ws", handleWebSocket)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "🌐 WebSocket server is running on /ws\n")
	})

	log.Printf("🚀 السيرفر شغال على البورت: %s", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("❌ فشل في تشغيل السيرفر: %v", err)
	}
}
