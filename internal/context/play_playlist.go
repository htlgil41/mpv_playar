package context

import (
	"database/sql"
	"fmt"
	"net"
	"net/http"
	"playar/internal/helpers"
	"playar/internal/libs"
	"playar/internal/repositories"

	"github.com/gin-gonic/gin"
)

func PlayListNew(db *sql.DB, cnet net.Conn, config *libs.ConfigApp) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		_, err_outcommand := cnet.Write([]byte(`{ "command": ["stop"] }` + "\n" + `{ "command": ["playlist-clear"] }` + "\n"))
		if err_outcommand != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   err_outcommand.Error(),
				"message": "No se ha podido limpiar la playlist",
			})
			return
		}

		videos, err_videos := repositories.GetListVideoByPlayList(db, 1)
		if err_videos != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   err_videos.Error(),
				"message": "No se ha podido correr la playlist",
			})
			return
		}

		var videos_noadd []string = []string{}
		videos_agregados := 0

		for _, v := range videos {
			path := fmt.Sprintf("%s%s", config.Paths.Path_mega, v)

			if !helpers.ExisteArchivo(path) {
				videos_noadd = append(videos_noadd, v)
				continue
			}

			var command string
			if videos_agregados == 0 {
				command = fmt.Sprintf(`{ "command": ["loadfile", "%s", "append-play"] }`, path)
			} else {
				command = fmt.Sprintf(`{ "command": ["loadfile", "%s", "append"] }`, path)
			}

			_, err_write := cnet.Write([]byte(command + "\n"))
			if err_write != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error":   err_write.Error(),
					"message": "No se ha podido agregar el video a la lista",
				})
				return
			}

			videos_agregados++
		}

		c.JSON(http.StatusOK, gin.H{
			"data":           videos,
			"message":        "Playlist unificada creada y reproduciendo correctamente",
			"video_notfound": videos_noadd,
		})
	})
}
