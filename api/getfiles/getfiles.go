package getfiles

import (
	"net/http"
	"os"
	"path"

	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/TheLazarusNetwork/go-helpers/logo"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/get-files")
	{
		g.POST("", getFiles)
	}
}

func getFiles(c *gin.Context) {
	var body GetFilesRequest
	err := c.BindJSON(&body)
	if err != nil {
		httpo.NewErrorResponse(http.StatusBadRequest, err.Error()).
			Send(c, http.StatusBadRequest)
		return
	}

	// user_id := c.GetString(pasetomiddleware.EmailIdInContext)
	dir_path := path.Join("/workspace/storage-app", "user_id", body.Path)
	logo.Info(dir_path)
	entries, err := os.ReadDir(dir_path)
	if err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, err.Error()).
			SendD(c)
		return
	}

	files := make([]file, len(entries))

	for i, v := range entries {
		is_dir := v.IsDir()
		is_name := v.Name()
		files[i].IsDir = is_dir
		files[i].Name = is_name
	}

	httpo.NewSuccessResponseP(http.StatusOK, "files fetched", files).
		SendD(c)
}
