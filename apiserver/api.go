/*
 * Device Simulator app API
 *
 * API to access and configure the Device Simulator app
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package apiserver

import (
	"context"
	"net/http"
)

// ConfigurationAPIRouter defines the required methods for binding the api requests to a responses for the ConfigurationAPI
// The ConfigurationAPIRouter implementation should parse necessary information from the http request,
// pass the data to a ConfigurationAPIServicer to perform the required actions, then write the service results to the http response.
type ConfigurationAPIRouter interface {
	GeneratorsGet(http.ResponseWriter, *http.Request)
	GeneratorsIdDelete(http.ResponseWriter, *http.Request)
	GeneratorsIdGet(http.ResponseWriter, *http.Request)
	GeneratorsIdPut(http.ResponseWriter, *http.Request)
	GeneratorsPost(http.ResponseWriter, *http.Request)
}

// VersionAPIRouter defines the required methods for binding the api requests to a responses for the VersionAPI
// The VersionAPIRouter implementation should parse necessary information from the http request,
// pass the data to a VersionAPIServicer to perform the required actions, then write the service results to the http response.
type VersionAPIRouter interface {
	GetOpenAPI(http.ResponseWriter, *http.Request)
	GetVersion(http.ResponseWriter, *http.Request)
}

// ConfigurationAPIServicer defines the api actions for the ConfigurationAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type ConfigurationAPIServicer interface {
	GeneratorsGet(context.Context) (ImplResponse, error)
	GeneratorsIdDelete(context.Context, int32) (ImplResponse, error)
	GeneratorsIdGet(context.Context, int32) (ImplResponse, error)
	GeneratorsIdPut(context.Context, int32, Generator) (ImplResponse, error)
	GeneratorsPost(context.Context, Generator) (ImplResponse, error)
}

// VersionAPIServicer defines the api actions for the VersionAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type VersionAPIServicer interface {
	GetOpenAPI(context.Context) (ImplResponse, error)
	GetVersion(context.Context) (ImplResponse, error)
}
