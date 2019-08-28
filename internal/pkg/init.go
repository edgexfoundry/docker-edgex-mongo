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
	"strconv"
	"time"

	"github.com/globalsign/mgo"

	"github.com/edgexfoundry/go-mod-core-contracts/clients/logger"
)

var LoggingClient logger.LoggingClient

func GetSession(config *Configuration) (*mgo.Session, error) {
	connectionString := config.Mongo.Host + ":" + strconv.Itoa(config.Mongo.Port)
	var session *mgo.Session
	var err error
	until := time.Now().Add(time.Millisecond * time.Duration(config.Service.BootTimeout))
	for time.Now().Before(until) {
		session, err = mgo.DialWithTimeout(connectionString, time.Duration(config.Mongo.Timeout)*time.Millisecond)
		if err != nil {
			LoggingClient.Error(fmt.Sprintf("cannot connect to the database: %s", err.Error()))
		} else {
			break
		}
		time.Sleep(time.Second * time.Duration(1))
	}
	return session, err
}
