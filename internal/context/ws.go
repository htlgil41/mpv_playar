package context

import (
	"log"
	"playar/internal/libs"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func Ws(upgrader websocket.Upgrader, hub *libs.Hub) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Printf("WebSocket upgrade error: %v", err)
			return
		}
		hub.AddClient(conn)
		defer hub.RemoveClient(conn)

		hub.Broadcast(1, []byte("Alguien mas se ha conectado"))
		for {
			conn.ReadMessage()
		}
	})
}
