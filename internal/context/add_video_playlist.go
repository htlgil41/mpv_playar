package context

import (
	"database/sql"
	"net/http"
	"playar/internal/repositories"
	"playar/internal/types"

	"github.com/gin-gonic/gin"
)

func ADDVIDEOPLAYLIST(db *sql.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		var body types.BODY_ADD_MUSIC_PLAYLIST
		if err_serealized := c.ShouldBindJSON(&body); err_serealized != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Error al serealizar",
				"message": "No se pudo serealizar los datos que se requieren guardar",
			})
			return
		}

		_, err_exec := repositories.InserVideoPlaylist(db, types.INSERT_MUSIC_PLAYLIST{
			Playlist_id: body.Playlist_id,
			Video:       body.Video,
			Orden:       body.Orden,
		})
		if err_exec != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "No se ha podido completar la consula",
				"details": err_exec.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   nil,
			"message": "Musica agregada correctamente",
			"data":    body,
		})
	})
}
