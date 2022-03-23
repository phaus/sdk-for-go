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

    var service := appwrite.Storage{
        client: &client
    }

    var response, error := service.CreateBucket("[BUCKET_ID]", "[NAME]", "file", ["role:all"], ["role:all"], false, null, [], false, false)

    if error != nil {
        panic(error)
    }

    fmt.Println(response)
}
