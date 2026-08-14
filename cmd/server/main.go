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

	config_vyper, err_configvyper := libs.LoadConfigWithVyper(".")
	if err_configvyper != nil {
		fmt.Println(err_configvyper.Error())
		return
	}

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

	serverUnix, err_initserverunix := helpers.StartServerUnix(libs.ScreenShotsView())
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

	server_unix, err_serverunix := libs.ServerSocketForUnix(config_vyper.UnixServerMpv.Path_server)
	if err_serverunix != nil {
		fmt.Println(err_serverunix.Error())
		return
	}

	defer server_unix.Close()
	time.Sleep(3 * time.Second)
	/* CONFIG API GIN EXECUTE */
	router := gin.Default()

	router.GET("/ping", context.PingContext(db_local))
	router.GET("/pid", context.GetLastPids(db_local))
	router.GET("/playlist", context.GETPLAYLISTCONTEXT(db_local))
	router.GET("/videos-mega", context.GETVIDEOPATHCONTEX(config_vyper.Paths.Path_mega))
	router.POST("/add-videoplaylist", context.ADDVIDEOPLAYCONTECXT(server_unix, config_vyper))
	router.POST("/next", context.NextVideosContext(server_unix))
	router.POST("/playlist", context.CREATEPLAYLIST(db_local))
	router.POST("/v-playlist", context.ADDVIDEOPLAYLIST(db_local))
	router.POST("/playlistnewply", context.PlayListNew(db_local, server_unix, config_vyper))
	router.DELETE("/playlist", context.CLEARPLAYLISTCONTEXT(server_unix))

	router.Run(fmt.Sprintf(":%d", config_vyper.Server.Port))
}
