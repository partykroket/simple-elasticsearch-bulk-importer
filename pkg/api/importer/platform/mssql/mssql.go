package mssql

import (
	"database/sql"
	"fmt"
	"strings"

	esimporter "github.com/stenraad/es-importer/pkg/utl/model"

	_ "github.com/denisenkom/go-mssqldb"
)

func New(host, dbName, user, pass string) (*sql.DB, error) {
	// Create connection string
	dbType := "sqlserver"

	connStr, err := createConnectionString(host, dbName, user, pass, dbType)
	if err != nil {
		return nil, err
	}

	db, err := sql.Open(dbType, connStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func ExecuteSP(db *sql.DB, r *esimporter.BulkRequest, sp string) (*sql.Rows, error) {
	// Create stored procedure with all parameters, make sure query is dynamic
	query := fmt.Sprintf("EXEC [dbo].[%s] @Index = '%s' ,@Offset = %v, @Fetch = %v ", sp, r.ElasticIndex, r.Offset, r.Fetch)

	fmt.Println(query)

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func parseMSSQLHostPort(info string) (string, string) {
	host, port := "127.0.0.1", "1433"
	if strings.Contains(info, ":") {
		host = strings.Split(info, ":")[0]
		port = strings.Split(info, ":")[1]
	} else if strings.Contains(info, ",") {
		host = strings.Split(info, ",")[0]
		port = strings.TrimSpace(strings.Split(info, ",")[1])
	} else if len(info) > 0 {
		host = info
	}
	return host, port
}

// createConnectionString getting the database information in the correct format for the database type
func createConnectionString(host, dbName, user, pass, dbType string) (string, error) {
	var connStr string
	Param := "?"
	switch dbType {
	case "sqlserver":
		// "server=%s; port=%s; database=%s; user id=%s; password=%s;"
		host, port := parseMSSQLHostPort(host)
		connStr = fmt.Sprintf("server=%s; port=%s; database=%s; user id=%s; password=%s;",
			host, port, dbName, user, pass)
	case "mysql":
		// "user:pass@host:port"
		connStr = fmt.Sprintf("%s:%s@tcp(%s)/%charset=utf8mb4&parseTime=true",
			user, pass, host, dbName, Param)
	default:
		return "", fmt.Errorf("%s, is no known database type (compatible databases: postgres, sqlserver and mysql)", dbType)
	}
	return connStr, nil
}
