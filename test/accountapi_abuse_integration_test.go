package main

import (
	accounts "accountapi_lib/src/go/v1"
	"testing"
)

// ================================ create fun =============================
func TestAbuseCaseCreateAccountWithInvalidOrganisationId(t *testing.T) {
  accountAttributes := new(accounts.AccountAttributes)

  _, err := accounts.Create(invalidOrganisationId, accountAttributes)

  if (err == nil) {
    t.Error("Expected error to be returned")
  }
}

func TestAbuseCaseCreateAccountWithNoCountry(t *testing.T) {
  accountAttributes := new(accounts.AccountAttributes)
  accountAttributes.Name = []string {"Samantha Holder"}  

  _, err := accounts.Create(organisationId, accountAttributes)

  if (err == nil) {
    t.Error("Expected error to be returned")
  }
}

func TestAbuseCaseCreateAccountWithNoName(t *testing.T) {
  country := "GB"

  accountAttributes := new(accounts.AccountAttributes)
  accountAttributes.Country = &country

  _, err := accounts.Create(organisationId, accountAttributes)

  if (err == nil) {
    t.Error("Expected error to be returned")
  }
}

func TestAbuseCaseCreateAccountWithTooManyNames(t *testing.T) {
  country := "GB"

  accountAttributes := new(accounts.AccountAttributes)
  accountAttributes.Country = &country
  accountAttributes.Name = []string {"Samantha Holder", "John Bel", "Ricky Martin", "Bryan Adams", "Mel Gibson"}

  _, err := accounts.Create(organisationId, accountAttributes)

  if (err == nil) {
    t.Error("Expected error to be returned")
  }
}

func TestAbuseCaseCreateAccountWithNilAttributes(t *testing.T) {
  _, err := accounts.Create(organisationId, nil)

  if (err == nil) {
    t.Error("Expected error to be returned")
  }
}

// ================================ fetch fun =============================
func TestAbuseCaseFetchNonExistingAccount(t *testing.T) {
  const nonExistingAccountId = "f56c75bf-69da-4861-97e8-18e43f9506df"

  _, err := accounts.Fetch(nonExistingAccountId)

  if (err == nil) {
    t.Error("Expected error to be returned")
  }
}

func TestAbuseCaseFetchInvalidAccount(t *testing.T) { 
  _, err := accounts.Fetch(invalidAccountId)

  if (err == nil) {
    t.Error("Expected error to be returned")
  }
}

// ================================ delete fun =============================
func TestAbuseCaseDeleteNonExistingAccount(t *testing.T) {
  const nonExistingAccountId = "f56c75bf-69da-4861-97e8-18e43f9506df"
  const version = 0
  
   _, err := accounts.Delete(nonExistingAccountId, version)

  if (err == nil) {
    t.Error("Expected error to be returned")
  }
}

func TestAbuseCaseDeleteInvalidAccount(t *testing.T) { 
  const version = 0

  _, err := accounts.Delete(invalidAccountId, version)

  if (err == nil) {
    t.Error("Expected error to be returned")
  }
}

func TestAbuseCaseDeleteInvalidVersion(t *testing.T) { 
  // create account
  country := "GB"
  accountAttributes := new(accounts.AccountAttributes)
  accountAttributes.Country = &country
  accountAttributes.Name = []string {"Samantha Holder"}

  resp, err := accounts.Create(organisationId, accountAttributes)

  if (err != nil || resp == nil || resp.Data.Version == nil) {
    t.Error("Something went wrong os setup level for this test - create account phase")
    return
  }

  accountId := resp.Data.ID
  version := *resp.Data.Version
  
  // delete invalid version
  _, err = accounts.Delete(accountId, version + 1)

  if (err == nil) {
    t.Error("Expected error to be returned")
  }

  // delete account
  _, _ = accounts.Delete(accountId, version)
}