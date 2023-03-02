package uploadfile

import "mime/multipart"

type UploadFileRequest struct {
	Path string                `form:"path" binding:"required"`
	File *multipart.FileHeader `form:"file" binding:"required"`
}
