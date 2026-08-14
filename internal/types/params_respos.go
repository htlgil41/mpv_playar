package types

type INSERTPID struct {
	Pid    int
	Path   string
	Estado string
}

type BODY_CREATE_NEW_VIDEO struct {
	Titulo         string `json:"titulo" binding:"required"`
	Descripcion    string `json:"descripcion" binding:"required"`
	Nombre_archivo string `json:"nombre_archivo" binding:"required"`
}

type BODY_ADD_VIDEO_PLAYLIST struct {
	Titulo string `json:"titulo" binding:"required"`
}
