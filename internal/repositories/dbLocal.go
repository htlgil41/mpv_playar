package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"playar/internal/types"
	"playar/internal/vars"
)

/* INSERTS */
func InsertProcess(db *sql.DB, input types.INSERTPID) (bool, error) {
	tx, err_tx := db.Begin()
	if err_tx != nil {
		return false, err_tx
	}
	defer tx.Commit()

	state, err_state := tx.Prepare(vars.INSERT_PID_STATEMENT)
	if err_state != nil {
		return false, err_state
	}
	defer state.Close()

	result, err_result := state.Exec(
		input.Pid,
		input.Path,
		input.Estado,
	)
	if err_result != nil {
		return false, err_result
	}

	_, err_rows := result.RowsAffected()
	if err_rows != nil {
		return false, err_rows
	}

	return true, nil
}

func InsetPlaylist(db *sql.DB, input types.INSERT_PLAYLIST) (bool, error) {
	prepare, err_prepare := db.Prepare(vars.INSERT_PLAYLIST_STATEMENT)
	if err_prepare != nil {
		return false, nil
	}
	defer prepare.Close()

	result, err_execute := prepare.Exec(input.Nombre, input.Descripcion)
	if err_execute != nil {
		return false, err_execute
	}

	_, err_rows := result.RowsAffected()
	if err_rows != nil {
		return false, err_rows
	}

	return true, nil

}

func InserVideoPlaylist(db *sql.DB, input types.INSERT_MUSIC_PLAYLIST) (bool, error) {
	statement, err_statement := db.Prepare(vars.INSERT_VIDEOS_ON_PLAYLIST)
	if err_statement != nil {
		return false, nil
	}
	defer statement.Close()

	var count_playlist int
	exists_playlist := db.QueryRow(
		"SELECT count(*) FROM playlists where id = ?",
		input.Playlist_id,
	).Scan(&count_playlist)

	if exists_playlist != nil {
		return false, exists_playlist
	}
	if count_playlist == 0 {
		return false, fmt.Errorf("No se ha podido encontrar la playlist")
	}

	var count int
	err := db.QueryRow(
		"SELECT COUNT(*) FROM playlist_videos WHERE video = ? AND playlist_id = ?",
		input.Video,
		input.Playlist_id,
	).Scan(&count)

	if err != nil {
		return false, err
	}
	if count > 0 {
		return false, fmt.Errorf("El video ya existe en la playlist")
	}

	result, err_result := statement.Exec(input.Playlist_id, input.Video, input.Orden)
	if err_result != nil {
		return false, err_result
	}

	_, err_rows := result.RowsAffected()
	if err_rows != nil {
		return false, err_rows
	}

	return true, nil
}

func InsertMetricasVideos(db *sql.DB, input types.INSERTMETRICASVIDEOS) (bool, error) {
	prepare, err_prepare := db.Prepare(vars.INSERT_METRICAS_VIDEOS)
	if err_prepare != nil {
		return false, err_prepare
	}
	defer prepare.Close()

	result, err_result := prepare.Exec(input.Video)
	if err_result != nil {
		return false, err_prepare
	}

	_, err_afected := result.RowsAffected()
	if err_afected != nil {
		return true, err_afected
	}

	return true, nil
}

/* SELECTS */
func GETLASTPIDPATH(db *sql.DB) (*[]types.PROCESS_EXECUTE, error) {
	results, err := db.Query(vars.GET_LAST_TEN_PROCESS_EXECUTE)
	if err != nil {
		return nil, err
	}
	defer results.Close()

	x := []types.PROCESS_EXECUTE{}
	for results.Next() {
		v := types.PROCESS_EXECUTE{}
		err := results.Scan(&v.PID, &v.Path, &v.Estado, &v.Aggregate)
		if err != nil {
			log.Fatal(err)
		}
		x = append(x, v)
	}

	if err_result := results.Err(); err_result != nil {
		fmt.Println(err_result.Error())
		return nil, err_result
	}

	return &x, nil
}

func GetListOfPlaylist(db *sql.DB) (*[]types.PLAYLIST, error) {
	rows, err_ros := db.Query(vars.GET_LIST_PLAYIST)
	if err_ros != nil {
		return nil, err_ros
	}

	var listado []types.PLAYLIST = []types.PLAYLIST{}
	for rows.Next() {
		var video types.PLAYLIST = types.PLAYLIST{}
		if err_scan := rows.Scan(&video.ID, &video.Nombre, &video.Descripcion); err_scan != nil {
			continue
		}

		listado = append(listado, video)
	}

	if err_result := rows.Err(); err_result != nil {
		fmt.Println(err_result.Error())
		return nil, err_result
	}

	return &listado, nil
}

func GetListVideoByPlayList(db *sql.DB, idplaylist int64) ([]string, error) {
	var listados []string = []string{}
	resultados, err_r := db.Query(vars.GET_VIDEO_BY_PLAYLIST, idplaylist)
	if err_r != nil {
		return nil, err_r
	}

	for resultados.Next() {
		var value string
		resultados.Scan(&value)
		listados = append(listados, value)
	}

	if err_res := resultados.Err(); err_res != nil {
		return nil, err_res
	}

	return listados, nil
}

func GetMetricaFuncToday(db *sql.DB) (*[]types.METRICA_DATE_VIDEOS, error) {
	var result []types.METRICA_DATE_VIDEOS = []types.METRICA_DATE_VIDEOS{}
	rows, err_rows := db.Query(vars.GET_VIDEOS_TOP_TODAY)
	if err_rows != nil {
		return nil, err_rows
	}

	for rows.Next() {
		var v types.METRICA_DATE_VIDEOS
		rows.Scan(&v.Fecha, &v.Video, &v.Repoducc)
		result = append(result, v)
	}

	if err_t := rows.Err(); err_t != nil {
		return nil, err_t
	}

	return &result, nil
}
