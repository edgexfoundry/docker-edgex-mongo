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
package pkg

import (
	"fmt"
	"github.com/edgexfoundry/go-mod-secrets/pkg/providers/vault"
	"net/url"
)

type Configuration struct {
	Service   ServiceInfo
	Writable  WritableInfo
	Mongo     MongoInfo
	Databases map[string]DatabaseInfo
}

func (c *Configuration) UpdateCredentials(databases map[string]DatabaseInfo) {
	c.Databases = databases
}

type ServiceInfo struct {
	// BootTimeout indicates, in milliseconds, how long the service will retry connecting to mongo database
	// before giving up. Default is 30,000.
	BootTimeout int
}

type MongoInfo struct {
	Host    string
	Port    int
	Timeout int
}

type DatabaseInfo struct {
	Username string
	Password string
}

type WritableInfo struct {
	LogLevel       string
	RequestTimeout int
}

type SecretStoreInfo struct {
	vault.SecretConfig
	// TokenFile provides a location to a token file.
	TokenFile string
}

func (s SecretStoreInfo) GetSecretStoreBaseURL() string {
	url := &url.URL{
		Scheme: "https",
		Host:   fmt.Sprintf("%s:%v", s.Host, s.Port),
	}
	return url.String()
}
