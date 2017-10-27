package phabricator

import (
	"encoding/json"
)

// GetDevices returns the list of Devices from Almanac
func GetDevices(apiUrl string, token string) (devices Almanac, err error){
	queryList := []Query{}
	queryList = AddToken(queryList, token)
	attachments := make(map[string]string)
	attachments["properties"] = "1"
	attachments["projects"] = "1"
	queryList = append(queryList, Query{"map", "attachments", attachments})
	queryList = append(queryList, Query{"string", "limit", "100"})
	resp, err := SendApiQuery(apiUrl, "almanac.device.search", queryList)

	json.Unmarshal(resp, &devices)

	return devices, err
}

// GetServices returns the list of available service
func GetServices(apiUrl string, token string) (services Almanac, err error){
	queryList := []Query{}
	queryList = AddToken(queryList, token)
	attachments := make(map[string]string)
	attachments["properties"] = "1"
	attachments["projects"] = "1"
	attachments["bindings"] = "1"
	queryList = append(queryList, Query{"map", "attachments", attachments})
	queryList = append(queryList, Query{"string", "limit", "100"})
	resp, err := SendApiQuery(apiUrl, "almanac.service.search", queryList)

	json.Unmarshal(resp, &services)

	return services, err
}

// GetDevice returns the specification for the given device
func GetDevice(apiUrl string, token string, hostName string) (device Almanac, err error){
	queryList := []Query{}
	queryList = AddToken(queryList, token)
	attachments := make(map[string]string)
	attachments["properties"] = "1"
	attachments["projects"] = "1"
	constraints := make(map[string][]string)
	nameConstraint := []string{hostName}
	constraints["names"] = nameConstraint
	queryList = append(queryList, Query{"map", "attachments", attachments})
	queryList = append(queryList, Query{"mapArray", "constraints", constraints})
	resp, err := SendApiQuery(apiUrl, "almanac.device.search", queryList)
	json.Unmarshal(resp, &device)

	return device, err
}

// GetService returns the specification for the given service
func GetService(apiUrl string, token string, serviceName string) (service Almanac, err error) {
	queryList := []Query{}
	queryList = AddToken(queryList, token)
	attachments := make(map[string]string)
	attachments["properties"] = "1"
	attachments["projects"] = "1"
	attachments["bindings"] = "1"
	constraints := make(map[string][]string)
	nameConstraint := []string{serviceName}
	constraints["names"] = nameConstraint
	queryList = append(queryList, Query{"map", "attachments", attachments})
	queryList = append(queryList, Query{"mapArray", "constraints", constraints})
	resp, err := SendApiQuery(apiUrl, "almanac.service.search", queryList)

	json.Unmarshal(resp, &service)

	return service, err
}

// GetProperties returns the properties for the given device or service
func GetProperties(apiUrl string, token string, name string, inventoryType string) (res map[string]string, err error){

	return res, err
}
