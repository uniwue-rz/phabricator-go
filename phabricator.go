package phabricator

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Query is the base type for the given system
type Query struct {
	QueryType string
	Key       string
	Value     interface{}
}

// APIResult The result of the given API Query will be saved here first
type APIResult struct{
	Result interface{}
	Error_Code string
	Error_Info string
}

// Cursor is the type that always describes the cursor
type Cursor struct {
	Limit string
	After string
	Before string
	Order string
}

// CheckApiUrl checks for the given url that it is valid
func CheckApiUrl(url string) (res bool, err error) {
	_, err = http.Get(url)
	if err != nil {

		return false, err
	}

	return true, err
}

// SendApiQuery Creates a Query URL and send it to Conduit for the given values as a map string
// There are 4 query types available. string, array, map and MapArray
func SendApiQuery(apiUrl string, method string, values []Query) (res []byte, err error) {
	// Create the buffer for the concat
	var urlBuffer bytes.Buffer
	urlBuffer.WriteString(apiUrl)
	urlBuffer.WriteString(method)
	urlValues := url.Values{}
	for _, q := range values {
		if q.QueryType == "array" {
			for key, value := range q.Value.([]string) {
				urlKey := q.Key + "[" + strconv.Itoa(key) + "]"
				urlValues.Add(urlKey, value)
			}
		} else if q.QueryType == "map"{
			for key, value := range q.Value.(map[string]string) {
				urlKey := q.Key + "[" + key + "]"
				urlValues.Add(urlKey, value)
			}
		} else if q.QueryType == "mapArray"{
			for key, value := range q.Value.(map[string][]string){
				for i, insideValue :=range value{
					urlKey := q.Key + "[" + key + "]"+"[" + strconv.Itoa(i) +"]"
					urlValues.Add(urlKey, insideValue)
				}
			}
		} else {
			urlValues.Add(q.Key, q.Value.(string))
		}
	}
	valuesAsString := urlValues.Encode()
	//log.Println(valuesAsString)
	req, err := http.NewRequest("GET", urlBuffer.String(), strings.NewReader(valuesAsString))
	client := http.Client{}
	resp, err := client.Do(req)
	res, err = ioutil.ReadAll(resp.Body)

	resp.Body.Close()

	return res, err
}

// AddToString adds concats a string to another one
func AddToString(base string, toAdd string, atStart bool) string {
	var buffer bytes.Buffer
	if atStart == false {
		buffer.WriteString(base)
		buffer.WriteString(toAdd)
	} else {
		buffer.WriteString(toAdd)
		buffer.WriteString(base)
	}

	return buffer.String()
}

// TypeOf Returns the type of the given interface
func TypeOf(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

// CheckApiLogin tries to login to the Phabricator API server
func CheckApiLogin(apiUrl string, token string) (res bool, err error) {
	queryList := []Query{}
	queryList = AddToken(queryList, token)
	resp, err := SendApiQuery(apiUrl, "user.whoami", queryList)
	var m APIResult
	json.Unmarshal(resp, &m)
	if m.Result == "" && m.Error_Code != "" {
		errorsMessage := AddToString(m.Error_Code, " : ", false)
		errorsMessage = AddToString(errorsMessage, m.Error_Info, false)
		err = errors.New(errorsMessage)

		return false, err
	}
	return true, err
}

// AddToken adds token to the given query. It should be done for every query
func AddToken(queries []Query, token string) []Query {
	query := Query{"string", "api.token", token}
	queries = append(queries, query)

	return queries
}

// GetAvailableQueries returns the list of available methods from the application
func GetAvailableMethods(apiUrl string, token string) (res []string, err error) {
	queries := []Query{}
	queries = AddToken(queries, token)
	resp, err := SendApiQuery(apiUrl, "conduit.query", queries)

	var dat map[string]interface{}
	json.Unmarshal(resp, &dat)
	result := dat["result"].(map[string]interface{})
	for i := range result {
		res = append(res, i)
	}

	return res, err
}
