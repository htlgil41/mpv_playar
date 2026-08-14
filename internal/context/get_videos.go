package context

import (
	"database/sql"
	"net/http"
	"playar/internal/repositories"

	"github.com/gin-gonic/gin"
)

func GETVIDEOSPAGES(
	db *sql.DB,
) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		videos, err_videos := repositories.GETVIDEOSPAGES(
			db,
			0,
		)
		if err_videos != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "No se ha podido recuperar los videos",
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "Videos pagina #1",
			"data":    &videos,
		})
	})
}
