package context

import (
	"database/sql"
	"net/http"
	"playar/internal/repositories"
	"playar/internal/types"

	"github.com/gin-gonic/gin"
)

func DELETEPLAYLIST(db *sql.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		var body types.BODY_DELETE_PLAYLIST
		if err_body := c.ShouldBindJSON(&body); err_body != nil {
			c.JSON(http.StatusOK, gin.H{"message": "matched error body params", "error": err_body.Error()})
			return
		}

		_, errDeletePlaylist := repositories.DELETEPLAYLIST(db, body.Playlist_id)
		if errDeletePlaylist != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   errDeletePlaylist.Error(),
				"message": "Error al tratar de eliminar la playlist valide e intente nuevamente",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   nil,
			"message": "Se ha eliminado la playlist correctamente junto con los videos asociados a ella",
		})
	})
}
