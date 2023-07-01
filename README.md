# ping-pong

`ping-pong` tries to be a generic pong implementation, written in Go. It
includes a Canvas engine playable in a browser wit Canvas support.

![Screenshot from 2023-06-06 07-46-48](https://github.com/ndabAP/ping-pong/assets/8510570/86c9569e-9892-4401-a96a-ce63adb6af82)

## How to start

To play in the browser with the Canvas engine, download or pull the repository
to start the backend server:

```bash
$ go run *.go -debug
```

Then open the frontend at [http://127.0.0.1:8080](http://127.0.0.1:8080) in your
browser. Alternatively, you can use `make`:

```bash
$ make play_canvas
```

## How to play

The left player is player one, the right player two. There are two supported
inputs available: <kbd>↑</kbd> moves player one up and <kbd>↓</kbd> moves him
down. Players get points for goals.

## API

You can create your own frontend with the provided API, which consists of
constraints and helpers. After you `go get` the module, import the `engine`
package:

```go
import github.com/ndabAP/ping-pong/engine
```

To get the module:

```bash
go get github.com/ndabAP/ping-pong
```

## Future goals

- Native sound generation
- AI-controlled player two
- More retro style (e. g. retro GUI)
