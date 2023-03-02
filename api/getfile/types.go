package getfile

type GetFilesRequest struct {
	FilePath string `json:"filePath"`
}

type file struct {
	Name  string `json:"name"`
	IsDir bool   `json:"isFolder"`
}

type GetFilesResponse struct {
	Files []file
}
