package context

import (
	"net/http"
	"playar/internal/libs"

	"github.com/gin-gonic/gin"
)

func NextVideosContext(cnet *libs.ConnectionUnix) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		_, err_outcommand := cnet.Connect.Write([]byte(`{ "command": ["playlist-next"] }` + "\n"))
		if err_outcommand != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   err_outcommand.Error(),
				"message": "No se ha podido pasar el video",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Se ha podido ejecutar correctamente el comando",
		})
	})
}
