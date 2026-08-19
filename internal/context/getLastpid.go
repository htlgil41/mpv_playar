package context

import (
	"database/sql"
	"net/http"
	"playar/internal/repositories"

	"github.com/gin-gonic/gin"
)

func GetLastPids(db *sql.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {

		execute, err := repositories.GETLASTPIDPATH(
			db,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   err,
				"message": "No se han podido recibir los ultimos pid de processos",
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"error":   nil,
			"message": "resultados obtenidos",
			"data":    &execute,
		})
	})
}
