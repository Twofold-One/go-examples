package datastorage

import "github.com/Twofold-One/go-examples/basic/datastorage/elasticsearch"

// "github.com/Twofold-One/go-examples/basic/datastorage/mongodb"
// "github.com/Twofold-One/go-examples/basic/datastorage/csv"
// "github.com/Twofold-One/go-examples/basic/datastorage/postgres"

// to get Go sql drivers
// https://github.com/golang/go/wiki/SQLDrivers

func DatastorageExamples() {
	// csv.CSVStorageExample()
	// postgres.PostgresExample()
	// mongodb.MongodbExample()
	elasticsearch.ElasticsearchExample()
}