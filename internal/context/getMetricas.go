package context

import (
	"database/sql"
	"net/http"
	"playar/internal/repositories"

	"github.com/gin-gonic/gin"
)

func GetMetricasVideo(db *sql.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		resultado, err_resultado := repositories.GetMetricaFuncToday(db)
		if err_resultado != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   err_resultado.Error(),
				"message": "No se ha podido obtener la informacion de la db",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"error":   nil,
			"message": "Las metricas obtenidas son las transcurridas el dia de HOY",
			"data":    resultado,
		})
	})
}
