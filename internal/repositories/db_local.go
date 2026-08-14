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
