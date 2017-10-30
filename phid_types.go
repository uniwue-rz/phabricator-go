package phabricator

// PHID is the construct to wrap the phid query results
type PHID struct {
	Result    map[string]PHIDObject `json:"result"`
	ErrorCode string                `json:"error_code"`
	ErrorInfo string                `json:"error_info"`
}

// PHIDObject every object in the result is an instance of this struct
type PHIDObject struct {
	PHID     string `json:"phid"`
	Uri      string `json:"uri"`
	TypeName string `json:"typeName"`
	PHIDType string `json:"type"`
	Name     string `json:"name"`
	FullName string `json:"fullName"`
	Status   string `json:"status"`
}
