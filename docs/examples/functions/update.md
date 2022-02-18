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
    client.SetKey("") // Your secret API key

    var service := appwrite.Functions{
        client: &client
    }

    var response, error := service.Update("[FUNCTION_ID]", "[NAME]", [], nil, [], "", 1)

    if error != nil {
        panic(error)
    }

    fmt.Println(response)
}
