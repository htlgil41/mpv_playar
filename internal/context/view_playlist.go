package context

import (
	"bufio"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
)

func VIEWPLAYLISTCONTEXT(cnet net.Conn) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		_, err_outcommand := cnet.Write([]byte(`{ "command": ["get_property", "playlist"] }` + "\n"))
		if err_outcommand != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   err_outcommand.Error(),
				"message": "No se ha podido pasar el video",
			})
			return
		}

		reader := bufio.NewReader(cnet)
		output, err_putput := reader.ReadString('\n')
		if err_putput != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": gin.H{
					"error_message": "No se ha podido leer la salida del servidor unix",
					"error":         err_putput.Error(),
				},
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Playlist actualmente en el servidor unix",
			"output":  output,
		})
	})
}
