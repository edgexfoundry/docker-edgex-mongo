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
	"fmt"

	"github.com/BurntSushi/toml"

	"github.com/edgexfoundry/docker-edgex-mongo/internal/pkg"
	secrets "github.com/edgexfoundry/go-mod-secrets/pkg/providers/vault"
	"github.com/edgexfoundry/go-mod-secrets/pkg/token/authtokenloader"
	"github.com/edgexfoundry/go-mod-secrets/pkg/token/fileioperformer"
)

type secureConfiguration struct {
	pkg.Configuration
	SecretStore pkg.SecretStoreInfo
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

	secretClient, err := secrets.NewSecretClient(secrets.SecretConfig{
		Port:                    secureConfig.SecretStore.Port,
		Host:                    secureConfig.SecretStore.Server,
		Path:                    secureConfig.SecretStore.Path,
		Protocol:                "https",
		RootCaCertPath:          secureConfig.SecretStore.CACertPath,
		ServerName:              secureConfig.SecretStore.SNI,
		Authentication:          secrets.AuthenticationInfo{AuthType: pkg.VaultToken, AuthToken: token},
		AdditionalRetryAttempts: secureConfig.SecretStore.AdditionalRetryAttempts,
		RetryWaitPeriod:         secureConfig.SecretStore.RetryWaitPeriod,
	})

	if err != nil {
		pkg.LoggingClient.Error(fmt.Sprintf("fail to connecto secret store: %v", err.Error()))
		return nil, err
	}

	var credentials = make(map[string]pkg.DatabaseInfo)
	for _, dbName := range getDatabaseNames(secureConfig) {
		pkg.LoggingClient.Debug(fmt.Sprintf("reading secrets from '%s/%s' path", secureConfig.SecretStore.Path, dbName))
		secrets, err := secretClient.GetSecrets("/"+dbName, "username", "password")
		if err != nil {
			pkg.LoggingClient.Error(fmt.Sprintf("failed to read secret stores data for '%s/%s' path: %s", secureConfig.SecretStore.Path, dbName, err.Error()))
			return nil, err
		}
		crInfo := pkg.DatabaseInfo{Username: secrets["username"], Password: secrets["password"]}
		credentials[dbName] = crInfo
	}
	secureConfig.UpdateCredentials(credentials)
	pkg.LoggingClient.Debug("Credentials successfully read from Secret Store")
	return &secureConfig.Configuration, err
}

func getDatabaseNames(secureConfig secureConfiguration) []string {
	databases := make([]string, len(secureConfig.Databases))
	i := 0
	for dbName := range secureConfig.Databases {
		databases[i] = dbName
		i++
	}
	return databases
}

func getAccessToken(filename string) (string, error) {
	fileOpener := fileioperformer.NewDefaultFileIoPerformer()
	tokenLoader := authtokenloader.NewAuthTokenLoader(fileOpener)
	token, err := tokenLoader.Load(filename)
	if err != nil {
		return "", err
	}

	return token, nil
}
