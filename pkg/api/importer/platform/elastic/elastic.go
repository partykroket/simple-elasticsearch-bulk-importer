package elastic

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	esimporter "github.com/stenraad/es-importer/pkg/utl/model"
)

func PostBulk(payload []byte, r *esimporter.BulkRequest) (int, error) {

	// Create http client
	client := http.Client{
		Timeout: time.Duration(5 * time.Second),
	}

	// Build request
	request, err := http.NewRequest(http.MethodPost, r.ElasticHost, bytes.NewBuffer(payload))
	if err != nil {
		return 0, err
	}
	request.Header.Set("Content-type", "application/json")

	// Do Request
	resp, err := client.Do(request)
	if err != nil {
		return resp.StatusCode, err
	}
	defer resp.Body.Close()

	// Read body for debugging purposes
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))

	// Return elastic response code
	return resp.StatusCode, nil
}
