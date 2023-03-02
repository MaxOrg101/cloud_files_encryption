package createdir

import (
	"net/http"
	"os"
	"path"

	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/create-dir")
	{
		g.POST("", createDir)
	}
}

func createDir(c *gin.Context) {
	var body CreateDirRequest
	err := c.BindJSON(&body)
	if err != nil {
		httpo.NewErrorResponse(http.StatusBadRequest, err.Error()).
			Send(c, http.StatusBadRequest)
		return
	}

	// user_id := c.GetString(pasetomiddleware.EmailIdInContext)
	file_path := path.Join("/workspace/storage-app", "user_id", body.Path)
	err = os.Mkdir(file_path, os.ModePerm)
	if err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, err.Error()).
			SendD(c)
		return
	}
	httpo.NewSuccessResponse(http.StatusOK, "dir created").
		SendD(c)

}
