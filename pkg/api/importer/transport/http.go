package transport

import (
	"github.com/labstack/echo"
	"github.com/stenraad/es-importer/pkg/api/importer"
	esimporter "github.com/stenraad/es-importer/pkg/utl/model"
)

// HTTP represents user http service
type HTTP struct {
	svc importer.Service
}

// NewHTTP creates new user http service
func NewHTTP(svc importer.Service, er *echo.Group) {
	h := HTTP{svc}
	ig := er.Group("/bulk")
	// swagger:route POST /v1/users users userCreate
	// Creates new user account.
	// responses:
	//  200: userResp
	//  400: errMsg
	//  401: err
	//  403: errMsg
	//  500: err
	ig.POST("", h.bulk)

	ig.GET("", h.bulk)
}

func (h *HTTP) bulk(c echo.Context) error {
	r := new(esimporter.BulkRequest)

	r.DBHost = "localhost"
	r.DBUser = "sa"
	r.DBPass = "SQLServerTest@1991"
	r.DBName = "TESTDB"
	r.ElasticHost = "http://localhost:9200/_bulk"
	r.ElasticIndex = "date_dimension"
	r.ElasticBulkType = "create"
	r.Fetch = 10
	r.Offset = 0

	if err := c.Bind(r); err != nil {
		return err
	}

	elasticPayload, elasticStatusCode, err := h.svc.Bulk(c, r)
	if err != nil {
		return err
	}

	return c.JSONBlob(elasticStatusCode, elasticPayload)
}
