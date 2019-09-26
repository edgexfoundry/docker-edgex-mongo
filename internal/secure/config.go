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
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/BurntSushi/toml"

	"github.com/edgexfoundry/docker-edgex-mongo/internal"
	"github.com/edgexfoundry/docker-edgex-mongo/internal/pkg"
	secrets "github.com/edgexfoundry/go-mod-secrets/pkg/providers/vault"
)

type secureConfiguration struct {
	pkg.Configuration
	SecretStore pkg.SecretStoreInfo
}

type auth struct {
	Token string `json:"root_token"`
}

func LoadConfig() (*pkg.Configuration, error) {
	pkg.LoggingClient.Info("loading configuration considering the secret store")
	configFileLocation := pkg.DefineConfigFileLocation()
	secureConfig := secureConfiguration{}
	_, err := toml.DecodeFile(configFileLocation, &secureConfig)
	if err != nil {
		return nil, err
	}

	token, err := getAccessToken(secureConfig.SecretStore.TokenPath)
	if err != nil {
		return nil, err
	}

	var credentials = make(map[string]pkg.CredentialsInfo)
	for _, dbName := range getDatabaseNames() {
		searchPath := fmt.Sprintf("%s/%s", secureConfig.SecretStore.DBStem, dbName)
		pkg.LoggingClient.Debug(fmt.Sprintf("reading secrets from '%s' path", searchPath))
		secretClient, err := secrets.NewSecretClient(secrets.SecretConfig{
			Port:           secureConfig.SecretStore.Port,
			Host:           secureConfig.SecretStore.Server,
			Path:           searchPath,
			Protocol:       "https",
			RootCaCert:     secureConfig.SecretStore.CACertPath,
			ServerName:     secureConfig.SecretStore.SNI,
			Authentication: secrets.AuthenticationInfo{AuthType: pkg.VaultToken, AuthToken: token},
		})

		if err != nil {
			pkg.LoggingClient.Error(fmt.Sprintf("fail to connecto secret store: %v", err.Error()))
			return nil, err
		}
		secrets, err := secretClient.GetSecrets("username", "password")
		if err != nil {
			pkg.LoggingClient.Error(fmt.Sprintf("failed to read secret stores data for '%s' path: %s", searchPath, err.Error()))
			return nil, err
		}
		crInfo := pkg.CredentialsInfo{Username: secrets["username"], Password: secrets["password"]}
		credentials[dbName] = crInfo
	}
	secureConfig.UpdateCredentials(credentials)
	pkg.LoggingClient.Debug("Credentials successfully read from Secret Store")
	return &secureConfig.Configuration, err
}

func getDatabaseNames() []string {
	databases := make([]string, len(internal.DatabaseCollectionsMap))
	i := 0
	for db := range internal.DatabaseCollectionsMap {
		databases[i] = db
		i++
	}
	return databases
}

func getAccessToken(filename string) (string, error) {
	a := auth{}
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		return a.Token, err
	}

	err = json.Unmarshal(raw, &a)
	return a.Token, err
}
