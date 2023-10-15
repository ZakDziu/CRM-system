// nolint
package api

import (
	"crm-system/pkg/authmiddleware"
	"crm-system/pkg/store"
	"github.com/gin-gonic/gin"
	"testing"
)

func initTestAPI(t *testing.T, middleware authmiddleware.AuthMiddleware, postgres *store.PostgresStore) *api {
	t.Helper()

	gin.SetMode(gin.ReleaseMode)
	api := &api{
		router:        gin.New(),
		auth:          middleware,
		postgresStore: postgres,
	}

	api.router = configureRouter(api)

	return api
}

func makeList(f ...func([]interface{}, []interface{})) []func([]interface{}, []interface{}) {
	funcs := make([]func([]interface{}, []interface{}), 0)
	for _, i := range f {
		funcs = append(funcs, i)
	}
	return funcs
}
