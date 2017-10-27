package phabricator

import (
	"encoding/json"
	"errors"
)

// GetPhidByName Returns the PHID for the given object using the Monogram name.
// For example for PassPhrase items you should use K{int}
func GetPhidByName(apiUrl string, token string, name string) (phid string, err error) {
	queryList := []Query{}
	queryList = AddToken(queryList, token)
	queryList = append(queryList, Query{"array", "names", []string{name}})
	resp, err := SendApiQuery(apiUrl, "phid.lookup", queryList)
	var message APIResult
	err = json.Unmarshal(resp, &message)
	if err != nil {

		panic(err)
	}
	resultType := TypeOf(message.Result)
	if resultType == "[]interface {}" {

		err = errors.New("The given name '" + name + "' did not belong to any phid")

		return "", err
	}

	phid = message.Result.(map[string]interface{})[name].(map[string]interface{})["phid"].(string)

	return phid, err
}
