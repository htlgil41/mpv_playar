package helpers

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"fmt"
	"path/filepath"
	"playar/internal/libs"
	"playar/internal/repositories"
	"playar/internal/types"
)

func ReaderServerUnix(cnet *libs.ConnectionUnix, db *sql.DB, config *libs.ConfigApp) {
	reader := bufio.NewReader(cnet.Connect)
	defer fmt.Println("LECTOR UNIX MUERTO")
	defer ExirProgram(
		config,
		"Playar se ha detenido porque ha ocurrido una interrupcion en el socket unix/pipe de Mpv",
		config.App.Sucursal,
		"Player ha dejado de leer el Buffer unix/pipe del servicio mpv por lo que se dejo de reproducir videos",
	)

	for {
		vals, errRead := reader.ReadString('\n')
		if errRead != nil {
			fmt.Printf("Conexión cerrada o error de lectura: %v\n", errRead)
			break
		}

		var mapa_eventData map[string]any
		if err := json.Unmarshal([]byte(vals), &mapa_eventData); err != nil {
			continue
		}

		if event, ok := mapa_eventData["event"].(string); ok && event == "start-file" {
			_, err := cnet.Connect.Write([]byte(`{ "command": ["get_property", "playlist"] }` + "\n"))
			if err != nil {
				continue
			}
			continue
		}

		var response types.Response
		if err := json.Unmarshal([]byte(vals), &response); err != nil {
			continue
		}

		if response.Error == "" && len(response.Data) == 0 {
			continue
		}

		for _, item := range response.Data {
			if item.Playing == true {
				var base_url string = filepath.Base(item.Filename)
				_, err_ := repositories.InsertMetricasVideos(
					db,
					types.INSERTMETRICASVIDEOS{
						Video: base_url,
					},
				)

				if err_ != nil {
					fmt.Println(err_.Error())
				}
			}
		}
	}
}
