// Package api provide support to create /api group
package api

import (
	"template-app/api/auth"
	"template-app/api/createdir"
	"template-app/api/getfile"
	"template-app/api/getfiles"
	"template-app/api/middleware/auth/pasetomiddleware"
	"template-app/api/publish"
	"template-app/api/uploadfile"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies the /api group and v1 routes to given gin Engine
func ApplyRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		auth.ApplyRoutes(api)
		api.Use(pasetomiddleware.PASETO)
		createdir.ApplyRoutes(api)
		getfiles.ApplyRoutes(api)
		getfile.ApplyRoutes(api)
		publish.ApplyRoutes(api)
		uploadfile.ApplyRoutes(api)
	}
}
