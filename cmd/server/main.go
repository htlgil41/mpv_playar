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

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	config_vyper, err_configvyper := libs.LoadConfigWithVyper(".")
	if err_configvyper != nil {
		fmt.Println(err_configvyper.Error())
		return
	}

	db_local, err_dblocacl := database.CreateDbSqlLite(config_vyper.Server.DBLite)
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

	serverUnix, err_initserverunix := helpers.StartServerUnix(
		libs.ScreenShotsView(), config_vyper.Paths.Path_servermpv,
	)
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

	server_unix, err_serverunix := libs.ServerSocketForUnix(config_vyper.Paths.Path_servermpv)
	if err_serverunix != nil {
		fmt.Printf("Error al conextarse a unix/pipe %s\n", err_serverunix.Error())
		return
	}

	time.Sleep(3 * time.Second)
	go helpers.ReaderServerUnix(server_unix, db_local, config_vyper)
	defer server_unix.Connect.Close()

	/* CONFIG API GIN EXECUTE */
	router := gin.Default()
	router.Use(libs.RateLimiter())
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/ping", context.PingContext(db_local))
	router.GET("/pid", context.GetLastPids(db_local))
	router.GET("/playlist", context.GETPLAYLISTCONTEXT(db_local))
	router.GET("/videos-mega", context.GETVIDEOPATHCONTEX(config_vyper.Paths.Path_mega))
	router.GET("/metrica", context.GetMetricasVideo(db_local))
	router.POST("/play-vtoplaylist", context.ADDVIDEOPLAYCONTECXT(server_unix, config_vyper))
	router.POST("/next-video", context.NextVideosContext(server_unix))
	router.POST("/playlist", context.CREATEPLAYLIST(db_local))
	router.POST("/add-vtoplaylist", context.ADDVIDEOPLAYLIST(db_local))
	router.POST("/playlist-newplay", context.PlayListNew(db_local, server_unix, config_vyper))
	router.DELETE("/playlist", context.DELETEPLAYLIST(db_local))
	router.DELETE("/stop-playlist", context.STOPPLAYLIST(server_unix))

	router.Run(fmt.Sprintf(":%d", config_vyper.Server.Port))
}
