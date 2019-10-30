// package importer contains user application services
package importer

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/stenraad/es-importer/pkg/api/importer/platform/elastic"
	"github.com/stenraad/es-importer/pkg/api/importer/platform/mssql"
	esimporter "github.com/stenraad/es-importer/pkg/utl/model"
)

// Create creates a new user account
func (u *Importer) Bulk(c echo.Context, r *esimporter.BulkRequest) ([]byte, int, error) {

	// Init DB connection, maybe in go routine
	db, err := mssql.New(r.DBHost, r.DBName, r.DBUser, r.DBPass)
	if err != nil {
		return nil, 0, err
	}

	// Execute bulk stored procedure
	rows, err := mssql.ExecuteSP(db, r, "GEN_Elastic_Bulk")
	if err != nil {
		return nil, 0, err
	}

	var payload []byte

	for rows.Next() {
		r := esimporter.ElasticBulkRequest{}

		err := rows.Scan(&r.DocID, &r.BulkInstructions, &r.Source)
		if err != nil {
			return nil, 0, err
		}

		// Get the bulk request in the correct order
		payload = append(payload, r.BulkInstructions...)
		payload = append(payload, []byte(fmt.Sprintf("\n"))...)
		payload = append(payload, r.Source...)
		payload = append(payload, []byte(fmt.Sprintf("\n"))...)

	}

	statusCode, err := elastic.PostBulk(payload, r)
	if err != nil {
		return nil, 0, err
	}

	return payload, statusCode, nil
}
