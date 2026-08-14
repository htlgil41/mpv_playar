package middleware

import (
	"bufio"
	"encoding/json"
	"net"
	"net/http"
	"playar/internal/types"

	"github.com/gin-gonic/gin"
)

func ValidateUnixMiddlewareContext(cnet net.Conn) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		_, err_unix := cnet.Write([]byte(`{ "command": ["get_property", "mpv-version"] }` + "\n"))
		if err_unix != nil {
			c.Next()
			return
		}

		reader := bufio.NewReader(cnet)
		output, _ := reader.ReadString('\n')
		var status types.ServerUnix_StatusResponse

		if err := json.Unmarshal([]byte(output), &status); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Error al deserealizar la respuesta del servidor unix",
				"message": "Debido a este erro no podemos dejarlo pasar al siguiente context",
			})
			return
		}

		if status.Error == "success" {
			c.Next()
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error en la verificacion valide el output",
			"message": "Debido a este error no podemos dejarlo pasar al siguiente context",
			"output":  status,
		})
	})
}
