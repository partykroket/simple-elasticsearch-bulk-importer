package esimporter

type BulkRequest struct {
	DBHost          string `json:"db_host"`
	DBName          string `json:"db_name"`
	DBUser          string `json:"db_user"`
	DBPass          string `json:"db_pass"`
	ElasticBulkType string `json:"elastic_bulk_type"`
	ElasticHost     string `json:"elastic_host"`
	ElasticIndex    string `json:"elastic_index"`
	Fetch           int    `json:"fetch"`
	Offset          int    `json:"offset"`
}

type ElasticBulkRequest struct {
	DocID            string `json:"doc_id"`
	BulkInstructions []byte `json:"bulk_instructions"`
	Source           []byte `json:"source"`
}
