package context

import (
	"database/sql"
	"net/http"
	"playar/internal/repositories"

	"github.com/gin-gonic/gin"
)

func GETPLAYLISTCONTEXT(db *sql.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		playlist, err_playlist := repositories.GetListOfPlaylist(db)
		if err_playlist != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Error al obtener el listado de playlist",
				"details": err_playlist.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data":    playlist,
			"message": "Messages obtenidos desde la db",
		})
	})
}
