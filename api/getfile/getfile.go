package getfile

import (
	"io/ioutil"
	"net/http"
	"path"

	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/get-file")
	{
		g.POST("", getFile)
	}
}

func getFile(c *gin.Context) {
	var body GetFilesRequest
	err := c.BindJSON(&body)
	if err != nil {
		httpo.NewErrorResponse(http.StatusBadRequest, err.Error()).
			Send(c, http.StatusBadRequest)
		return
	}

	// user_id := c.GetString(pasetomiddleware.EmailIdInContext)
	file_path := path.Join("/workspace/storage-app", "user_id", body.FilePath)

	byteFile, err := ioutil.ReadFile(file_path)
	if err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, err.Error()).
			SendD(c)
		return
	}

	c.Header("Content-Disposition", "attachment; filename=file-name.txt")
	c.Data(http.StatusOK, "application/octet-stream", byteFile)
}
