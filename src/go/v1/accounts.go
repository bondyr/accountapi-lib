package accounts

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
	// "reflect"
)

func createHttpClient() *http.Client {
    // let's fix default http client configuration problems:
    // timeout and DefaultMaxIdleConnsPerHost 

    var t = http.DefaultTransport.(*http.Transport).Clone()
    t.MaxIdleConns = 100
    t.MaxConnsPerHost = 100
    t.MaxIdleConnsPerHost = 100
	
    return &http.Client{
        Timeout:   5 * time.Second,
        Transport: t,
    }
}

var fullUrl = getFullUrl()
var httpClient = createHttpClient()

func getFullUrl() string {
    host := os.Getenv("ACCOUNT_API_LIB_HOST")
    if host == "" {
        host = "https://api.staging-form3.tech" // let's assume this is production ip
    }

    const endpoint = "/v1/organisation/accounts"
    fullUrl := host + endpoint
    return fullUrl
}

func prepareCreateAccountRequest(organisationId string, accountAttributes *AccountAttributes) *AccountStruct {
    accountStruct := new(AccountStruct)

    accountStruct.Data.Attributes = accountAttributes
    accountStruct.Data.OrganisationID = organisationId
    accountStruct.Data.Type = "accounts"
    accountStruct.Data.ID = uuid.New().String()

    return accountStruct
}

func jsonUnmarshalToAccountStruct(resp *http.Response) (*AccountStruct, error) {
    respBodyBytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    accountStruct := new(AccountStruct)
    json.Unmarshal(respBodyBytes, accountStruct)
 
    return accountStruct, nil
}

func isStatusCodeSuccess(status int) bool {
    return status >= 200 && status < 300
}

func Create(organisationId string, accountAttributes *AccountAttributes) (*AccountStruct, error) {
    url := fullUrl
    accountRequest := prepareCreateAccountRequest(organisationId, accountAttributes)
    jsonReq, err := json.Marshal(accountRequest)
    if err != nil {
        return nil, err
    }

    resp, err := httpClient.Post(url, "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if (!isStatusCodeSuccess(resp.StatusCode)) {
        return nil, errors.New(resp.Status)
    }

    return jsonUnmarshalToAccountStruct(resp)
}

func Fetch(account_id string) (*AccountStruct, error) {
    url := fullUrl + "/" + account_id

    resp, err := httpClient.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if (!isStatusCodeSuccess(resp.StatusCode)) {
        return nil, errors.New(resp.Status)
    }  

    return jsonUnmarshalToAccountStruct(resp)
}

func Delete(account_id string, version int64) (bool, error) {
    versionStr := strconv.FormatInt(int64(version), 10)
    url := fullUrl + "/" + account_id + "?version=" + versionStr
    req, err := http.NewRequest(http.MethodDelete, url, nil)
    if err != nil {
        return false, err
    }

    resp, err := httpClient.Do(req)
    if err != nil {
        return false, err
    }
    defer resp.Body.Close()

    if (!isStatusCodeSuccess(resp.StatusCode)) {
        return false, errors.New(resp.Status)
    }

    return true, nil
}
