package main

import (
    "fmt"
    "time"
    "github.com/appwrite/sdk-for-go"
)

func main() {
    var client := appwrite.NewClient(10 * time.Second)

    client.SetEndpoint("https://[HOSTNAME_OR_IP]/v1") // Your API Endpoint
    client.SetProject("") // Your project ID

    var service := appwrite.Database{
        client: &client
    }

    var response, error := service.CreateDocument("[COLLECTION_ID]", "[DOCUMENT_ID]", nil, ["role:all"], ["role:all"])

    if error != nil {
        panic(error)
    }

    fmt.Println(response)
}
