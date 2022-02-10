package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go"
)

func main() {
    var client := appwrite.Client{}

    client.SetEndpoint("https://[HOSTNAME_OR_IP]/v1") // Your API Endpoint
    client.SetProject("") // Your project ID
    client.SetJWT("") // Your secret JSON Web Token

    var service := appwrite.Account{
        client: &client
    }

    var response, error := service.UpdateVerification("[USER_ID]", "[SECRET]")

    if error != nil {
        panic(error)
    }

    fmt.Println(response)
}
