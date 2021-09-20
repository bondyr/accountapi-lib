package main

import (
	accounts "accountapi_lib/src/go/v1"
	"fmt"
)

func main() {
    const organisationId = "cdb76e34-74a1-4901-8179-196e7b133d55"

    accountAttributes := new(accounts.AccountAttributes)
    accountAttributes.Name = []string {"Samantha Holder"}
    country := "GB"
    accountAttributes.Country = &country
    accountAttributes.BaseCurrency = "PLN"
    accountAttributes.BankID = "MYBANKID"

    fmt.Println("\n============ Creating account ==========")
    createResponse, _ := accounts.Create(organisationId, accountAttributes)
    fmt.Printf("- API Response struct: %+v\n", *createResponse)
    fmt.Printf("- API Response data: %+v\n", createResponse.Data)
    fmt.Printf("- API Response attributes: %+v\n", *createResponse.Data.Attributes)

    fmt.Println("\n============ Fetching account ==========")
    accountId := createResponse.Data.ID
    fetchResponse, _ := accounts.Fetch(accountId)
    fmt.Printf("- API Response struct: %+v\n", *fetchResponse)
    fmt.Printf("- API Response data: %+v\n", fetchResponse.Data)
    fmt.Printf("- API Response attributes: %+v\n", *fetchResponse.Data.Attributes)
    fmt.Printf("- API Response links: %+v\n", *fetchResponse.Links)

    fmt.Println("\n============ Deleting account ==========")
    version := *createResponse.Data.Version
    deleteResponse, _ := accounts.Delete(accountId, version)
    fmt.Printf("- Delete response: %+v\n", deleteResponse)
}
