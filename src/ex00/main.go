package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v7"
)

func main() {
	elClient, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatal("Error creating client:", err)
	}
	index := "places"
	mapping := `
    {
      "settings": {
        "number_of_shards": 1
      },
      "mappings": {
        "properties": {
          "field1": {
            "type": "text"
          }
        }
      }
    }`
	slice := []string{index}
	_, err = elClient.Indices.Delete(slice, elClient.Indices.Delete.WithAllowNoIndices(true))
	if err != nil {
		log.Fatal(err)
	}
	res, err := elClient.Indices.Create(index, elClient.Indices.Create.WithBody(strings.NewReader(mapping)))
	if err != nil {
		log.Fatal()
	}
	fmt.Println(res.String())
}
