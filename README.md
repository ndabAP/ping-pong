# ping-pong

`ping-pong` is a generic pong implementation in Go with a Canvas engine.

## How to start

To play with the Canvas engine, download or pull the repository to start the
backend server:

```bash
$ go run *.go -debug
```

Then open the frontend at [http://127.0.0.1:8080](http://127.0.0.1:8080) in your
browser.

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
