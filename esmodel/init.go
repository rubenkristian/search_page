package esmodel

import "github.com/elastic/go-elasticsearch/v8"

type EsModel struct {
	esClient *elasticsearch.Client
}

func New(_esClient *elasticsearch.Client) *EsModel {
	return &EsModel{
		esClient: _esClient,
	}
}
