package phabricator

import (
	"encoding/json"
)

// GetServicesFuture Returns an Async FutureAlmanac
func (p *Phabricator) GetServicesFuture() (FutureAlmanac, FutureError) {
	fAlmanac := make(FutureAlmanac)
	fError := make(FutureError)
	request := p.NewRequest()
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

// GetServicesAsync Returns the list of services asynchronously
func (p *Phabricator) GetServicesAsync() (Almanac, error) {
	fAlmanac, fError := p.GetServicesFuture()
	return <-fAlmanac, <-fError
}

// GetDeviceFuture Returns an Async FutureAlmanac
func (p *Phabricator) GetDeviceFuture(hostName string) (FutureAlmanac, FutureError) {
	fd := make(FutureAlmanac)
	ferr := make(FutureError)
	request := p.NewRequest()
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

// GetDeviceAsync Returns the device information asynchronously
func (p *Phabricator) GetDeviceAsync(hostName string) (Almanac, error) {
	fDevice, fError := p.GetDeviceFuture(hostName)
	return <-fDevice, <-fError
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

// GetDevices returns the list of Devices from Almanac
// Backward compatibility
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

// GetDevice returns the specification for the given device
// Backward compatibility
func GetDevice(request *Request, hostName string) (device Almanac, err error) {
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
	resp, err := SendRequest(request)
	err = json.Unmarshal(resp, &device)

	return device, err
}

// GetService returns the specification for the given service
// Backward compatibility
func GetService(request *Request, serviceName string) (service Almanac, err error) {
	queryList := make([]Query, 0)
	attachments := make(map[string]string)
	attachments["properties"] = "1"
	attachments["projects"] = "1"
	attachments["bindings"] = "1"
	constraints := make(map[string][]string)
	nameConstraint := []string{serviceName}
	constraints["names"] = nameConstraint
	queryList = append(queryList, Query{"map", "attachments", attachments})
	queryList = append(queryList, Query{"mapArray", "constraints", constraints})
	queryList = append(queryList, Query{"string", "limit", "100"})
	request.SetMethod("almanac.service.search")
	request.AddValues(queryList)
	resp, err := SendRequest(request)
	err = json.Unmarshal(resp, &service)
	return service, err
}
