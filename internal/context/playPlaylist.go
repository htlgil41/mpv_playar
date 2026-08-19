package context

import (
	"database/sql"
	"encoding/json" // Importante para empaquetar JSON de forma segura
	"fmt"
	"net/http"
	"playar/internal/helpers"
	"playar/internal/libs"
	"playar/internal/repositories"

	"github.com/gin-gonic/gin"
)

func PlayListNew(db *sql.DB, cnet *libs.ConnectionUnix, config *libs.ConfigApp) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		var totalCommands []byte = []byte{}

		clearCmd := `{"command": ["stop"]}` + "\n" + `{"command": ["playlist-clear"]}` + "\n"
		totalCommands = append(totalCommands, []byte(clearCmd)...)

		videos, err_videos := repositories.GetListVideoByPlayList(db, 1)
		if err_videos != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err_videos.Error(), "message": "No se ha podido correr la playlist"})
			return
		}

		if len(videos) == 0 {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "No hay un error en ejecucion sino mas bien la playlist esta vacia",
				"message": "Agregue correctamente la lista de los videos a esta playlist por favor",
			})
			return
		}
		var videos_noadd []string
		videos_agregados := 0

		for _, v := range videos {
			path := fmt.Sprintf("%s%s", config.Paths.Path_mega, v)

			if !helpers.ExisteArchivo(path) {
				videos_noadd = append(videos_noadd, v)
				continue
			}

			cmdStruct := struct {
				Command []string `json:"command"`
			}{
				Command: []string{"loadfile", path, "append"},
			}

			cmdBytes, err_json := json.Marshal(cmdStruct)
			if err_json != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err_json.Error()})
				return
			}

			totalCommands = append(totalCommands, cmdBytes...)
			totalCommands = append(totalCommands, '\n')
			videos_agregados++
		}

		if videos_agregados > 0 {
			playCmd := `{"command": ["playlist-next"]}` + "\n" + `{"command": ["set_property", "pause", false]}` + "\n"
			totalCommands = append(totalCommands, []byte(playCmd)...)
		}
		_, err_write_all := cnet.Connect.Write(totalCommands)
		if err_write_all != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   err_write_all.Error(),
				"message": "Error crítico al comunicar los comandos agrupados a mpv",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data":           videos,
			"message":        "Playlist unificada enviada en un solo bloque con éxito",
			"video_notfound": videos_noadd,
		})
	})
}
