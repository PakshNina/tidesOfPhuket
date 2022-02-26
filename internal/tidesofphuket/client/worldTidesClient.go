package client

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/errgo.v2/fmt/errors"

	"tidesOfPhuket/internal/tools"
)

type TidesExtreme struct {
	Date     int64  `json:"dt"`
	Height   float64 `json:"height"`
	TideType string  `json:"type"`
}

type TidesExtremeResult struct {
	Extremes []TidesExtreme `json:"extremes"`
}

type TidesError struct {
	Status json.Number `json:"status"`
	Error  string      `json:"error"`
}

// WorldTidesClient REST API client
type WorldTidesClient interface {
	Get(worldTidesMethod string, queryParams map[string]string) ([]byte, error)
	GetExtremes(date, lat, lon string) (*TidesExtremeResult, error)
}

type Connection struct {
	Url    string
	ApiKey string
}

type Client struct {
	conn   Connection
	client *http.Client
}

func NewWorldTidesClient(conn Connection) WorldTidesClient {
	return &Client{
		conn:   conn,
		client: &http.Client{},
	}
}

func (w Client) Get(worldTidesMethod string, queryParams map[string]string) ([]byte, error) {
	url := tools.CreateFullUrl(w.conn.Url, worldTidesMethod)
	for key, value := range queryParams {
		url = tools.AddToUrlParameter(url, key, value)
	}
	url = tools.AddToUrlParameter(url, "key", w.conn.ApiKey)
	req, reqErr := http.NewRequest("GET", url, nil)
	if reqErr != nil {
		return nil, reqErr
	}
	res, resErr := w.client.Do(req)
	if resErr != nil {
		return nil, resErr
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(res.Body)
	answer, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return answer, errors.Newf("Status code is not ok: %d", res.StatusCode)
	}
	return answer, nil
}

func (w Client) GetExtremes(date, lat, lon string) (*TidesExtremeResult, error) {
	queryParams := map[string]string{}
	queryParams["date"] = date
	queryParams["lat"] = lat
	queryParams["lon"] = lon
	answer, err := w.Get("extremes", queryParams)
	if err != nil {
		var errorResponse TidesError
		if jsonErr := json.Unmarshal(answer, &errorResponse); jsonErr != nil {
			return nil, errors.Newf("Error with unmarshalling json")
		}
		return nil, errors.Newf("Status code: %s, error message: %s", errorResponse.Status, errorResponse.Error)
	}
	var tidesExtremes TidesExtremeResult
	if jsonErr := json.Unmarshal(answer, &tidesExtremes); jsonErr != nil {
		return nil, jsonErr
	}
	return &tidesExtremes, nil
}
