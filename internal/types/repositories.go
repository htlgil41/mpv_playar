package types

import "time"

/* RESPONSES */
type VIDEOS_RESPONSES struct {
	Id             int64
	Titulo         string
	Descripcion    string
	Nombre_archivo string
	Creado_en      time.Time
}

type PROCESS_EXECUTE struct {
	PID       int
	Path      string
	Estado    string
	Aggregate time.Time
}

type VIDEOS struct {
	Titulo         string
	Descripcion    string
	Nombre_archivo string
}
