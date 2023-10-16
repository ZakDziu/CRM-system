package api

import (
	"crm-system/pkg/model"

	"github.com/gin-gonic/gin"

	"net/http"
)

func configureRouter(api *api) *gin.Engine {
	router := gin.Default()

	router.Use(CORSMiddleware())

	public := router.Group("api/v1")

	public.POST("/login", api.Auth().Login)
	public.POST("/refresh", api.Auth().Refresh)
	public.PATCH("/change-password", api.Auth().ChangePassword)

	private := router.Group("api/v1")

	private.Use(api.auth.Authorize)

	private.POST("/registration", api.Auth().Register)

	privateUser := private.Group("/user")

	privateUser.PATCH("/update-info", api.User().UpdateInfo)
	privateUser.GET("/", api.User().Get)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, model.ErrRecordNotFound)
	})

	return router
}
