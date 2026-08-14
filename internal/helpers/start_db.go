package helpers

import (
	"database/sql"
	"playar/internal/vars"
)

func InitDbCreateRegisters(db *sql.DB) (bool, error) {
	result, err_execute := db.Exec(vars.MIGRATION_OOONE)
	if err_execute != nil {
		return false, err_execute
	}

	rows, err_rows := result.RowsAffected()
	if err_rows != nil {
		return false, err_rows
	}

	if rows != 0 {
		return true, nil
	}

	return false, nil
}
