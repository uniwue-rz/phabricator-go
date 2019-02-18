package phabricator

import (
	"encoding/json"
)

// GetAllPassPhrase returns the list of all available passPhrases to the given user
func GetAllPassPhrase(request *Request) (passPhrases PassPhrase, err error) {
	queryList := make([]Query,0)
	queryList = append(queryList, Query{"string", "needSecrets", "1"})
	queryList = append(queryList, Query{"string", "limit", "100"})
	request.AddValues(queryList)
	request.SetMethod("passphrase.query")
	resp, err := request.Send()
	err = json.Unmarshal(resp, &passPhrases)

	return passPhrases, err
}

// GetPassPhrase returns the passPhrase for the given name
// The name should be the Monogram and the number afterwards. Like K23
func GetPassPhrase(request *Request, name string) (passPhrase PassPhrase, err error) {
	phid, err := GetPhid(request, name)
	if err != nil {
		return passPhrase, err
	}
	queryList := make([]Query,0)
	queryList = append(queryList, Query{"string", "needSecrets", "1"})
	queryList = append(queryList, Query{"array", "phids", []string{phid.ExtractPhid(name)}})
	request.SetMethod("passphrase.query")
	request.AddValues(queryList)
	resp, err := request.Send()
	err = json.Unmarshal(resp, &passPhrase)

	return passPhrase, err
}
