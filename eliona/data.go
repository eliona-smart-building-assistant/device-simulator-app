package eliona

import (
	confmodel "device-simulator/v2/model/conf"
	"fmt"

	api "github.com/eliona-smart-building-assistant/go-eliona-api-client/v3"
	"github.com/eliona-smart-building-assistant/go-eliona/v2/asset"
	"github.com/eliona-smart-building-assistant/go-utils/log"
)

const ClientReference string = "device-simulator"

func UpsertAssetData(generator confmodel.Generator, data map[string]any) error {
	log.Debug("Eliona", "upserting data for asset: generator %d", generator.Id)

	cr := ClientReference // Needed to take a reference for NewNullableString call
	d := api.Data{
		AssetId:         generator.AssetId,
		Data:            data,
		ClientReference: *api.NewNullableString(&cr),
		Subtype:         api.DataSubtype(generator.Subtype),
	}
	if err := asset.UpsertDataIfAssetExists(generator.ApiEndpoint, generator.ApiKey, d); err != nil {
		return fmt.Errorf("upserting data: %v", err)
	}

	return nil
}
