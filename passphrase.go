package phabricator

import (
	"encoding/json"
	"errors"
)

// GetAllPassPhrase returns the list of all available passPhrases to the given user
func GetAllPassPhrase(apiUrl string, token string) (passPhrase string, err error) {
	queryList := []Query{}
	queryList = AddToken(queryList, token)
	queryList = append(queryList, Query{"string", "needSecrets", "1"})
	queryList = append(queryList, Query{"string", "limit", "100"})
	resp, err := SendApiQuery(apiUrl, "passphrase.query", queryList)
	if err != nil {

		return "", err
	}

	return string(resp), err
}

// GetPassPhrase returns the passPhrase for the given name
// The name should be the Monogram and the number afterwards. Like K23
func GetPassPhrase(apiUrl string, token string, name string) (passPhrase string, err error) {
	phid, err := GetPhidByName(apiUrl, token, name)
	if err != nil {
		return passPhrase, err
	}
	queryList := []Query{}
	queryList = AddToken(queryList, token)
	queryList = append(queryList, Query{"string", "needSecrets", "1"})
	queryList = append(queryList, Query{"array", "phids", []string{phid}})
	resp, err := SendApiQuery(apiUrl, "passphrase.query", queryList)

	var m APIResult
	err = json.Unmarshal(resp, &m)
	if err != nil {

		panic(err)
	}
	resultType := TypeOf(m.Result)
	if resultType == "[]interface {}" {

		err = errors.New("The given name '" + name + "' did not have any passPhrase")

		return "", err
	}

	dataResult := m.Result.(map[string]interface{})["data"].(map[string]interface{})[phid].(map[string]interface{})["material"].(map[string]interface{})
	value, privateKeyOk := dataResult["privateKey"]
	if privateKeyOk {
		passPhrase = value.(string)
	}
	value, passPhraseOK := dataResult["password"]
	if passPhraseOK {
		passPhrase = value.(string)
	}

	return passPhrase, err
}

// GetUsername Returns the Username for the given passPhrase name
// This should be something starting with K and a number
func GetUsername(apiUrl string, token string, name string) (username string, err error) {
	phid, err := GetPhidByName(apiUrl, token, name)
	if err != nil {
		return username, err
	}
	queryList := []Query{}
	queryList = AddToken(queryList, token)
	queryList = append(queryList, Query{"array", "phids", []string{phid}})
	resp, err := SendApiQuery(apiUrl, "passphrase.query", queryList)
	var m APIResult
	err = json.Unmarshal(resp, &m)
	if err != nil {
		panic(err)
	}
	resultType := TypeOf(m.Result)
	if resultType == "[]interface {}" {
		err = errors.New("The given name '" + name + "' did not have any username")
		return "", err
	}
	dataResult := m.Result.(map[string]interface{})["data"].(map[string]interface{})[phid].(map[string]interface{})
	value, usernameOk := dataResult["username"]
	if usernameOk {

		return value.(string), err
	}
	return "", err
}
