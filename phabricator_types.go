package phabricator

import (
	"net/url"
	"sync"
)

// Request is the placeholder for the given Request to the Phabricator server
type Request struct {
	Url    string     // Url that should be used to for the phabricator api
	Token  string     // Token that should be used for the given request
	Method string     // The method that should be used for the request
	Values Values // The Values that should be parsed to the given URL string
}

type Values struct {
	sync.RWMutex
	val url.Values
}

// Query is the base type for the given system
type Query struct {
	QueryType string
	Key       string
	Value     interface{}
}

// Cursor is the type that always describes the cursor
type Cursor struct {
	Limit  string
	After  string
	Before string
	Order  string
}
