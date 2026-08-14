package context

import (
	"bufio"
	"database/sql"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PingContext(
	db_local *sql.DB,
	cnet net.Conn,
) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		err_ping := db_local.Ping()
		_, err_unix := cnet.Write([]byte(`{ "command": ["get_property", "mpv-version"] }` + "\n"))

		if err_ping != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"sucess":    "not",
				"status":    "server error ping db local",
				"status_db": err_ping.Error(),
			})
			return
		}

		if err_unix != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"sucess":      "not",
				"status":      "server error unix",
				"status_unix": err_unix.Error(),
			})
			return
		}

		reader := bufio.NewReader(cnet)
		output, errout := reader.ReadString('\n')

		if errout != nil {
			c.JSON(http.StatusOK, gin.H{
				"sucess":      "ok",
				"status":      "server ok",
				"status_db":   "ok",
				"server_unix": "No se ha podido leer la respuesta pero se ejecuto el comando",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"sucess":      "ok",
			"status":      "server ok",
			"status_db":   "ok",
			"server_unix": output,
		})
	})
}
