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

package main

import (
	"context"
	"device-simulator/v2/apiserver"
	"device-simulator/v2/apiservices"
	"device-simulator/v2/conf"
	"device-simulator/v2/eliona"
	confmodel "device-simulator/v2/model/conf"
	"net/http"
	"sync"
	"time"

	"github.com/eliona-smart-building-assistant/go-eliona/v2/app"
	"github.com/eliona-smart-building-assistant/go-eliona/v2/asset"
	"github.com/eliona-smart-building-assistant/go-eliona/v2/dashboard"
	"github.com/eliona-smart-building-assistant/go-eliona/v2/frontend"
	"github.com/eliona-smart-building-assistant/go-utils/common"
	"github.com/eliona-smart-building-assistant/go-utils/db"
	utilshttp "github.com/eliona-smart-building-assistant/go-utils/http"
	"github.com/eliona-smart-building-assistant/go-utils/log"
)

func initialization() {
	ctx := context.Background()

	// Necessary to close used init resources
	conn := db.NewInitConnectionWithContextAndApplicationName(ctx, app.AppName())
	defer conn.Close(ctx)

	apiEndpoint := common.Getenv("API_ENDPOINT", "")
	apiToken := common.Getenv("API_TOKEN", "")

	// Init the app before the first run.
	app.Init(
		apiEndpoint,
		apiToken,
		conn, app.AppName(),
		app.ExecSqlFile("conf/init.sql"),
		asset.InitAssetTypeFiles(apiEndpoint, apiToken, "resources/asset-types/*.json"),
		dashboard.InitWidgetTypeFiles(apiEndpoint, apiToken, "resources/widget-types/*.json"),
	)
	app.Patch(apiEndpoint, apiToken, conn, app.AppName(), "010100",
		app.ExecSqlFile("conf/010100.sql"),
	)
	app.Patch(apiEndpoint, apiToken, conn, app.AppName(), "020000",
		app.ExecSqlFile("conf/020000.sql"),
	)
}

var once sync.Once

func collectData() {
	generators, err := conf.GetGenerators(context.Background())
	if err != nil {
		log.Fatal("conf", "Couldn't read generators from DB: %v", err)
		return
	}
	if len(generators) == 0 {
		once.Do(func() {
			log.Info("conf", "No generators in DB. Please configure the app in Eliona.")
		})
		return
	}

	for _, generator := range generators {
		common.RunOnceWithParam(func(generator confmodel.Generator) {
			log.Info("main", "Collecting %d started.", generator.Id)
			if err := generateData(generator); err != nil {
				return // Error is handled in the method itself.
			}
			log.Info("main", "Collecting %d finished.", generator.Id)

			time.Sleep(time.Second * time.Duration(generator.IntervalSeconds))
		}, generator, generator.Id)
	}
}

func generateData(generator confmodel.Generator) error {
	value := generator.Generate()
	if err := eliona.UpsertAssetData(generator, value); err != nil {
		return err
	}
	return nil
}

// listenApi starts the API server and listen for requests
func listenApi() {
	err := http.ListenAndServe(":"+common.Getenv("API_SERVER_PORT", "3000"),
		frontend.NewEnvironmentHandler(
			utilshttp.NewCORSEnabledHandler(
				apiserver.NewRouter(
					apiserver.NewConfigurationAPIController(apiservices.NewConfigurationAPIService()),
					apiserver.NewVersionAPIController(apiservices.NewVersionAPIService()),
				))))
	log.Fatal("main", "API server: %v", err)
}
