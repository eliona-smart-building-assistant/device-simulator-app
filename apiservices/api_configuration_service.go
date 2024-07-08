//  This file is part of the Eliona project.
//  Copyright © 2024 IoTEC AG. All Rights Reserved.
//  ______ _ _
// |  ____| (_)
// | |__  | |_  ___  _ __   __ _
// |  __| | | |/ _ \| '_ \ / _` |
// | |____| | | (_) | | | | (_| |
// |______|_|_|\___/|_| |_|\__,_|
//
//  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING
//  BUT NOT LIMITED  TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
//  NON INFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
//  DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package apiservices

import (
	"context"
	"device-simulator/apiserver"
	"errors"
	"net/http"
)

// ConfigurationAPIService is a service that implements the logic for the ConfigurationAPIServicer
// This service should implement the business logic for every endpoint for the ConfigurationAPI API.
// Include any external packages or services that will be required by this service.
type ConfigurationAPIService struct {
}

// NewConfigurationAPIService creates a default api service
func NewConfigurationAPIService() apiserver.ConfigurationAPIServicer {
	return &ConfigurationAPIService{}
}

// GeneratorsGet - List all generators
func (s *ConfigurationAPIService) GeneratorsGet(ctx context.Context) (apiserver.ImplResponse, error) {
	// TODO - update GeneratorsGet with the required logic for this service method.
	// Add api_configuration_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, []Generator{}) or use other options such as http.Ok ...
	// return apiserver.Response(200, []Generator{}), nil

	return apiserver.Response(http.StatusNotImplemented, nil), errors.New("GeneratorsGet method not implemented")
}

// GeneratorsIdDelete - Delete a generator by ID
func (s *ConfigurationAPIService) GeneratorsIdDelete(ctx context.Context, id int32) (apiserver.ImplResponse, error) {
	// TODO - update GeneratorsIdDelete with the required logic for this service method.
	// Add api_configuration_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return apiserver.Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return apiserver.Response(404, nil),nil

	return apiserver.Response(http.StatusNotImplemented, nil), errors.New("GeneratorsIdDelete method not implemented")
}

// GeneratorsIdGet - Get a generator by ID
func (s *ConfigurationAPIService) GeneratorsIdGet(ctx context.Context, id int32) (apiserver.ImplResponse, error) {
	// TODO - update GeneratorsIdGet with the required logic for this service method.
	// Add api_configuration_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, Generator{}) or use other options such as http.Ok ...
	// return apiserver.Response(200, Generator{}), nil

	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return apiserver.Response(404, nil),nil

	return apiserver.Response(http.StatusNotImplemented, nil), errors.New("GeneratorsIdGet method not implemented")
}

// GeneratorsIdPut - Update a generator by ID
func (s *ConfigurationAPIService) GeneratorsIdPut(ctx context.Context, id int32, generator apiserver.Generator) (apiserver.ImplResponse, error) {
	// TODO - update GeneratorsIdPut with the required logic for this service method.
	// Add api_configuration_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	// return apiserver.Response(200, nil),nil

	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return apiserver.Response(404, nil),nil

	return apiserver.Response(http.StatusNotImplemented, nil), errors.New("GeneratorsIdPut method not implemented")
}

// GeneratorsPost - Create a new generator
func (s *ConfigurationAPIService) GeneratorsPost(ctx context.Context, generator apiserver.Generator) (apiserver.ImplResponse, error) {
	// TODO - update GeneratorsPost with the required logic for this service method.
	// Add api_configuration_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(201, {}) or use other options such as http.Ok ...
	// return apiserver.Response(201, nil),nil

	return apiserver.Response(http.StatusNotImplemented, nil), errors.New("GeneratorsPost method not implemented")
}
