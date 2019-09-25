/*******************************************************************************
 * Copyright 2019 VMWare.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/

package main

import (
	"fmt"
	"os"

	"github.com/edgexfoundry/go-mod-core-contracts/clients/logger"
	"github.com/edgexfoundry/go-mod-core-contracts/models"

	"github.com/edgexfoundry/docker-edgex-mongo/internal"
	"github.com/edgexfoundry/docker-edgex-mongo/internal/pkg"
	"github.com/edgexfoundry/docker-edgex-mongo/internal/secure"
	"github.com/edgexfoundry/docker-edgex-mongo/internal/unsecure"
)

func main() {
	pkg.LoggingClient = logger.NewClient(pkg.EdgexMongoServiceKey, false, "", models.InfoLog)
	pkg.LoggingClient.Info(fmt.Sprintf("starting %s process ...", pkg.EdgexMongoServiceKey))

	var err error
	var config *pkg.Configuration

	switch env := os.Getenv(pkg.SecretStore); {
	case env == "false":
		config, err = unsecure.LoadConfig()
	default:
		config, err = secure.LoadConfig()
	}
	exitIfError(err)

	pkg.LoggingClient = logger.NewClient(pkg.EdgexMongoServiceKey, false, "", config.Writable.LogLevel)

	dbInitClient := internal.DBInitClient{
		Configuration: config,
	}

	err = dbInitClient.PopulateDatabase()
	exitIfError(err)
}

func exitIfError(err error) {
	if err != nil {
		pkg.LoggingClient.Error(fmt.Sprintf("%s failed because of: %s", pkg.EdgexMongoServiceKey, err.Error()))
		os.Exit(1)
	}
}
