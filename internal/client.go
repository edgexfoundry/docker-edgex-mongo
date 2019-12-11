package internal

import (
	"fmt"

	"github.com/globalsign/mgo"

	"github.com/edgexfoundry/docker-edgex-mongo/internal/pkg"
)

var DatabaseCollectionsMap = map[string]func(db *mgo.Database){
	"authorization":       nil,
	"metadata":            createMetadataCollections,
	"coredata":            createCoredataCollections,
	"rulesengine":         nil,
	"notifications":       createNotificationCollections,
	"scheduler":           createSchedulerCollections,
	"logging":             createLoggingCollections,
	"application-service": createApplicationServiceCollections,
}

type DBInitClient struct {
	Configuration *pkg.Configuration
}

func (client *DBInitClient) PopulateDatabase() (err error) {
	session, err := pkg.GetSession(client.Configuration)

	if err != nil {
		return
	}

	defer session.Close()

	//User clearance should be done first, so further created users will be present.
	client.createDatabase(session, "admin", cleanupUsers)

	for dbName, createCollectionsFunc := range DatabaseCollectionsMap {
		client.createDatabase(session, dbName, createCollectionsFunc)
	}
	return
}

func (client *DBInitClient) createDatabase(session *mgo.Session, dbName string, createCollectionsFunc func(db *mgo.Database)) {
	pkg.LoggingClient.Info(fmt.Sprintf("Settting up %v database", dbName))
	db := mgo.Database{
		Session: session,
		Name:    dbName,
	}

	err := db.UpsertUser(&mgo.User{
		Username: client.Configuration.Databases[dbName].Username,
		Password: client.Configuration.Databases[dbName].Password,
		Roles: []mgo.Role{
			mgo.RoleReadWrite,
		},
	})

	if err != nil {
		pkg.LoggingClient.Error("Error during execution: " + err.Error())
	}
	if createCollectionsFunc != nil {
		createCollectionsFunc(&db)
	}
}
