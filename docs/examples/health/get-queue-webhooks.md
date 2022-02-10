package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go"
)

func main() {
    var client := appwrite.Client{}

    client.SetEndpoint("https://[HOSTNAME_OR_IP]/v1") // Your API Endpoint
    client.SetProject("") // Your project ID
    client.SetKey("") // Your secret API key

    var service := appwrite.Health{
        client: &client
    }

    var response, error := service.GetQueueWebhooks()

    if error != nil {
        panic(error)
    }

    fmt.Println(response)
}
