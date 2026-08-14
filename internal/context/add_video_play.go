package context

import (
	"fmt"
	"net"
	"net/http"
	"playar/internal/helpers"
	"playar/internal/libs"
	"playar/internal/types"

	"github.com/gin-gonic/gin"
)

func ADDVIDEOPLAYCONTECXT(
	cnect net.Conn,
	config *libs.ConfigApp,
) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		var body types.BODY_ADD_VIDEO_PLAYLIST
		if err_body := c.ShouldBindJSON(&body); err_body != nil {
			c.JSON(http.StatusOK, gin.H{"message": "matched error body params"})
			return
		}

		command := fmt.Sprintf(
			`{ "command": ["loadfile", "%s%s", "append-play", "fullscreen=yes"] }`,
			config.Paths.Path_mega,
			body.Titulo,
		)
		path := fmt.Sprintf("%s%s", config.Paths.Path_mega, body.Titulo)

		if !helpers.ExisteArchivo(path) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "No se ha podido encontrar el archivo dento del directorio configurado",
				"path":  path,
			})
			return
		}

		_, err_write := cnect.Write([]byte(command + "\n"))
		if err_write != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   err_write.Error(),
				"message": "No se ha podido agregar el video",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Se ha podido ejecutar correctamente el comando.",
			"error":   nil,
		})
	})
}
