package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"os"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// السماح بأي origin (مهم لو شغال من Flutter أو Web)
		return true
	},
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	log.Println("🔌 اتصال WebSocket جديد")

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("❌ فشل الترقية:", err)
		return
	}
	defer conn.Close()

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("❌ خطأ في قراءة الرسالة:", err)
			break
		}
		log.Printf("📩 استلم: %s\n", message)

		// إرسال نفس الرسالة كـ echo
		if err := conn.WriteMessage(messageType, message); err != nil {
			log.Println("❌ خطأ في الإرسال:", err)
			break
		}
	}
}

func main() {
	http.HandleFunc("/ws", handleWebSocket)

	// استخدام بورت من environment (Railway بتوفر PORT تلقائيًا)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("🚀 السيرفر شغال على البورت:", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
