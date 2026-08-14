package types

type INSERTPID struct {
	Pid    int
	Path   string
	Estado string
}

type BODY_CREATE_PLAYLIST struct {
	Nombre      string `json:"nombre" binding:"required"`
	Descripcion string `json:"descripcion" binding:"required"`
}

type BODY_ADD_VIDEO_PLAYLIST struct {
	Titulo string `json:"titulo" binding:"required"`
}
