package main

import (
	"context"
	"errors"
	"flag"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/ndabAP/ping-pong/engine"

	"golang.org/x/net/websocket"
)

var (
	serverLogger = log.New(os.Stdout, "[SERVER] ", 0)

	debug = flag.Bool("debug", false, "")
)

func main() {
	flag.Parse()

	http.HandleFunc("/", serveHome)
	http.Handle("/ws", websocket.Handler(serveWs))

	serverLogger.Println("try to open http://127.0.0.1:8080")

	serverLogger.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}

	serverLogger.Println("serving home.html")
	http.ServeFile(w, r, "home.html")
}

func serveWs(ws *websocket.Conn) {
	defer ws.Close()

	game := engine.NewGame(
		800,
		600,
		engine.NewPlayer(10, 150),
		engine.NewPlayer(10, 150),
		engine.NewBall(5, 5),
	)
	engine := engine.NewCanvasEngine(game)
	engine.SetDebug(*debug)
	engine.SetFPS(50)

	// User interface
	frames := make(chan []byte, 1)
	go func(frames chan []byte) {
		for frame := range frames {
			ws.Write(frame)
		}
	}(frames)

	// User input
	input := make(chan []byte, 1)
	defer close(input)
	go func() {
		for {
			buf := make([]byte, 2<<6)
			_, err := ws.Read(buf)
			if err != nil && !errors.Is(err, io.EOF) {
				serverLogger.Fatal(err.Error())
			}
			input <- buf
		}
	}()

	ctx := context.Background()
	engine.NewRound(ctx, frames, input)
}
