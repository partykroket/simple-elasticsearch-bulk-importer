// Copyright 2017 Emir Ribic. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// GORSK - Go(lang) restful starter kit
//
// API Docs for GORSK v1
//
// 	 Terms Of Service:  N/A
//     Schemes: http
//     Version: 2.0.0
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Emir Ribic <ribice@gmail.com> https://ribice.ba
//     Host: localhost:8080
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - bearer: []
//
//     SecurityDefinitions:
//     bearer:
//          type: apiKey
//          name: Authorization
//          in: header
//
// swagger:meta
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

	log := zlog.New()

	e := server.New()

	v1 := e.Group("/v1")

	it.NewHTTP(il.New(importer.Initialize(), log), v1)

	server.Start(e, &server.Config{
		Port:                cfg.Server.Port,
		ReadTimeoutSeconds:  cfg.Server.ReadTimeout,
		WriteTimeoutSeconds: cfg.Server.WriteTimeout,
		Debug:               cfg.Server.Debug,
	})

	return nil
}
