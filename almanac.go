package phabricator

import (
	"encoding/json"
)

// GetDevices returns the list of Devices from Almanac
func GetDevices(request *Request) (devices Almanac, err error) {
	queryList := []Query{}
	attachments := make(map[string]string)
	attachments["properties"] = "1"
	attachments["projects"] = "1"
	queryList = append(queryList, Query{"map", "attachments", attachments})
	queryList = append(queryList, Query{"string", "limit", "100"})
	request.SetMethod("almanac.device.search")
	request.AddValues(queryList)
	resp, err := SendRequest(request)
	json.Unmarshal(resp, &devices)

	return devices, err
}

// GetServices returns the list of available service
func GetServices(request *Request) (services Almanac, err error) {
	queryList := []Query{}
	attachments := make(map[string]string)
	attachments["properties"] = "1"
	attachments["projects"] = "1"
	attachments["bindings"] = "1"
	queryList = append(queryList, Query{"map", "attachments", attachments})
	queryList = append(queryList, Query{"string", "limit", "100"})
	request.AddValues(queryList)
	request.SetMethod("almanac.service.search")
	resp, err := request.Send()
	json.Unmarshal(resp, &services)

	return services, err
}

// GetDevice returns the specification for the given device
func GetDevice(request *Request, hostName string) (device Almanac, err error) {
	queryList := []Query{}
	attachments := make(map[string]string)
	attachments["properties"] = "1"
	attachments["projects"] = "1"
	constraints := make(map[string][]string)
	nameConstraint := []string{hostName}
	constraints["names"] = nameConstraint
	queryList = append(queryList, Query{"map", "attachments", attachments})
	queryList = append(queryList, Query{"mapArray", "constraints", constraints})
	request.SetMethod("almanac.device.search")
	request.AddValues(queryList)
	resp, err := request.Send()
	json.Unmarshal(resp, &device)

	return device, err
}

// GetService returns the specification for the given service
func GetService(request *Request, serviceName string) (service Almanac, err error) {
	queryList := []Query{}
	attachments := make(map[string]string)
	attachments["properties"] = "1"
	attachments["projects"] = "1"
	attachments["bindings"] = "1"
	constraints := make(map[string][]string)
	nameConstraint := []string{serviceName}
	constraints["names"] = nameConstraint
	queryList = append(queryList, Query{"map", "attachments", attachments})
	queryList = append(queryList, Query{"mapArray", "constraints", constraints})
	request.SetMethod("almanac.service.search")
	request.AddValues(queryList)
	resp, err := request.Send()
	json.Unmarshal(resp, &service)

	return service, err
}

// GetProperties Returns the Properties for the result of Almanac Query
func (almanac *Almanac) GetProperties() (properties []Property) {
	for _, v := range almanac.Result.Data {
		properties = append(properties, v.Attachments.Properties.Properties...)
	}

	return properties
}

// GetBindings Returns the bindings for the result of Almanac Query
func (almanac *Almanac) GetBindings() (bindings []Binding) {
	for _, v := range almanac.Result.Data {
		bindings = append(bindings, v.Attachments.Bindings.Bindings...)
	}

	return bindings
}
