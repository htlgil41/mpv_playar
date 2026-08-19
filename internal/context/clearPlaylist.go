package context

import (
	"net/http"
	"playar/internal/libs"

	"github.com/gin-gonic/gin"
)

func STOPPLAYLIST(cnet *libs.ConnectionUnix) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		_, err_outcommand := cnet.Connect.Write([]byte(`{ "command": ["stop"] }` + "\n"))
		if err_outcommand != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   err_outcommand.Error(),
				"message": "No se ha podido limpiar la playlist",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "Comando correctamente ejecutado",
		})
	})
}
