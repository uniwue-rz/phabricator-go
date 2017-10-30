package phabricator

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// NewRequest creates a new request for the given one
func NewRequest(apiUrl string, token string) Request {

	return Request{apiUrl, token, "", url.Values{}}
}

// SetMethod sets the method for the given Request
func (r *Request) SetMethod(method string) {
	r.Method = method
}

// AddValue adds a new value to the given request URL
func (r *Request) AddValues(values []Query) {
	for _, q := range values {
		if q.QueryType == "array" {
			for key, value := range q.Value.([]string) {
				urlKey := q.Key + "[" + strconv.Itoa(key) + "]"
				r.Values.Add(urlKey, value)
			}
		} else if q.QueryType == "map" {
			for key, value := range q.Value.(map[string]string) {
				urlKey := q.Key + "[" + key + "]"
				r.Values.Add(urlKey, value)
			}
		} else if q.QueryType == "mapArray" {
			for key, value := range q.Value.(map[string][]string) {
				for i, insideValue := range value {
					urlKey := q.Key + "[" + key + "]" + "[" + strconv.Itoa(i) + "]"
					r.Values.Add(urlKey, insideValue)
				}
			}
		} else if q.QueryType == "string" {
			r.Values.Add(q.Key, q.Value.(string))
		}
	}
}

// Reset restart the given request query string.
func (r *Request) Reset() {
	r.Values = url.Values{}
}

// Send sends the given request to the server. The result will be the error and response body bytes
func (r *Request) Send() (resp []byte, err error) {
	var urlBuffer bytes.Buffer
	urlBuffer.WriteString(r.Url)
	urlBuffer.WriteString(r.Method)
	r.Values.Add("api.token", r.Token)
	valuesAsString := r.Values.Encode()
	httpRequest, err := http.NewRequest("GET", urlBuffer.String(), strings.NewReader(valuesAsString))
	client := http.Client{}
	queryResult, err := client.Do(httpRequest)
	resp, err = ioutil.ReadAll(queryResult.Body)
	queryResult.Body.Close()
	// Always restart the request data so it can be reused with a new query
	r.Reset()
	return resp, err
}

// SendRequest send a request to the given phabricator server. And Returns the response as bytes array
func SendRequest(request *Request) (resp []byte, err error) {
	var urlBuffer bytes.Buffer
	urlBuffer.WriteString(request.Url)
	urlBuffer.WriteString(request.Method)
	request.Values.Add("api.token", request.Token)
	valuesAsString := request.Values.Encode()
	httpRequest, err := http.NewRequest("GET", urlBuffer.String(), strings.NewReader(valuesAsString))
	client := http.Client{}
	queryResult, err := client.Do(httpRequest)
	resp, err = ioutil.ReadAll(queryResult.Body)
	queryResult.Body.Close()
	// Always restart the request data so it can be reused with a new query
	request.Reset()
	return resp, err
}
