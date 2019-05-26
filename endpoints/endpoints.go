package main

import (
	"fmt"

	endpointdiscovery "github.com/eelcoh/endpoint-discovery"
)

func add(manifest endpointdiscovery.Manifest, hostname string, port int) {
	for _, api := range manifest.APIs {
		for _, endpoint := range api.Endpoints {
			endpointInstance := makeEndpointInstance(api, endpoint, hostname, port)
			apiEndpoints[api.API] = append(apiEndpoints[api.API], endpointInstance)
		}
	}
}

func makeEndpointInstance(api *endpointdiscovery.API, endpoint *endpointdiscovery.Endpoint, hostname string, port int) endpointdiscovery.EndpointInstance {
	endpointInstance := endpointdiscovery.EndpointInstance{
		api.Occurence,
		api.Version,
		hostname,
		port,
		endpoint.Method,
		endpoint.Path,
	}
	return endpointInstance
}

func find(api, occurence, version string) []endpointdiscovery.EndpointInstance {
	endpoints := make([]endpointdiscovery.EndpointInstance, 0)

	for _, e := range apiEndpoints[api] {
		isCompatible := compatible(version, e.Version)
		if (e.Occurence == occurence) && isCompatible {
			endpoints = append(endpoints, e)
		}
	}

	return endpoints
}

func compatible(version1, version2 string) bool {
	return (version1 == version2)
}

func del(api, occurence, version, method, path string) {

	endpoints := apiEndpoints[api]

	for i, e := range endpoints {
		if (e.Occurence == occurence) &&
			(e.Version == version) &&
			(e.Method == method) &&
			(e.Path == path) {
			endpoints = append(endpoints[:i], endpoints[i+1:]...)
		}
	}
	apiEndpoints[api] = endpoints
	fmt.Println(endpoints)

}
