package websocket

import (
	"fmt"

	"github.com/kataras/iris"
	"github.com/kataras/iris/websocket"
)

func SetupWebsocket(app *iris.Application) {
	// create our echo websocket server
	ws := websocket.New(websocket.Config{
		// These are low-level optionally fields,
		// user/client can't see those values.
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		// only javascript client-side code has the same rule,
		// which you serve using the ws.ClientSource (see below).
		EvtMessagePrefix: []byte("my-custom-prefix:"),
	})
	ws.OnConnection(handleConnection)
	// register the server on an endpoint.
	// see the inline javascript code in the websockets.html, this endpoint is used to connect to the server.
	app.Get("/websocket", ws.Handler())
}

func handleConnection(c websocket.Connection) {
	// Read events from browser
	c.OnMessage(func(msg []byte) {
		fmt.Printf("%s sent: %s\n", c.Context().RemoteAddr(), msg)
		c.To(websocket.Broadcast).Emit("chat", msg)
	})
}
