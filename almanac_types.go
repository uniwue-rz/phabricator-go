package phabricator

// AlmanacDevices contains all the data related to a list of Almanac Devices
type Almanac struct {
	Result struct {
		Data   []Device    `json:"data"`
		Cursor Cursor      `json:"cursor"`
		Maps   interface{} `json:"maps"`
		Query  struct {
			QueryKey string `json:"queryKey"`
		} `json:"query"`
	}
	ErrorCode string `json:"error_code"`
	ErrorInfo string `json:"error_info"`
}

// Attachments are handled using this
type Attachment struct {
	Properties struct {
		Properties []Property `json:"properties"`
	} `json:"properties"`
	Projects struct {
		ProjectPHIDs []string `json:"projectPHIDs"`
	} `json:"projects"`
	Bindings struct {
		Bindings []Binding `json:"bindings"`
	} `json:"bindings"`
}

// Binding describes the service given bindings
type Binding struct {
	Id         int        `json:"id"`
	PHID       string     `json:"phid"`
	Properties []Property `json:"properties"`
	Interface  Interface  `json:"interface"`
	Disabled   bool       `json:"disabled"`
}

// Interface describes the interface in a given binding
type Interface struct {
	Id      int           `json:"id"`
	PHID    string        `json:"phid"`
	Address string        `json:"address"`
	Port    int           `json:"port"`
	Device  ServiceDevice `json:"device"`
	Network Network       `json:"network"`
}

// Network describes the network property of a given service
type Network struct {
	Id   int    `json:"id"`
	PHID string `json:"phid"`
	Name string `json:"name"`
}

// Field describes the field in data result
type Field struct {
	Name         string `json:"name"`
	DateCreated  int    `json:"dateCreated"`
	DateModified int    `json:"dateModified"`
	Policy       Policy `json:"policy"`
	ServiceType  string `json:"serviceType"`
}

// Policy describes the policies in data results
type Policy struct {
	View string `json:"view"`
	Edit string `json:"edit"`
}

// Property describes every given property
type Property struct {
	Key     string `json:"key"`
	Value   string `json:"value"`
	BuiltIn bool   `json:"builtin"`
}

// ServiceDevice is special kind of device that is delivered by the service api
type ServiceDevice struct {
	Id         int        `json:"id"`
	PHID       string     `json:"phid"`
	Name       string     `json:"name"`
	Properties []Property `json:"properties"`
}

// Device contains the structure that should be used for devices
type Device struct {
	Id          int        `json:"id"`
	DevType     string     `json:"type"`
	Phid        string     `json:"phid"`
	Fields      Field      `json:"fields"`
	Attachments Attachment `json:"attachments"`
}

// Service contains the structure that should be used for the services
type Service struct {
	Name       string
	Properties map[string]string
	Devices    []Device
}
