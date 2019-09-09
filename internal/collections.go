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
package internal

import (
	"github.com/globalsign/mgo"

	"github.com/edgexfoundry/docker-edgex-mongo/internal/pkg"
)

func cleanupUsers(db *mgo.Database) {
	_, err := db.C("system.users").RemoveAll(nil)
	if err != nil {
		pkg.LoggingClient.Error("Error during execution: " + err.Error())
	}
}

func createMetadataCollections(db *mgo.Database) {
	err := db.C("addressable").EnsureIndex(mgo.Index{Key: []string{"name"}, Name: "name_1", Unique: true})
	if err != nil {
		pkg.LoggingClient.Error("Error during execution: " + err.Error())
	}

	command := mgo.Collection{
		Database: db,
		Name:     "command",
		FullName: "db.command",
	}
	err = command.Create(&mgo.CollectionInfo{})
	if err != nil {
		pkg.LoggingClient.Error("Error during execution: " + err.Error())
	}

	err = db.C("device").EnsureIndex(mgo.Index{Key: []string{"name"}, Name: "name_1", Unique: true})
	if err != nil {
		pkg.LoggingClient.Error("Error during execution: " + err.Error())
	}

	err = db.C("deviceProfile").EnsureIndex(mgo.Index{Key: []string{"name"}, Name: "name_1", Unique: true})
	if err != nil {
		pkg.LoggingClient.Error("Error during execution: " + err.Error())
	}

	err = db.C("deviceReport").EnsureIndex(mgo.Index{Key: []string{"name"}, Name: "name_1", Unique: true})
	if err != nil {
		pkg.LoggingClient.Error("Error during execution: " + err.Error())
	}

	err = db.C("deviceService").EnsureIndex(mgo.Index{Key: []string{"name"}, Name: "name_1", Unique: true})
	if err != nil {
		pkg.LoggingClient.Error("Error during execution: " + err.Error())
	}

	err = db.C("provisionWatcher").EnsureIndex(mgo.Index{Key: []string{"name"}, Name: "name_1", Unique: true})
	if err != nil {
		pkg.LoggingClient.Error("Error during execution: " + err.Error())
	}

	err = db.C("schedule").EnsureIndex(mgo.Index{Key: []string{"name"}, Name: "name_1", Unique: true})
	if err != nil {
		pkg.LoggingClient.Error("Error during execution: " + err.Error())
	}

	err = db.C("scheduleEvent").EnsureIndex(mgo.Index{Key: []string{"name"}, Name: "name_1", Unique: true})
	if err != nil {
		pkg.LoggingClient.Error("Error during execution: " + err.Error())
	}
}

func createCoredataCollections(db *mgo.Database) {
	err := db.C("event").EnsureIndex(mgo.Index{Key: []string{"device"}, Name: "device_1", Unique: false})
	if err != nil {
		pkg.LoggingClient.Error("Error during execution: " + err.Error())
	}

	err = db.C("reading").EnsureIndex(mgo.Index{Key: []string{"device"}, Name: "device_1", Unique: false})
	if err != nil {
		pkg.LoggingClient.Error("Error during execution: " + err.Error())
	}

	err = db.C("valueDescriptor").EnsureIndex(mgo.Index{Key: []string{"name"}, Name: "name_1", Unique: true})
	if err != nil {
		pkg.LoggingClient.Error("Error during execution: " + err.Error())
	}
}

func createNotificationCollections(db *mgo.Database) {
	err := db.C("notification").EnsureIndex(mgo.Index{Key: []string{"slug"}, Name: "slug_1", Unique: true})
	if err != nil {
		pkg.LoggingClient.Error("Error during execution: " + err.Error())
	}

	transmission := mgo.Collection{
		Database: db,
		Name:     "transmission",
		FullName: "db.transmission",
	}
	err = transmission.Create(&mgo.CollectionInfo{})
	if err != nil {
		pkg.LoggingClient.Error("Error during execution: " + err.Error())
	}

	err = db.C("subscription").EnsureIndex(mgo.Index{Key: []string{"slug"}, Name: "slug_1", Unique: true})
	if err != nil {
		pkg.LoggingClient.Error("Error during execution: " + err.Error())
	}
}

func createSchedulerCollections(db *mgo.Database) {
	err := db.C("interval").EnsureIndex(mgo.Index{Key: []string{"name"}, Name: "name_1", Unique: true})
	if err != nil {
		pkg.LoggingClient.Error("Error during execution: " + err.Error())
	}

	err = db.C("intervalAction").EnsureIndex(mgo.Index{Key: []string{"name"}, Name: "name_1", Unique: true})
	if err != nil {
		pkg.LoggingClient.Error("Error during execution: " + err.Error())
	}
}

func createLoggingCollections(db *mgo.Database) {
	logEntry := mgo.Collection{
		Database: db,
		Name:     "logEntry",
		FullName: "db.logEntry",
	}
	err := logEntry.Create(&mgo.CollectionInfo{})
	if err != nil {
		pkg.LoggingClient.Error("Error during execution: " + err.Error())
	}
}

func createExportClientCollections(db *mgo.Database) {
	logEntry := mgo.Collection{
		Database: db,
		Name:     "exportConfiguration",
		FullName: "db.exportConfiguration",
	}
	err := logEntry.Create(&mgo.CollectionInfo{})
	if err != nil {
		pkg.LoggingClient.Error("Error during execution: " + err.Error())
	}
}
