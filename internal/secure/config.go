/*******************************************************************************
 * Copyright 2019 VMWare Inc.
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
 *
 *******************************************************************************/
package secure

import (
	"flag"
	"fmt"

	"github.com/BurntSushi/toml"

	"github.com/edgexfoundry/docker-edgex-mongo/internal/pkg"
)

type secureConfiguration struct {
	pkg.Configuration
	pkg.SecretStoreInfo
}

func LoadConfig() (*pkg.Configuration, error) {
	configFileLocation := flag.String("config", pkg.SecureConfigPath, "configuration file")
	flag.Parse()
	pkg.LoggingClient.Info(fmt.Sprintf("loading the configuration from: %s", *configFileLocation))
	secureConfig := secureConfiguration{}
	_, err := toml.DecodeFile(*configFileLocation, &secureConfig)
	if err != nil {
		return nil, err
	}

	//TODO 1 Connect to Vault
	//TODO 2 Read the credentials
	//TODO 3 If ok ->
	//TODO 4 Update the Configuration.Credentials with what is read from Vault
	//TODO 5 Return the Configuration struct. From that point we should not care if we are working in secure or none-secure envieroment

	return &secureConfig.Configuration, err
}
