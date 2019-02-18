package phabricator

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// NewRequest creates a new request for the given one
func NewRequest(apiUrl string, token string) Request {

	return Request{apiUrl, token, "", Values{val:url.Values{}}}
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
				r.Values.Lock()
				r.Values.val.Add(urlKey, value)
				r.Values.Unlock()
			}
		} else if q.QueryType == "map" {
			for key, value := range q.Value.(map[string]string) {
				urlKey := q.Key + "[" + key + "]"
				r.Values.Lock()
				r.Values.val.Add(urlKey, value)
				r.Values.Unlock()
			}
		} else if q.QueryType == "mapArray" {
			for key, value := range q.Value.(map[string][]string) {
				for i, insideValue := range value {
					urlKey := q.Key + "[" + key + "]" + "[" + strconv.Itoa(i) + "]"
					r.Values.Lock()
					r.Values.val.Add(urlKey, insideValue)
					r.Values.Unlock()
				}
			}
		} else if q.QueryType == "string" {
			r.Values.Lock()
			r.Values.val.Add(q.Key, q.Value.(string))
			r.Values.Unlock()
		}
	}
}

// Reset restart the given request query string.
func (r *Request) Reset() {
	r.Values = Values{val:url.Values{}}
}

// Send sends the given request to the server. The result will be the error and response body bytes
func (r *Request) Send() (resp []byte, err error) {
	var urlBuffer bytes.Buffer
	urlBuffer.WriteString(r.Url)
	urlBuffer.WriteString(r.Method)
	r.Values.Lock()
	r.Values.val.Add("api.token", r.Token)
	r.Values.Unlock()
	valuesAsString := r.Values.val.Encode()
	fmt.Println(valuesAsString)
	httpRequest, err := http.NewRequest("GET", urlBuffer.String(), strings.NewReader(valuesAsString))
	client := http.Client{}
	queryResult, err := client.Do(httpRequest)
	resp, err = ioutil.ReadAll(queryResult.Body)
	err = queryResult.Body.Close()
	// Always restart the request data so it can be reused with a new query
	r.Reset()
	return resp, err
}

// SendRequest send a request to the given phabricator server. And Returns the response as bytes array
func SendRequest(request *Request) (resp []byte, err error) {
	var urlBuffer bytes.Buffer
	urlBuffer.WriteString(request.Url)
	urlBuffer.WriteString(request.Method)
	request.Values.Lock()
	request.Values.val.Add("api.token", request.Token)
	request.Values.Unlock()
	valuesAsString := request.Values.val.Encode()
	httpRequest, err := http.NewRequest("GET", urlBuffer.String(), strings.NewReader(valuesAsString))
	client := http.Client{}
	queryResult, err := client.Do(httpRequest)
	resp, err = ioutil.ReadAll(queryResult.Body)
	err = queryResult.Body.Close()
	// Always restart the request data so it can be reused with a new query
	request.Reset()
	return resp, err
}
