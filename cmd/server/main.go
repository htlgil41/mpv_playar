package main

import (
	"fmt"
	"playar/internal/context"
	"playar/internal/database"
	"playar/internal/helpers"
	"playar/internal/libs"
	"playar/internal/repositories"
	"playar/internal/types"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
)

func main() {
	db_local, err_dblocacl := database.CreateDbSqlLite()
	if err_dblocacl != nil {
		fmt.Println(err_dblocacl)
	}
	defer db_local.Close()

	init_db, err_initdb := helpers.InitDbCreateRegisters(db_local)
	if err_initdb != nil {
		fmt.Println(err_initdb)
	}
	if init_db {
		fmt.Println("EXECUTE INIT DB NO SE CREO NADA DE NUEVO SE PUDO CARGAR LA DB")
	}

	serverUnix, err_initserverunix := helpers.StartServerUnix()
	if err_initserverunix != nil {
		return
	}
	result_execute, _ := repositories.InsertProcess(
		db_local,
		types.INSERTPID{
			Pid:    serverUnix.Pid,
			Path:   serverUnix.Path,
			Estado: "RUN - WILL STOP WITH THIS SERVER GIN",
		},
	)
	if result_execute {
		fmt.Println("PROCESS REGISTER SUCESS")
	}

	server_unix, err_serverunix := libs.ServerSocketForUnix("/tmp/mpvsocket")
	if err_serverunix != nil {
		fmt.Println(err_serverunix.Error())
		return
	}

	defer server_unix.Close()

	time.Sleep(8 * time.Second)
	/* CONFIG API GIN EXECUTE */
	router := gin.Default()

	/* GETS */
	router.GET("/ping", context.PingContext(db_local, server_unix))
	router.GET("/pid", context.GetLastPids(db_local))
	router.GET("/videos", context.GETVIDEOSPAGES(db_local))

	/* INSERT */
	router.GET("/video", context.CreateVideoContext(db_local, types.VIDEOS{Titulo: "Prueba", Descripcion: "Prueba de video", Nombre_archivo: "Prueba.mp4"}))

	router.Run(":8000")
}
