package main

import (
	"log"
	"os"
	"time"

	"github.com/appwrite/sdk-for-go/appwrite"
)

const (
	EmptyParentDocument     = ""
	EmptyParentProperty     = ""
	EmptyParentPropertyType = ""
)

func main() {
	var EmptyArray = []interface{}{}

	client := appwrite.NewClient(10 * time.Second)
	client.SetEndpoint(os.Getenv("YOUR_ENDPOINT"))
	client.SetProject(os.Getenv("YOUR_PROJECT_ID"))
	client.SetKey(os.Getenv("YOUR_KEY"))

	db := appwrite.NewDatabase(client)
	data := map[string]string{
		"hello": "world",
	}
	doc, err := db.CreateDocument(
		os.Getenv("COLLECTION_ID"),
		data,
		EmptyArray,
		EmptyArray,
		EmptyParentDocument,
		EmptyParentProperty,
		EmptyParentPropertyType,
	)
	if err != nil {
		log.Printf("Error creating document: %v", err)
	}
	log.Printf("Created document: %v", doc)
}
