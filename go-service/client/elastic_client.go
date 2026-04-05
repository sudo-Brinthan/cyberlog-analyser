package client

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

var es *elasticsearch.Client

// InitElastic initializes the connection to Elasticsearch
func InitElastic(url string) {
	cfg := elasticsearch.Config{
		Addresses: []string{url},
	}
	var err error
	es, err = elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("❌ Error creating the elastic client: %s", err)
	}

	res, err := es.Info()
	if err != nil {
		log.Fatalf("❌ Error getting response from elasticsearch: %s", err)
	}
	defer res.Body.Close()
	log.Println("✅ Connected to Elasticsearch")
}

// SaveLog saves the AI response to the "cyberlogs" index
func SaveLog(aiResult map[string]interface{}) {
	if es == nil {
		log.Println("⚠️ Elasticsearch client not initialized")
		return
	}

	// Add a timestamp so we can plot it on a graph in Kibana
	aiResult["@timestamp"] = time.Now().Format(time.RFC3339)

	data, err := json.Marshal(aiResult)
	if err != nil {
		log.Printf("❌ Error marshaling document: %s", err)
		return
	}

	req := esapi.IndexRequest{
		Index:   "cyberlogs",
		Body:    bytes.NewReader(data),
		Refresh: "true",
	}

	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Printf("❌ Error indexing document: %s", err)
		return
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("❌ Error indexing document: %s", res.String())
	} else {
		log.Println("💾 Successfully saved threat data to Elasticsearch")
	}
}
