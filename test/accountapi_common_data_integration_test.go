package main

import (
	accounts "accountapi_lib/src/go/v1"
	"encoding/json"
	"reflect"
	"testing"

	"github.com/google/uuid"
)

const accountsDataType = "accounts"

const organisationId = "cdb76e34-74a1-4901-8179-196e7b133d55"
const invalidOrganisationId = "abc"
const invalidAccountId = "def"

var accountIdList []string
var versionList []int64

func isInvalidUuid(u string) bool {
    _, err := uuid.Parse(u)
    return err != nil
 }

func logExpectedAndReceived(t *testing.T, accountLinks *accounts.AccountLinks, organisationId string,
                            accountAttributes *accounts.AccountAttributes, resp *accounts.AccountStruct) {
  t.Logf("Expected: %s, Received: %s\n", accountLinks.Self, *resp.Links)
  t.Logf("Expected: %s, Received: %s\n", organisationId, resp.Data.OrganisationID)
  t.Logf("Expected: %s, Received: %s\n", accountsDataType, resp.Data.Type)
  t.Logf("Expected ID in UUID format, Received: %s\n", resp.Data.ID)
  attributesExpected, _ := json.Marshal(accountAttributes)
  attributesReceived, _ := json.Marshal(resp.Data.Attributes)
  t.Logf("Expected: %s, Received: %s\n", attributesExpected, attributesReceived)
}

func isTestFailed(accountLinks *accounts.AccountLinks, organisationId string,
                  accountAttributes *accounts.AccountAttributes, resp *accounts.AccountStruct) bool {
  accountLinks.Self = "/v1/organisation/accounts/" + resp.Data.ID                        
  
  testFailed := *resp.Links != *accountLinks ||
                resp.Data.OrganisationID != organisationId ||
                resp.Data.Type != accountsDataType ||
                isInvalidUuid(resp.Data.ID) ||
                resp.Data.CreatedOn == nil ||
                resp.Data.ModifiedOn == nil ||
                !reflect.DeepEqual(resp.Data.Attributes, accountAttributes)
  return testFailed
}   