package eliona

import (
	confmodel "device-simulator/model/conf"
	"errors"
	"fmt"

	api "github.com/eliona-smart-building-assistant/go-eliona-api-client/v3"
	"github.com/eliona-smart-building-assistant/go-eliona/v2/asset"
	"github.com/eliona-smart-building-assistant/go-utils/common"
	"github.com/eliona-smart-building-assistant/go-utils/log"
)

const ClientReference string = "device-simulator"

var (
	apiToken    string
	apiEndpoint string
)

var (
	ErrNoAPIEndpoint = errors.New("no API endpoint defined")
	ErrNoAPIToken    = errors.New("no API token defined")
)

func init() {
	apiToken = common.Getenv("API_TOKEN", "")
	apiEndpoint = common.Getenv("API_ENDPOINT", "")
}

func UpsertAssetData(generator confmodel.Generator, data map[string]any) error {
	log.Debug("Eliona", "upserting data for asset: generator %d", generator.Id)

	if apiEndpoint == "" {
		return ErrNoAPIEndpoint
	}

	if apiToken == "" {
		return ErrNoAPIToken
	}

	cr := ClientReference // Needed to take a reference for NewNullableString call
	d := api.Data{
		AssetId:         generator.AssetId,
		Data:            data,
		ClientReference: *api.NewNullableString(&cr),
		Subtype:         api.DataSubtype(generator.Subtype),
		AssetTypeName:   *api.NewNullableString(&generator.AssetType),
	}
	if err := asset.UpsertDataIfAssetExists(apiEndpoint, apiToken, d); err != nil {
		return fmt.Errorf("upserting data: %v", err)
	}

	return nil
}
