package context

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PingContext(db_local *sql.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		err_ping := db_local.Ping()

		if err_ping != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"sucess":    "not",
				"status":    "server error ping db local",
				"status_db": err_ping.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"sucess":    "ok",
			"status":    "server ok",
			"status_db": "ok",
		})
	})
}
