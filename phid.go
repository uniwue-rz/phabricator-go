package phabricator

import (
	"encoding/json"
)

// GetPhid Returns the PHID of given Monogram
func GetPhid(request *Request, name string) (phid PHID, err error) {
	queryList := []Query{}
	queryList = append(queryList, Query{"array", "names", []string{name}})
	request.Method = "phid.lookup"
	request.AddValues(queryList)
	resp, err := SendRequest(request)
	json.Unmarshal(resp, &phid)

	return phid, err
}

// GetName Returns the name of the given object by its PHID
func GetName(request *Request, phid string) (name PHID, err error) {
	queryList := []Query{}
	queryList = append(queryList, Query{"array", "phids", []string{phid}})
	request.Method = "phid.query"
	request.AddValues(queryList)
	resp, err := SendRequest(request)
	json.Unmarshal(resp, &name)

	return name, err
}

// ExtractPhid returns the PHID string for the given PHID object
func (phid *PHID) ExtractPhid(name string) string {
	for k, v := range phid.Result {
		if v.Name == k {

			return v.PHID
		}
	}

	return ""
}
