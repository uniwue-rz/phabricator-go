package phabricator

// PassPhrase is the wrapper for the PassPhrase API data
type PassPhrase struct {
	Result struct {
		Data   map[string]PassPhraseObject `json:"data"`
		Cursor Cursor                      `json:"cursor"`
	} `json:"result"`
	ErrorCode string `json:"error_code"`
	ErrorInfo string `json:"error_info"`
}

// PassPhraseObject is the main construct that contains every passphrase element
type PassPhraseObject struct {
	Id             string             `json:"id"`
	PHID           string             `json:"phid"`
	PassPhraseType string             `json:"type"`
	Name           string             `json:"name"`
	Description    string             `json:"description"`
	URI            string             `json:"uri"`
	Monogram       string             `json:"monogram"`
	Username       string             `json:"username"`
	Material       PassPhraseMaterial `json:"material"`
}

// PassPhraseMaterial contains the secure data that is received from the Phabricator Server
type PassPhraseMaterial struct {
	PrivateKey string `json:"privateKey"`
	Password   string `json:"password"`
}
