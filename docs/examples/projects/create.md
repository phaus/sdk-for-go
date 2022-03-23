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

    var service := appwrite.Projects{
        client: &client
    }

    var response, error := service.Create("[PROJECT_ID]", "[NAME]", "[TEAM_ID]", "[DESCRIPTION]", "[LOGO]", "https://example.com", "[LEGAL_NAME]", "[LEGAL_COUNTRY]", "[LEGAL_STATE]", "[LEGAL_CITY]", "[LEGAL_ADDRESS]", "[LEGAL_TAX_ID]")

    if error != nil {
        panic(error)
    }

    fmt.Println(response)
}
