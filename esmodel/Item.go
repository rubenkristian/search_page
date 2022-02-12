package esmodel

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/elastic/go-elasticsearch/v8/esutil"
)

type Item struct {
	Title string `json:"title"`
}

func (es *EsModel) SearchItem(from int, size int) ([]interface{}, error) {
	var (
		err  error
		hits []interface{}
	)
	type QueryString struct {
		Query string `json:"query"`
	}

	type Query struct {
		QueryString QueryString `json:"query_string"`
	}

	type Structure struct {
		Query Query `json:"query"`
	}

	query := Structure{
		Query: Query{
			QueryString: QueryString{
				Query: "*",
			},
		},
	}

	res, err := es.esClient.Search(
		es.esClient.Search.WithContext(context.Background()),
		es.esClient.Search.WithIndex("item"),
		es.esClient.Search.WithBody(esutil.NewJSONReader(query)),
		es.esClient.Search.WithTrackTotalHits(true),
		es.esClient.Search.WithPretty(),
		es.esClient.Search.WithFrom(from),
		es.esClient.Search.WithSize(size),
	)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var e map[string]interface{}

	if res.IsError() {
		err = json.NewDecoder(res.Body).Decode(&e)

		if err != nil {
			return nil, err
		}
	}

	if err = json.NewDecoder(res.Body).Decode(&e); err == nil {
		hits = e["hits"].(map[string]interface{})["hits"].([]interface{})
	}

	return hits, err
}

func (es *EsModel) CreateIndex(id string, title string) (string, error) {
	var result string
	var err error

	doc := Item{Title: title}

	req := esapi.IndexRequest{
		Index:      "item",
		DocumentID: id,
		Body:       esutil.NewJSONReader(doc),
		Refresh:    "wait_for",
	}

	res, err := req.Do(context.Background(), es.esClient)

	if err != nil {
		return result, err
	}

	defer res.Body.Close()

	if res.IsError() {
		err = fmt.Errorf("[%s]", res)
	} else {
		// Deserialize the response into a map.
		var r map[string]interface{}
		if errjson := json.NewDecoder(res.Body).Decode(&r); errjson != nil {
			err = fmt.Errorf("error parsing the response body: %s", errjson)
		} else {
			// Print the response status and indexed document version.
			result = fmt.Sprintf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
		}
	}

	return result, err
}
