package context

import (
	"database/sql"
	"net/http"
	"playar/internal/repositories"
	"playar/internal/types"

	"github.com/gin-gonic/gin"
)

func CREATEPLAYLIST(db *sql.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		var body types.BODY_CREATE_PLAYLIST
		if err_serealized := c.ShouldBindJSON(&body); err_serealized != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":        "No se ha podido materializar los parametros en el body",
				"detail_error": err_serealized.Error(),
			})
			return
		}

		_, err_result := repositories.InsetPlaylist(db, types.INSERT_PLAYLIST{Nombre: body.Nombre, Descripcion: body.Descripcion})
		if err_result != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Error al tratar de crear una playlist",
				"details": err_result.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"error":   nil,
			"message": "Plyalist creada correctamente",
		})
	})
}
