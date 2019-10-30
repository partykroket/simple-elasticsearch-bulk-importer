package api

import (
	"github.com/stenraad/es-importer/pkg/utl/zlog"

	"github.com/stenraad/es-importer/pkg/api/importer"
	il "github.com/stenraad/es-importer/pkg/api/importer/logging"
	it "github.com/stenraad/es-importer/pkg/api/importer/transport"

	"github.com/stenraad/es-importer/pkg/utl/config"
	"github.com/stenraad/es-importer/pkg/utl/server"
)

// Start starts the API service
func Start(cfg *config.Configuration) error {
	// New log instance
	log := zlog.New()

	// Init echo server
	e := server.New()
	v1 := e.Group("/v1")

	// Init bulk call
	it.NewHTTP(il.New(importer.Initialize(), log), v1, cfg.App.DefaultDbhost, cfg.App.DefaultElasticBulkType, cfg.App.DefaultElasticHost, cfg.App.DefaultFetch, cfg.App.DefaultOffset)

	// Start server
	server.Start(e, &server.Config{
		Port:                cfg.Server.Port,
		ReadTimeoutSeconds:  cfg.Server.ReadTimeout,
		WriteTimeoutSeconds: cfg.Server.WriteTimeout,
		Debug:               cfg.Server.Debug,
	})

	return nil
}
