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
	DeleteConfigurationById(http.ResponseWriter, *http.Request)
	GetConfigurationById(http.ResponseWriter, *http.Request)
	GetConfigurations(http.ResponseWriter, *http.Request)
	PostConfiguration(http.ResponseWriter, *http.Request)
	PutConfigurationById(http.ResponseWriter, *http.Request)
}

// CustomizationAPIRouter defines the required methods for binding the api requests to a responses for the CustomizationAPI
// The CustomizationAPIRouter implementation should parse necessary information from the http request,
// pass the data to a CustomizationAPIServicer to perform the required actions, then write the service results to the http response.
type CustomizationAPIRouter interface {
	GetDashboardTemplateByName(http.ResponseWriter, *http.Request)
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
	DeleteConfigurationById(context.Context, int64) (ImplResponse, error)
	GetConfigurationById(context.Context, int64) (ImplResponse, error)
	GetConfigurations(context.Context) (ImplResponse, error)
	PostConfiguration(context.Context, Configuration) (ImplResponse, error)
	PutConfigurationById(context.Context, int64, Configuration) (ImplResponse, error)
}

// CustomizationAPIServicer defines the api actions for the CustomizationAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type CustomizationAPIServicer interface {
	GetDashboardTemplateByName(context.Context, string, string) (ImplResponse, error)
}

// VersionAPIServicer defines the api actions for the VersionAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type VersionAPIServicer interface {
	GetOpenAPI(context.Context) (ImplResponse, error)
	GetVersion(context.Context) (ImplResponse, error)
}
