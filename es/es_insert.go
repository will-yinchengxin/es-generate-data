package es

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"time"
)

func InsertToES(esAddress, index string, records []map[string]interface{}) (string, error) {
	newIndex := getIndice(index)
	es, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{esAddress},
	})
	if err != nil {
		panic(fmt.Sprintf("Error creating the client: %s", err))
	}

	var buf bytes.Buffer
	for _, record := range records {
		meta := map[string]interface{}{
			"index": map[string]interface{}{
				"_index": newIndex,
			},
		}
		metaBytes, _ := json.Marshal(meta)
		dataBytes, _ := json.Marshal(record)

		buf.Write(metaBytes)
		buf.WriteString("\n")
		buf.Write(dataBytes)
		buf.WriteString("\n")
	}

	res, err := es.Bulk(bytes.NewReader(buf.Bytes()), es.Bulk.WithContext(context.Background()))
	if err != nil {
		return "", fmt.Errorf("failed to execute bulk request: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return "", fmt.Errorf("bulk request returned an error: %s", res.String())
	}

	return newIndex, nil
}

func getIndice(indexBase string) string {
	return fmt.Sprintf("%s-%s", indexBase, time.Now().Format("2006.01.02"))
}
