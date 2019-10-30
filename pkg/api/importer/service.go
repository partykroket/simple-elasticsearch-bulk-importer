package importer

import (
	"github.com/labstack/echo"
	esimporter "github.com/stenraad/es-importer/pkg/utl/model"
)

// Service represents user application interface
type Service interface {
	Bulk(echo.Context, *esimporter.BulkRequest) ([]byte, int, error)
}

// New creates new user application service
func New() *Importer {
	return &Importer{}
}

// Initialize initalizes User application service with defaults
func Initialize() *Importer {
	return New()
}

// Importer represents Importer application service
type Importer struct {
}
