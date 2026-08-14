package context

import (
	"database/sql"
	"fmt"
	"net"
	"net/http"
	"playar/internal/helpers"
	"playar/internal/libs"
	"playar/internal/repositories"
	"strings"

	"github.com/gin-gonic/gin"
)

func PlayListNew(db *sql.DB, cnet net.Conn, config *libs.ConfigApp) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		_, err_outcommand := cnet.Write([]byte(`{ "command": ["stop"] }` + "\n"))
		if err_outcommand != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   err_outcommand.Error(),
				"message": "No se ha podido limpiar la playlist",
			})
			return
		}

		videos, err_videos := repositories.GetListVideoByPlayList(
			db,
			1,
		)
		if err_videos != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   err_videos.Error(),
				"message": "No se ha podido correr la playlist",
			})
			return
		}

		var command strings.Builder
		command.Grow(len(videos))
		var videos_noadd []string = []string{}

		for _, v := range videos {
			fmt.Fprintf(&command, `{ "command": ["loadfile", "%s%s", "append-play", "fullscreen=yes"] }`+"\n",
				config.Paths.Path_mega,
				v,
			)

			path := fmt.Sprintf("%s%s", config.Paths.Path_mega, v)
			if !helpers.ExisteArchivo(path) {
				videos_noadd = append(videos_noadd, v)
			}
		}

		_, err_write := cnet.Write([]byte(command.String()))
		if err_write != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   err_write.Error(),
				"message": "No se ha podido agregar el video",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data":           videos,
			"message":        "Videos enviada a lista correctamente",
			"video_notfound": videos_noadd,
		})
	})
}
