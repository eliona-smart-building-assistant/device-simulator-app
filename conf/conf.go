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

package conf

import (
	"context"
	"database/sql"
	"device-simulator/v2/appdb"
	confmodel "device-simulator/v2/model/conf"
	"errors"
	"fmt"
	"github.com/eliona-smart-building-assistant/go-eliona/v2/app"
	"github.com/eliona-smart-building-assistant/go-eliona/v2/client"
	"github.com/eliona-smart-building-assistant/go-eliona/v2/frontend"
	"github.com/google/uuid"
	"sync"

	"github.com/aarondl/sqlboiler/v4/boil"
)

var ErrBadRequest = errors.New("bad request")
var ErrNotFound = errors.New("not found")

func InsertGenerator(ctx context.Context, generator confmodel.Generator) (confmodel.Generator, error) {
	dbGenerator := toDbGenerator(generator)
	if err := dbGenerator.InsertG(ctx, boil.Infer()); err != nil {
		return confmodel.Generator{}, fmt.Errorf("inserting DB generator: %v", err)
	}
	return generator, nil
}

func UpsertGenerator(ctx context.Context, generator confmodel.Generator) (confmodel.Generator, error) {
	dbGenerator := toDbGenerator(generator)
	if err := dbGenerator.UpsertG(ctx, true, []string{"id"}, boil.Blacklist("id"), boil.Infer()); err != nil {
		return confmodel.Generator{}, fmt.Errorf("inserting DB generator: %v", err)
	}
	return generator, nil
}

func GetGenerator(ctx context.Context, generatorID int64) (confmodel.Generator, error) {
	dbGenerator, err := appdb.Generators(
		appdb.GeneratorWhere.ID.EQ(generatorID),
	).OneG(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		return confmodel.Generator{}, ErrNotFound
	}
	if err != nil {
		return confmodel.Generator{}, fmt.Errorf("fetching generator from database: %v", err)
	}
	return toAppGenerator(dbGenerator), nil
}

func DeleteGenerator(ctx context.Context, generatorID int64) error {
	count, err := appdb.Generators(
		appdb.GeneratorWhere.ID.EQ(generatorID),
	).DeleteAllG(ctx)
	if err != nil {
		return fmt.Errorf("deleting generator from database: %v", err)
	}
	if count > 1 {
		return fmt.Errorf("shouldn't happen: deleted more (%v) generators by ID", count)
	}
	if count == 0 {
		return ErrNotFound
	}
	return nil
}

func toDbGenerator(appGenerator confmodel.Generator) appdb.Generator {
	return appdb.Generator{
		ID:              int64(appGenerator.Id),
		TenantID:        appGenerator.ElionaTenantId,
		AssetID:         appGenerator.AssetId,
		Attribute:       appGenerator.Attribute,
		Subtype:         appGenerator.Subtype,
		FunctionType:    appGenerator.FunctionType,
		AssetType:       appGenerator.AssetType,
		MinValue:        appGenerator.MinValue,
		MaxValue:        appGenerator.MaxValue,
		Integer:         appGenerator.Integer,
		Frequency:       appGenerator.Frequency,
		IntervalSeconds: appGenerator.IntervalSeconds,
	}
}

func toAppGenerator(dbGenerator *appdb.Generator) confmodel.Generator {
	tenantUuid := uuid.MustParse(dbGenerator.TenantID)
	apiKey, err := GetApiKey(tenantUuid)
	if err != nil {
		return confmodel.Generator{}
	}

	return confmodel.Generator{
		Id:              int32(dbGenerator.ID),
		AssetId:         dbGenerator.AssetID,
		Attribute:       dbGenerator.Attribute,
		Subtype:         dbGenerator.Subtype,
		FunctionType:    dbGenerator.FunctionType,
		AssetType:       dbGenerator.AssetType,
		MinValue:        dbGenerator.MinValue,
		MaxValue:        dbGenerator.MaxValue,
		Integer:         dbGenerator.Integer,
		IntervalSeconds: dbGenerator.IntervalSeconds,
		Frequency:       dbGenerator.Frequency,
		StartTime:       dbGenerator.InitializedAt,

		ElionaTenantId: dbGenerator.TenantID,
		ApiEndpoint:    client.ApiEndpointString(),
		ApiKey:         apiKey,
	}
}

func GetGenerators(ctx context.Context) ([]confmodel.Generator, error) {
	dbGenerators, err := appdb.Generators().AllG(ctx)
	if err != nil {
		return nil, err
	}
	var appGenerators []confmodel.Generator
	ResetApiKeyCache()
	for _, dbGenerator := range dbGenerators {
		appGenerators = append(appGenerators, toAppGenerator(dbGenerator))
	}
	return appGenerators, nil
}

func ParseTenantIdFromEnv(ctx context.Context) (uuid.UUID, error) {
	env := frontend.GetEnvironment(ctx)
	if env == nil {
		return uuid.UUID{}, fmt.Errorf("missing environment JWT")
	}
	parsed, err := uuid.Parse(env.TenantId)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("tenant isn't a valid UUID: %s", env.TenantId)
	}
	return parsed, err
}

// --- cache (tenantUuid -> apiKey) ---

var apiKeyCache = struct {
	mu sync.RWMutex
	m  map[string]string
}{
	m: make(map[string]string),
}

// ResetApiKeyCache clears the whole cache.
func ResetApiKeyCache() {
	apiKeyCache.mu.Lock()
	defer apiKeyCache.mu.Unlock()
	apiKeyCache.m = make(map[string]string)
}

// GetApiKey is the cached wrapper (same signature as your call-site).
func GetApiKey(tenantUuid uuid.UUID) (string, error) {
	// Fast path: cache hit
	apiKeyCache.mu.RLock()
	if v, ok := apiKeyCache.m[tenantUuid.String()]; ok {
		apiKeyCache.mu.RUnlock()
		return v, nil
	}
	apiKeyCache.mu.RUnlock()

	// Cache miss: fetch using the uncached implementation
	key, err := app.GetApiKey("device-simulator", GetDB(), tenantUuid)
	if err != nil {
		return "", err // don't cache failures
	}

	// Store in cache
	apiKeyCache.mu.Lock()
	apiKeyCache.m[tenantUuid.String()] = key
	apiKeyCache.mu.Unlock()

	return key, nil
}
