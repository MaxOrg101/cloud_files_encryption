package uploadfile

import (
	"net/http"
	"path"
	"template-app/api/middleware/auth/pasetomiddleware"

	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/TheLazarusNetwork/go-helpers/logo"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/upload-file")
	{
		g.POST("", uploadFile)
	}
}

func uploadFile(c *gin.Context) {
	var request UploadFileRequest
	err := c.Bind(&request)
	if err != nil {
		httpo.NewErrorResponse(http.StatusBadRequest, err.Error()).
			Send(c, http.StatusBadRequest)
		return
	}
	user_id := c.GetString(pasetomiddleware.EmailIdInContext)
	dir_path := path.Join("/workspace/storage-app", user_id, request.Path, request.File.Filename)
	logo.Info("max file", dir_path)
	err = c.SaveUploadedFile(request.File, dir_path)

	if err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, err.Error()).
			SendD(c)
		return
	}
	httpo.NewSuccessResponse(http.StatusOK, "file uploaded").
		SendD(c)

}
