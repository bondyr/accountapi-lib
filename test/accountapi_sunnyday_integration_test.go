package main

import (
	accounts "accountapi_lib/src/go/v1"
	"testing"
)

var accountLinksForFetchTesting *accounts.AccountLinks
var accountAttributesForFetchTesting *accounts.AccountAttributes

// ================================ create fun =============================
func TestShouldSuccessfullyCreateAccountWithRequiredAttributes(t *testing.T) {
  testFailed := false
  accountLinks := new(accounts.AccountLinks)

  country := "GB"

  accountAttributes := new(accounts.AccountAttributes)
  accountAttributes.Country = &country
  accountAttributes.Name = []string {"Samantha Holder"}

  resp, err := accounts.Create(organisationId, accountAttributes)

  if (err == nil) {
    testFailed = isTestFailed(accountLinks, organisationId, accountAttributes, resp)
  }

  if (err != nil) {
    t.Errorf("Received unexpected error: %s\n", err)    
  } else if (testFailed) {
    logExpectedAndReceived(t, accountLinks, organisationId, accountAttributes, resp)
    t.Error("Received unexpected response")
  } else {
    accountIdList = append(accountIdList, resp.Data.ID)
    versionList = append(versionList, *resp.Data.Version)
  
    accountLinksForFetchTesting = accountLinks
    accountAttributesForFetchTesting = accountAttributes
  }
}

func TestShouldSuccessfullyCreateAccountWithPostmanExample1(t *testing.T) {
  testFailed := false
  accountLinks := new(accounts.AccountLinks)  

  accountClassification := "Personal"
  country := "GB"
  
  accountAttributes := new(accounts.AccountAttributes)
  accountAttributes.AccountClassification = &accountClassification
  accountAttributes.AccountNumber = "10000004"
  accountAttributes.Country = &country
  accountAttributes.BankID = "400302"
  accountAttributes.BankIDCode = "GBDSC"
  accountAttributes.BaseCurrency = "GBP"
  accountAttributes.Bic = "NWBKGB42"
  accountAttributes.Iban = "GB28NWBK40030212764204"
  accountAttributes.Name = []string {"Samantha Holder"}  

  resp, err := accounts.Create(organisationId, accountAttributes)

  if (err == nil) {
    testFailed = isTestFailed(accountLinks, organisationId, accountAttributes, resp)
  }

  if (err != nil) {
    t.Errorf("Received unexpected error: %s\n", err)    
  } else if (testFailed) {
    logExpectedAndReceived(t, accountLinks, organisationId, accountAttributes, resp)
    t.Error("Received unexpected response")
  } else {
    accountIdList = append(accountIdList, resp.Data.ID)
    versionList = append(versionList, *resp.Data.Version)
  }
}

func TestShouldSuccessfullyCreateAccountWithPostmanExample2(t *testing.T) {
  testFailed := false
  accountLinks := new(accounts.AccountLinks)  

  accountClassification := "Personal"
  accountMatchingOptOut := false
  country := "GB"
  jointAccount := false

  accountAttributes := new(accounts.AccountAttributes)
  accountAttributes.AccountClassification = &accountClassification
  accountAttributes.AccountMatchingOptOut = &accountMatchingOptOut
  accountAttributes.AccountNumber = "10000004"
  accountAttributes.AlternativeNames = []string {"Sam Holder"}
  accountAttributes.Country = &country
  accountAttributes.BankID = "400302"
  accountAttributes.BankIDCode = "GBDSC"
  accountAttributes.BaseCurrency = "GBP"
  accountAttributes.Bic = "NWBKGB42"
  accountAttributes.Iban = "GB28NWBK40030212764204"
  accountAttributes.JointAccount = &jointAccount  
  accountAttributes.Name = []string {"Samantha Holder", "Sammy Holdy"}
  accountAttributes.SecondaryIdentification = "A1B2C3D4"

  resp, err := accounts.Create(organisationId, accountAttributes)

  if (err == nil) {
    testFailed = isTestFailed(accountLinks, organisationId, accountAttributes, resp)
  }

  if (err != nil) {
    t.Errorf("Received unexpected error: %s\n", err)    
  } else if (testFailed) {
    logExpectedAndReceived(t, accountLinks, organisationId, accountAttributes, resp)
    t.Error("Received unexpected response")
  } else {
    accountIdList = append(accountIdList, resp.Data.ID)
    versionList = append(versionList, *resp.Data.Version)
  }
}

func TestShouldSuccessfullyCreateAccountWithAnotherCountry(t *testing.T) {
  testFailed := false
  accountLinks := new(accounts.AccountLinks)  

  accountClassification := "Personal"
  country := "PL"
  
  accountAttributes := new(accounts.AccountAttributes)
  accountAttributes.AccountClassification = &accountClassification
  accountAttributes.Country = &country
  accountAttributes.BankID = "400302"
  accountAttributes.BankIDCode = "GBDSC"
  accountAttributes.BaseCurrency = "PLN"
  accountAttributes.Bic = "LHVBEE22"
  accountAttributes.Name = []string {"Samantha Holder", "Sammy Holdy", "James Bond", "Julia Roberts"}

  resp, err := accounts.Create(organisationId, accountAttributes)

  if (err == nil) {
    testFailed = isTestFailed(accountLinks, organisationId, accountAttributes, resp)
  }

  if (err != nil) {
    t.Errorf("Received unexpected error: %s\n", err)    
  } else if (testFailed) {
    logExpectedAndReceived(t, accountLinks, organisationId, accountAttributes, resp)
    t.Error("Received unexpected response")
  } else {
    accountIdList = append(accountIdList, resp.Data.ID)
    versionList = append(versionList, *resp.Data.Version)
  }
}

// =========================== fetch ==================================
func TestShouldSuccessfullyFetchAccount(t *testing.T) {
  testFailed := false
  
  if (len(accountIdList) == 0) {
    t.Log("Test skipped, no input data provided")
    return
  }

  resp, err := accounts.Fetch(accountIdList[0])

  if (err != nil) {
    testFailed = isInvalidUuid(resp.Data.ID)
  }

  if (err != nil) {
    t.Errorf("Received unexpected error: %s\n", err)    
  } else if (testFailed) {
    logExpectedAndReceived(t, accountLinksForFetchTesting, organisationId, accountAttributesForFetchTesting, resp)    
    t.Error("Received unexpected response")
  }  
}

// =========================== delete ==================================
func TestShouldSuccessfullyDeleteAccount(t *testing.T) {
  for idx, accountId := range accountIdList {
    resp, err := accounts.Delete(accountId, versionList[idx])
  
    if (err != nil || resp == false) {
      t.Errorf("Delete operation failed with resp: %t, err: %s", resp, err)
    }
  }
}