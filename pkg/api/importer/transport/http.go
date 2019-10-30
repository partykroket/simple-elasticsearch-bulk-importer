package transport

import (
	"github.com/labstack/echo"
	"github.com/stenraad/es-importer/pkg/api/importer"
	esimporter "github.com/stenraad/es-importer/pkg/utl/model"
)

// HTTP represents importer http service
type HTTP struct {
	svc                    importer.Service
	DefaultDbhost          string
	DefaultElasticBulkType string
	DefaultElasticHost     string
	DefaultFetch           int
	DefaultOffset          int
}

// NewHTTP creates new importer http service
func NewHTTP(svc importer.Service, er *echo.Group, DefaultDbhost, DefaultElasticBulkType, DefaultElasticHost string, DefaultFetch, DefaultOffset int) {
	h := HTTP{
		svc:                    svc,
		DefaultDbhost:          DefaultDbhost,
		DefaultElasticBulkType: DefaultElasticBulkType,
		DefaultElasticHost:     DefaultElasticHost,
		DefaultFetch:           DefaultFetch,
		DefaultOffset:          DefaultOffset,
	}
	ig := er.Group("/bulk")
	ig.POST("", h.bulk)
	ig.GET("", h.bulk)
}

func (h *HTTP) bulk(c echo.Context) error {
	r := new(esimporter.BulkRequest)

	// Set defaults
	r.DBHost = h.DefaultDbhost
	r.ElasticHost = h.DefaultElasticHost
	r.ElasticBulkType = h.DefaultElasticBulkType
	r.Fetch = h.DefaultFetch
	r.Offset = h.DefaultOffset

	// Bind form params or json payload
	if err := c.Bind(r); err != nil {
		return err
	}

	// Perform the elastic bulk upload
	elasticPayload, elasticStatusCode, err := h.svc.Bulk(c, r)
	if err != nil {
		return err
	}

	// Return bulk payload query
	return c.JSONBlob(elasticStatusCode, elasticPayload)
}
