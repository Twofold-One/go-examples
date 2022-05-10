package elasticsearch

import (
	"context"
	"fmt"
	"log"

	"github.com/olivere/elastic/v7"
)


const Mapping = `{
	"mappings": {
		"_doc": {
			"properties": {
				"type": {"type": "keyword"},
				"name": {"type": "text"},
				"lastname": {"type": "keyword"},
				"create_time": {"type": "date"},
				"update_time": {"type": "date"}
			}
		}
	}
}`

// 

func ElasticsearchExample() {
	// pass a context to execute each service
	ctx := context.Background()

	// obtain a client and connect
	client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		log.Fatal(err)
	}

	// ping the Elasticsearch server to get the version number
	info, code, err := client.Ping("http://127.0.0.1:9200").Do(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Elasticsearch returend with code %d and version %s\n", code, info.Version.Number)
}

// good guide:
// https://olivere.github.io/elastic/