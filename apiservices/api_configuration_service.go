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
	"device-simulator/conf"
	"device-simulator/eliona"
	confmodel "device-simulator/model/conf"
	"errors"
	"fmt"
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
	appGenerators, err := conf.GetGenerators(ctx)
	if err != nil {
		return apiserver.ImplResponse{Code: http.StatusInternalServerError}, err
	}
	generators := make([]apiserver.Generator, 0, len(appGenerators))
	for _, appGenerator := range appGenerators {
		generators = append(generators, toAPIGenerator(appGenerator))
	}
	return apiserver.Response(http.StatusOK, generators), nil
}

// GeneratorsIdDelete - Delete a generator by ID
func (s *ConfigurationAPIService) GeneratorsIdDelete(ctx context.Context, id int32) (apiserver.ImplResponse, error) {
	err := conf.DeleteGenerator(ctx, int64(id))
	if errors.Is(err, conf.ErrNotFound) {
		return apiserver.ImplResponse{Code: http.StatusNotFound}, nil
	}
	if err != nil {
		return apiserver.ImplResponse{Code: http.StatusInternalServerError}, err
	}
	return apiserver.ImplResponse{Code: http.StatusNoContent}, nil
}

// GeneratorsIdGet - Get a generator by ID
func (s *ConfigurationAPIService) GeneratorsIdGet(ctx context.Context, id int32) (apiserver.ImplResponse, error) {
	generator, err := conf.GetGenerator(ctx, int64(id))
	if errors.Is(err, conf.ErrNotFound) {
		return apiserver.ImplResponse{Code: http.StatusNotFound}, nil
	}
	if err != nil {
		return apiserver.ImplResponse{Code: http.StatusInternalServerError}, err
	}
	return apiserver.Response(http.StatusOK, toAPIGenerator(generator)), nil
}

// GeneratorsIdPut - Update a generator by ID
func (s *ConfigurationAPIService) GeneratorsIdPut(ctx context.Context, id int32, generator apiserver.Generator) (apiserver.ImplResponse, error) {
	generator.Id = id
	appGenerator, err := toAppGenerator(ctx, generator)
	if err != nil {
		return apiserver.ImplResponse{Code: http.StatusInternalServerError}, err
	}
	upsertedGenerator, err := conf.UpsertGenerator(ctx, appGenerator)
	if err != nil {
		return apiserver.ImplResponse{Code: http.StatusInternalServerError}, err
	}
	return apiserver.Response(http.StatusCreated, toAPIGenerator(upsertedGenerator)), nil
}

// GeneratorsPost - Create a new generator
func (s *ConfigurationAPIService) GeneratorsPost(ctx context.Context, generator apiserver.Generator) (apiserver.ImplResponse, error) {
	appGenerator, err := toAppGenerator(ctx, generator)
	if err != nil {
		return apiserver.ImplResponse{Code: http.StatusInternalServerError}, err
	}
	insertedGenerator, err := conf.InsertGenerator(ctx, appGenerator)
	if err != nil {
		return apiserver.ImplResponse{Code: http.StatusInternalServerError}, err
	}
	return apiserver.Response(http.StatusCreated, toAPIGenerator(insertedGenerator)), nil
}

func toAPIGenerator(appGenerator confmodel.Generator) apiserver.Generator {
	return apiserver.Generator{
		Id:              int32(appGenerator.Id),
		AssetId:         appGenerator.AssetId,
		Attribute:       appGenerator.Attribute,
		Subtype:         appGenerator.Subtype,
		FunctionType:    appGenerator.FunctionType,
		MinValue:        appGenerator.MinValue,
		MaxValue:        appGenerator.MaxValue,
		Integer:         appGenerator.Integer,
		IntervalSeconds: appGenerator.IntervalSeconds,
		Frequency:       appGenerator.Frequency,
	}
}

func toAppGenerator(ctx context.Context, apiGenerator apiserver.Generator) (confmodel.Generator, error) {
	assetTypeName, err := eliona.GetAssetType(ctx, apiGenerator.AssetId)
	if err != nil {
		return confmodel.Generator{}, fmt.Errorf("getting asset type name: %v", err)
	}
	return confmodel.Generator{
		Id:              int32(apiGenerator.Id),
		AssetId:         apiGenerator.AssetId,
		Attribute:       apiGenerator.Attribute,
		Subtype:         apiGenerator.Subtype,
		AssetType:       assetTypeName,
		FunctionType:    apiGenerator.FunctionType,
		MinValue:        apiGenerator.MinValue,
		MaxValue:        apiGenerator.MaxValue,
		Integer:         apiGenerator.Integer,
		IntervalSeconds: apiGenerator.IntervalSeconds,
		Frequency:       apiGenerator.Frequency,
	}, nil
}
