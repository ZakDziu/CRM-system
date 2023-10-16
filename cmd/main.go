package main

//nolint:stylecheck
import (
	"context"
	"crm-system/docs"
	"crm-system/pkg/api"
	"crm-system/pkg/authmiddleware/appauth"
	"crm-system/pkg/config"
	"crm-system/pkg/logger"
	"crm-system/pkg/store"
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

// @title           CRM System API
// @version         1.0
// @description     All handlers for the CRM System API
// @termsOfService  https://tos.santoshk.dev

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	conf, err := config.New()
	if err != nil {
		logger.Fatalf("Can't read config file: %s", err)
	}

	atKey, err := LoadKey(conf.Keys.AccessKey)
	if err != nil {
		logger.Fatalf("main.go--->main()--->LoadAccKey: %s", err)
	}

	rtKey, err := LoadKey(conf.Keys.RefreshKey)
	if err != nil {
		logger.Fatalf("main.go--->main()--->LoadRefKey: %s", err)
	}

	postgresStore, err := store.NewPostgres(&conf.DB)
	if err != nil {
		logger.Fatalf("main.go--->main()--->NewPostgres: %s", err)
	}

	middleware := appauth.NewAuthMiddleware(postgresStore, atKey, rtKey)

	apiServer := api.NewServer(&conf.Server, postgresStore, middleware)
	runErr := make(chan error, 1)
	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, syscall.SIGINT, syscall.SIGTERM)
	logger.Infof("Start api: %s", time.Now())

	docs.SwaggerInfo.Host = os.Getenv("API_URL")

	go func() {
		logger.Infof("HTTP API server start listen on: %s", apiServer.Addr)
		err = apiServer.ListenAndServe()
		if err != nil {
			runErr <- fmt.Errorf("can't start http server: %w", err)
		}
	}()

	select {
	case err = <-runErr:
		logger.Fatalf("Running error: %s", err)
	case s := <-quitCh:
		logger.Infof("Received signal: %v. Running graceful shutdown...", s)
		ctx := context.Background()

		err = apiServer.Shutdown(ctx)
		if err != nil {
			logger.Infof("Can't shutdown API server: %s", err)
		}
	}
}

func LoadKey(path string) (*ecdsa.PrivateKey, error) {
	txt, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(txt)
	x509Encoded := block.Bytes

	return x509.ParseECPrivateKey(x509Encoded)
}
