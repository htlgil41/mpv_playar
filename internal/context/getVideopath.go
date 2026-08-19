package context

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GETVIDEOPATHCONTEX(
	dir string,
) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		dirs, err_dirs := os.ReadDir(dir)
		if err_dirs != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err_dirs.Error(),
				"path":  dir,
			})
			return
		}

		videos := []string{}
		for _, dir := range dirs {
			if dir.IsDir() {
				continue
			}
			videos = append(videos, dir.Name())
		}

		c.JSON(http.StatusOK, gin.H{
			"data":    videos,
			"message": fmt.Sprintf("resultados de archivos leidos en el directorio %s LAS CARPETAS ESTAN OMITIDAS", dir),
		})
	})
}
