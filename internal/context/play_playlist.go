package context

import (
	"database/sql"
	"net"
	"net/http"
	"playar/internal/repositories"

	"github.com/gin-gonic/gin"
)

func PlayListNew(db *sql.DB, cnet net.Conn) gin.HandlerFunc {
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

		c.JSON(http.StatusOK, gin.H{
			"data":    videos,
			"message": "Videos enviada a lista correctamente",
		})
	})
}
