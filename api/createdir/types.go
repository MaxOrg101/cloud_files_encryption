package createdir

type CreateDirRequest struct {
	Path string `json:"path" binding:"required"`
}
