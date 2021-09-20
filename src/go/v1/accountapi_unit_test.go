package accounts

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/jarcoal/httpmock"
)

const organisationId = "cdb76e34-74a1-4901-8179-196e7b133d55"
const accountsDataType = "accounts"
const someAcountId = "f56c75bf-69da-4861-97e8-18e43f9506df"

func isInvalidUuid(u string) bool {
    _, err := uuid.Parse(u)
    return err != nil
}

func logExpectedAndReceived(t *testing.T, organisationId string,
                            accountAttributes *AccountAttributes, resp *AccountStruct) {
  t.Logf("Expected: nil, Received: %s\n", resp.Links)
  t.Logf("Expected: %s, Received: %s\n", organisationId, resp.Data.OrganisationID)
  t.Logf("Expected: %s, Received: %s\n", accountsDataType, resp.Data.Type)
  t.Logf("Expected ID in UUID format, Received: %s\n", resp.Data.ID)
  attributesExpected, _ := json.Marshal(accountAttributes)
  attributesReceived, _ := json.Marshal(resp.Data.Attributes)
  t.Logf("Expected: %s, Received: %s\n", attributesExpected, attributesReceived)
}


func TestPrepareCreateAccountRequest(t *testing.T) {
  accountAttributes := new(AccountAttributes)
  accountAttributes.Name = []string {"Samantha Holder"}
  country := "GB"
  accountAttributes.Country = &country
  accountAttributes.BaseCurrency = "PLN"
  accountAttributes.BankID = "MYBANKID"

  resp := prepareCreateAccountRequest(organisationId, accountAttributes)

  testFailed := resp.Links != nil ||
                resp.Data.OrganisationID != organisationId ||
                resp.Data.Type != "accounts" ||
                isInvalidUuid(resp.Data.ID) ||
                resp.Data.Attributes != accountAttributes

  if (testFailed) {
    logExpectedAndReceived(t, organisationId, accountAttributes, resp)
    t.Error("Received unexpected response")
  } 
}

func TestJsonUnmarshalToAccountStruct(t *testing.T) {
  accountStruct := AccountStruct{}
  accountStruct.Data.ID = someAcountId
  accountStructBytes, _ := json.Marshal(accountStruct)
  input := http.Response{
    Body: ioutil.NopCloser(bytes.NewBuffer(accountStructBytes)),
  }

  resp, _ := jsonUnmarshalToAccountStruct(&input)

  if (*resp != accountStruct) {
    attributesExpected, _ := json.Marshal(&accountStruct)
    attributesReceived, _ := json.Marshal(resp)
    t.Errorf("Received unexpected response: Expected: %s, Received: %s\n", attributesExpected, attributesReceived)    
  }
}

func TestGetFullUrlFromEnv(t *testing.T) {
  const expectedFullUrl = "www.example.com/v1/organisation/accounts"  
  os.Setenv("ACCOUNT_API_LIB_HOST", "www.example.com")
  
  fullUrl := getFullUrl()

  if (expectedFullUrl != fullUrl) {
    t.Errorf("Received unexpected url. Expected: %s, Received: %s\n", expectedFullUrl, fullUrl )
  }

  os.Unsetenv("ACCOUNT_API_LIB_HOST")
}

func TestGetFullUrlFromDefault(t *testing.T) {
  const expectedFullUrl = "https://api.staging-form3.tech/v1/organisation/accounts"

  fullUrl := getFullUrl()

  if (expectedFullUrl != fullUrl) {
    t.Errorf("Received unexpected url. Expected: %s, Received: %s\n", expectedFullUrl, fullUrl )
  }
}

func TestCreateError(t *testing.T) {
  accountAttributes := new(AccountAttributes)  
  
  httpmock.Activate()
  defer httpmock.DeactivateAndReset()

  _, err := Create(organisationId, accountAttributes)

  if (err == nil) {
    t.Error("Expected error to be returned")
  }
}
func TestFetchError(t *testing.T) {
  httpmock.Activate()
  defer httpmock.DeactivateAndReset()

  _, err := Fetch(someAcountId)

  if (err == nil) {
    t.Error("Expected error to be returned")
  }
}

func TestDeleteError(t *testing.T) {
  const version = 0
  
  httpmock.Activate()
  defer httpmock.DeactivateAndReset()

  _, err := Delete(someAcountId, version)

  if (err == nil) {
    t.Error("Expected error to be returned")
  }
}
