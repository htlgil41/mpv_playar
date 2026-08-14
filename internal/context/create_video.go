package context

import (
	"database/sql"
	"net/http"
	"playar/internal/repositories"
	"playar/internal/types"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateVideoContext(db *sql.DB, input types.VIDEOS) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		_, err_execute := repositories.InsertVideo(
			db,
			input,
		)

		if err_execute != nil {
			if strings.Contains(err_execute.Error(), "UNIQUE constraint") {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error":   err_execute.Error(),
					"message": "El video no se ha podido crear ya que hay uno que tiene le mismo nombre",
				})
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   err_execute.Error(),
				"message": "No se ha podido registrar un nuevo video",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   nil,
			"message": "Se ha registrado el video correctamente",
			"data":    input,
		})
	})
}
