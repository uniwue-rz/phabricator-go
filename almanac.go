package phabricator

import (
	"encoding/json"
)

// GetDevices returns the list of Devices from Almanac
// @TODO Update to async if still necessary
func GetDevices(request *Request) (devices Almanac, err error) {
	queryList := make([]Query, 0)
	attachments := make(map[string]string)
	attachments["properties"] = "1"
	attachments["projects"] = "1"
	queryList = append(queryList, Query{"map", "attachments", attachments})
	queryList = append(queryList, Query{"string", "limit", "100"})
	request.SetMethod("almanac.device.search")
	request.AddValues(queryList)
	resp, err := SendRequest(request)
	err = json.Unmarshal(resp, &devices)

	return devices, err
}

// GetServices returns the list of available service

func GetServicesAsync(p *Phabricator) (FutureAlmanac, FutureError) {
	fAlmanac := make(FutureAlmanac)
	fError := make(FutureError)
	request := NewRequest(p)
	queryList := make([]Query, 0)
	attachments := make(map[string]string)
	attachments["properties"] = "1"
	attachments["projects"] = "1"
	attachments["bindings"] = "1"
	queryList = append(queryList, Query{"map", "attachments", attachments})
	queryList = append(queryList, Query{"string", "limit", "100"})
	request.AddValues(queryList)
	request.SetMethod("almanac.service.search")

	go func(request Request) {
		resp, err := request.Send(p)
		var services Almanac
		err = json.Unmarshal(resp, &services)
		fAlmanac <- services
		fError <- err
	}(request)

	return fAlmanac, fError
}

func GetServices(p *Phabricator) (Almanac, error) {
	fAlmanac, fError := GetServicesAsync(p)
	return <- fAlmanac, <- fError
}

// GetDevice gives the specification for the given device
func GetDeviceAsync(p *Phabricator, hostName string) (FutureAlmanac, FutureError) {
	fd := make(FutureAlmanac)
	ferr := make(FutureError)

	request := NewRequest(p)
	queryList := make([]Query, 0)
	attachments := make(map[string]string)
	attachments["properties"] = "1"
	attachments["projects"] = "1"
	constraints := make(map[string][]string)
	nameConstraint := []string{hostName}
	constraints["names"] = nameConstraint
	queryList = append(queryList, Query{"map", "attachments", attachments})
	queryList = append(queryList, Query{"mapArray", "constraints", constraints})
	queryList = append(queryList, Query{"string", "limit", "100"})
	request.SetMethod("almanac.device.search")
	request.AddValues(queryList)

	go func(request Request) {
		resp, err := request.Send(p)
		var device Almanac
		err = json.Unmarshal(resp, &device)
		fd <- device
		ferr <- err
	}(request)

	return fd, ferr
}

func GetDevice(p *Phabricator, hostName string) (Almanac, error) {
	fDevice, fError := GetDeviceAsync(p, hostName)
	return <- fDevice, <- fError
}

// GetService returns the specification for the given service
// @TODO Update to async if still necessary
//func GetService(request *Request, serviceName string) (service Almanac, err error) {
//	queryList := make([]Query, 0)
//	attachments := make(map[string]string)
//	attachments["properties"] = "1"
//	attachments["projects"] = "1"
//	attachments["bindings"] = "1"
//	constraints := make(map[string][]string)
//	nameConstraint := []string{serviceName}
//	constraints["names"] = nameConstraint
//	queryList = append(queryList, Query{"map", "attachments", attachments})
//	queryList = append(queryList, Query{"mapArray", "constraints", constraints})
//	queryList = append(queryList, Query{"string", "limit", "100"})
//	request.SetMethod("almanac.service.search")
//	request.AddValues(queryList)
//	resp, err := request.Send()
//	err = json.Unmarshal(resp, &service)
//
//	return service, err
//}

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
