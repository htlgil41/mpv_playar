package types

import "time"

type INSERTPID struct {
	Pid    int
	Path   string
	Estado string
}

type INSERTMETRICASVIDEOS struct {
	Video string
}

type RANGEBEETWENDATEREPOSITORIE struct {
	Lte time.Time
	Gte time.Time
}

type BODY_CREATE_PLAYLIST struct {
	Nombre      string `json:"nombre" binding:"required"`
	Descripcion string `json:"descripcion" binding:"required"`
}

type BODY_ADD_VIDEO_PLAYLIST struct {
	Titulo string `json:"titulo" binding:"required"`
}

type BODY_ADD_MUSIC_PLAYLIST struct {
	Playlist_id int64  `json:"playlist" binding:"required"`
	Video       string `json:"video" binding:"required"`
	Orden       int    `json:"orden" binding:"required"`
}

type BODY_PLAY_PLAYLIST struct {
	Playlist_id int64 `json:"playlist" binding:"required"`
}

type BODY_DELETE_PLAYLIST struct {
	Playlist_id int64 `json:"playlist" binding:"required"`
}

type PARAM_GET_METRICAS_RANGES struct {
	Gte time.Time `form:"gte" time_format:"2006-01-02" binding:"required"`
	Lte time.Time `form:"lte" time_format:"2006-01-02" binding:"required"`
}
