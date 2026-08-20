package context

import (
	"database/sql"
	"net/http"
	"playar/internal/repositories"
	"playar/internal/types"

	"github.com/gin-gonic/gin"
)

func MetricaByRangeDate(db *sql.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		var params types.PARAM_GET_METRICAS_RANGES
		if err_serealized := c.ShouldBindQuery(&params); err_serealized != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Error al tratar de serealizar los datos",
				"message": err_serealized.Error(),
			})
			return
		}

		resultado, err_resultado := repositories.GetMetricasByRanger(db, types.RANGEBEETWENDATEREPOSITORIE{
			Gte: params.Gte,
			Lte: params.Lte,
		})
		if err_resultado != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   err_resultado.Error(),
				"message": "No se ha podido obtener la informacion de la db",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"querys": params,
			"data":   resultado,
		})
	})
}
