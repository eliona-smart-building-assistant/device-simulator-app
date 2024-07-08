package eliona

import (
	confmodel "device-simulator/model/conf"
	"fmt"

	"github.com/eliona-smart-building-assistant/go-eliona/asset"
	"github.com/eliona-smart-building-assistant/go-utils/log"
)

const ClientReference string = "device-simulator"

func UpsertAssetData(generator confmodel.Generator, data any) error {
	log.Debug("Eliona", "upserting data for asset: generator %d", generator.Id)
	d := asset.Data{
		AssetId:         generator.AssetId,
		Data:            data, // todo, I guess
		ClientReference: ClientReference,
	}
	if err := asset.UpsertAssetDataIfAssetExists(d); err != nil {
		return fmt.Errorf("upserting data: %v", err)
	}

	return nil
}
