package phabricator

import (
	"encoding/json"
	"regexp"
)

// GetAllPassPhrase returns the list of all available passPhrases to the given user
// @ TODO Update to avoid request passing
//func GetAllPassPhrase(request *Request) (passPhrases PassPhrase, err error) {
//	queryList := make([]Query, 0)
//	queryList = append(queryList, Query{"string", "needSecrets", "1"})
//	queryList = append(queryList, Query{"string", "limit", "100"})
//	request.AddValues(queryList)
//	request.SetMethod("passphrase.query")
//	resp, err := request.Send()
//	err = json.Unmarshal(resp, &passPhrases)
//
//	return passPhrases, err
//}

// Get the passphrase from the id
// The name should be a monogram
func GetPassPhraseWithId(p *Phabricator, name string) (passPhrase PassPhrase, err error) {
	request := NewRequest(p)
	var idsString []string
	re := regexp.MustCompile("[0-9]+")
	idsString = re.FindAllString(name, -1);
	queryList := make([]Query, 0)
	queryList = append(queryList, Query{"string", "needSecrets", "1"})
	queryList = append(queryList, Query{"array", "ids", idsString})
	request.SetMethod("passphrase.query")
	request.AddValues(queryList)
	resp, err := request.Send(p)
	err = json.Unmarshal(resp, &passPhrase)

	return passPhrase, err
}

// GetPassPhrase returns the passPhrase for the given name
// The name should be the Monogram and the number afterwards. Like K23
// @ TODO Update to avoid request passing
//func GetPassPhrase(request *Request, name string) (passPhrase PassPhrase, err error) {
//	phid, err := GetPhid(request, name)
//	if err != nil {
//		return passPhrase, err
//	}
//	queryList := make([]Query, 0)
//	queryList = append(queryList, Query{"string", "needSecrets", "1"})
//	queryList = append(queryList, Query{"array", "phids", []string{phid.ExtractPhid(name)}})
//	request.SetMethod("passphrase.query")
//	request.AddValues(queryList)
//	resp, err := request.Send()
//	err = json.Unmarshal(resp, &passPhrase)
//
//	return passPhrase, err
//}
