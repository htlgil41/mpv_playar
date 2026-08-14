package context

import (
	"database/sql"
	"net/http"
	"playar/internal/repositories"
	"playar/internal/types"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateVideoContext(db *sql.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		var body types.BODY_CREATE_NEW_VIDEO
		if err_body := c.ShouldBindJSON(&body); err_body != nil {
			c.JSON(http.StatusOK, gin.H{"message": "matched error body params"})
		}

		_, err_execute := repositories.InsertVideo(
			db,
			types.VIDEOS{
				Titulo:         body.Titulo,
				Descripcion:    body.Descripcion,
				Nombre_archivo: body.Nombre_archivo,
			},
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
			"data":    body,
		})
	})
}
