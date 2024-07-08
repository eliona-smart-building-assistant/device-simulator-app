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
	"device-simulator/appdb"
	confmodel "device-simulator/model/conf"
	"errors"
	"fmt"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

var ErrBadRequest = errors.New("bad request")
var ErrNotFound = errors.New("not found")

func InsertConfig(ctx context.Context, config confmodel.Generator) (confmodel.Generator, error) {
	dbGenerator := toDbGenerator(ctx, config)
	if err := dbGenerator.InsertG(ctx, boil.Infer()); err != nil {
		return confmodel.Generator{}, fmt.Errorf("inserting DB config: %v", err)
	}
	return config, nil
}

func UpsertConfig(ctx context.Context, config confmodel.Generator) (confmodel.Generator, error) {
	dbGenerator := toDbGenerator(ctx, config)
	if err := dbGenerator.UpsertG(ctx, true, []string{"id"}, boil.Blacklist("id"), boil.Infer()); err != nil {
		return confmodel.Generator{}, fmt.Errorf("inserting DB config: %v", err)
	}
	return config, nil
}

func GetGenerator(ctx context.Context, configID int64) (confmodel.Generator, error) {
	dbGenerator, err := appdb.Generators(
		appdb.GeneratorWhere.ID.EQ(configID),
	).OneG(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		return confmodel.Generator{}, ErrNotFound
	}
	if err != nil {
		return confmodel.Generator{}, fmt.Errorf("fetching config from database: %v", err)
	}
	return toAppGenerator(dbGenerator), nil
}

func DeleteGenerator(ctx context.Context, configID int64) error {
	count, err := appdb.Generators(
		appdb.GeneratorWhere.ID.EQ(configID),
	).DeleteAllG(ctx)
	if err != nil {
		return fmt.Errorf("deleting config from database: %v", err)
	}
	if count > 1 {
		return fmt.Errorf("shouldn't happen: deleted more (%v) configs by ID", count)
	}
	if count == 0 {
		return ErrNotFound
	}
	return nil
}

func toDbGenerator(ctx context.Context, appGenerator confmodel.Generator) appdb.Generator {
	return appdb.Generator{
		ID:              int64(appGenerator.Id),
		AssetID:         appGenerator.AssetId,
		Attribute:       appGenerator.Attribute,
		FunctionType:    appGenerator.FunctionType,
		MinValue:        appGenerator.MinValue,
		MaxValue:        appGenerator.MaxValue,
		Frequency:       appGenerator.Frequency,
		IntervalSeconds: appGenerator.IntervalSeconds,
	}
}

func toAppGenerator(dbGenerator *appdb.Generator) confmodel.Generator {
	return confmodel.Generator{
		Id:              int32(dbGenerator.ID),
		AssetId:         dbGenerator.AssetID,
		Attribute:       dbGenerator.Attribute,
		FunctionType:    dbGenerator.FunctionType,
		MinValue:        dbGenerator.MinValue,
		MaxValue:        dbGenerator.MaxValue,
		IntervalSeconds: dbGenerator.IntervalSeconds,
		Frequency:       dbGenerator.Frequency,
	}
}

func GetGenerators(ctx context.Context) ([]confmodel.Generator, error) {
	dbGenerators, err := appdb.Generators().AllG(ctx)
	if err != nil {
		return nil, err
	}
	var appGenerators []confmodel.Generator
	for _, dbGenerator := range dbGenerators {
		appGenerators = append(appGenerators, toAppGenerator(dbGenerator))
	}
	return appGenerators, nil
}
