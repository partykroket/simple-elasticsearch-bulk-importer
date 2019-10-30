package importer

import (
	"github.com/stenraad/es-importer/pkg/api/importer"
	esimporter "github.com/stenraad/es-importer/pkg/utl/model"
)

// New creates new importer logging service
func New(svc importer.Service, logger esimporter.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger:  logger,
	}
}

// LogService represents user logging service
type LogService struct {
	importer.Service
	logger esimporter.Logger
}

const name = "importer"
