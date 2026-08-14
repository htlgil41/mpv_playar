package context

import (
	"bufio"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NextVideosContext(cnet net.Conn) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		_, err_outcommand := cnet.Write([]byte(`{ "command": ["playlist-next"] }` + "\n"))
		if err_outcommand != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   err_outcommand.Error(),
				"message": "No se ha podido pasar el video",
			})
			return
		}

		reader := bufio.NewReader(cnet)
		output, errout := reader.ReadString('\n')
		if errout != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": gin.H{
					"error_message": "No se ha podido leer la salida del servidor unix",
					"error":         errout.Error(),
				},
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Se ha podido ejecutar correctamente el comando verifique el output directamente",
			"output":  output,
		})
	})
}
