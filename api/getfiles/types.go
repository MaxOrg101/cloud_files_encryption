package getfiles

type GetFilesRequest struct {
	Path string `json:"path" binding:"required"`
}

type file struct {
	Name  string `json:"name"`
	IsDir bool   `json:"isDir"`
}

type GetFilesResponse []file
