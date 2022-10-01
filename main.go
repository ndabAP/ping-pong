package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

var serverLogger = log.New(os.Stdout, "[SERVER] ", 0)

var debug = flag.Bool("debug", false, "")

func main() {
	flag.Parse()

	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", serveWs)

	serverLogger.Println("now open your browser at http://localhost:8080")

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}

var upgrader = websocket.Upgrader{}

func serveWs(w http.ResponseWriter, r *http.Request) {
	serverLogger.Println("found websocket handler")

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		panic(err)
	}
	defer ws.Close()

	done := make(chan struct{}, 1)

	g := NewGame(1000, 600, 10, 150, 10, 150, 8, 8)
	jsonFrames := make(chan []byte, 10) // 10 frames buffer
	go func() {
		for f := range jsonFrames {
			ws.WriteMessage(websocket.TextMessage, f)
		}
	}()
	g.Start(r.Context(), jsonFrames)

	<-done
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home.html")
}
