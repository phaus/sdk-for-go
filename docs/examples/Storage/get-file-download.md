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

    var service := appwrite.Storage{
        client: &client
    }

    var response, error := service.GetFileDownload("[FILE_ID]")

    if error != nil {
        panic(error)
    }

    fmt.Println(response)
}
