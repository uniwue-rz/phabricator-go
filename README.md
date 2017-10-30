# Phabricator Go API

This is a simple Phabricator Go Conduit API, covering Almanac and Passphrase.
This API can be used with different automatization tools like and Ansible, Puppet, Chef or Salt to
create a dynamic inventory of the devices, services and their properties. The insecure
data can be saved in Phabricator as property of given device or service and the secure data can be handled using the
Passphrase.

## Installation

To install this packet you should add the repository in your import section
and the run the `goget` command.

```lang=go
import 	phabricator "github.com/uniwue-rz/phabricator-go"
```


## Usage
The following code exmaple is how this API is used:
```lang=go
request := phabricator.NewRequest("https://phabricator-domain/api/", "api-token")
resp, err := phabricator.GetService(&request, "serivce-name")
```


## Supported API

This Conduit API wrapper for go supports following Phabricator apis:
- `almanac`
- `passphrase`
- `phid`

## Extending

You can extend this library to enhance the functionality. To do so
you should create a new file with your `conduit_api.go` and `conduit_api_types.go`. In the `conduit_api` like the existing files you should declare your Methods, that
fetch the data from the server. In this you simply can use the `Request` struct and
add your modification to it. In `coundit_api_types.go` you describe the way the
data should be decoded. See `passphrase.go` and `passphrase_types.go` for hints.

## License
See LICENSE file