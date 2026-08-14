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
	}

	return &x, nil
}
