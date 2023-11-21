package api

//nolint:revive
import (
	_ "crm-system/docs"
	"crm-system/pkg/authmiddleware"
	"crm-system/pkg/config"
	"crm-system/pkg/store"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	*http.Server
}

type api struct {
	postgresStore *store.Store
	router        *gin.Engine
	config        *config.ServerConfig
	auth          authmiddleware.AuthMiddleware

	authHandler *AuthHandler
	userHandler *UserHandler
}

func NewServer(
	config *config.ServerConfig,
	postgresStore *store.Store,
	auth authmiddleware.AuthMiddleware,
) *Server {
	handler := newAPI(config, postgresStore, auth)

	srv := &http.Server{
		Addr:              config.ServerPort,
		Handler:           handler,
		ReadHeaderTimeout: config.ReadTimeout.Duration,
	}

	return &Server{
		Server: srv,
	}
}

func newAPI(
	config *config.ServerConfig,
	postgresStore *store.Store,
	auth authmiddleware.AuthMiddleware,
) *api {
	api := &api{
		config:        config,
		postgresStore: postgresStore,
		auth:          auth,
	}

	api.router = configureRouter(api)

	api.router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return api
}

//nolint:varnamelen
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding,"+
			"X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)

			return
		}

		c.Next()
	}
}

func (a *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}

func (a *api) Auth() *AuthHandler {
	if a.authHandler == nil {
		a.authHandler = NewAuthHandler(a)
	}

	return a.authHandler
}

func (a *api) User() *UserHandler {
	if a.userHandler == nil {
		a.userHandler = NewUserHandler(a)
	}

	return a.userHandler
}
