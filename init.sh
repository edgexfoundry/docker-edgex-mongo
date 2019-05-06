#!/bin/sh
# read root key file of vault from volume share and parse key
JSON=$(cat </vault/config/assets/resp-init.json)
#TODO: switch to root so that resp-init.json can be read
#JSON=$(cat </tmp/resp-init.json)
RKEY=$(echo $JSON | sed 's/.*"\(.*\)"[^"]*$/\1/')

echo $RKEY
# reach vault to get the credential
# curl edgex-vault:8200/v1/secret/mongodb/my-secret
# $SECRET_SERVICE_HOST needs to be defined in the dockercompose file
CREDS=$(curl -k --header "X-Vault-Token: $RKEY" https://$SECRET_SERVICE_HOST:8200/v1/secret/edgex/mongodbinit)

#CREDS=$(cat </tmp/cred.json)
ADMIN=$(echo $CREDS | jq -r ".data .admin")
echo $ADMIN
ADMINPASSWD=$(echo $CREDS | jq -r ".data .adminpasswd")
echo $ADMINPASSWD
METADATA=$(echo $CREDS | jq -r ".data .metadata")
METADATAPASSWD=$(echo $CREDS | jq -r ".data .metadatapasswd")
COREDATA=$(echo $CREDS | jq -r ".data .coredata")
COREDATAPASSWD=$(echo $CREDS | jq -r ".data .coredatapasswd")
RULESENGINE=$(echo $CREDS | jq -r ".data .rulesengine")
RULESENGINEPASSWD=$(echo $CREDS | jq -r ".data .rulesenginepasswd")
NOTIFICATIONS=$(echo $CREDS | jq -r ".data .notifications")
NOTIFICATIONSPASSWD=$(echo $CREDS | jq -r ".data .notificationspasswd")
LOGGING=$(echo $CREDS | jq -r ".data .logging")
LOGGINGPASSWD=$(echo $CREDS | jq -r ".data .loggingpasswd")
SCHEDULER=$(echo $CREDS | jq -r ".data .scheduler")
SCHEDULERPASSWD=$(echo $CREDS | jq -r ".data .schedulerpasswd")

#ADMIN="admin"
#ADMINPASSWD="pass"
#METADATA="metadata"
#METADATAPASSWD="pass"
#COREDATA="coredata"
#COREDATAPASSWD="pass"
#RULESENGINE="rules_engine"
#RULESENGINEPASSWD="pass"
#NOTIFICATIONS="notifications"
#NOTIFICATIONSPASSWD="pass"
#LOGGING="logging"
#LOGGINGPASSWD="pass"
#SCHEDULER="scheduler"
#SCHEDULERPASSWD="pass"

authDatabase='admin'

"${mongo[@]}" "$authDatabase" <<EOJS
    //Create user for security service in Mongo    
    db.createUser({ user: "$ADMIN",pwd: "$ADMINPASSWD",roles: [ { role: "root", db: "admin" } ]});
    db.auth('$ADMIN', '$ADMINPASSWD');    
    //Create keystore collection
    db.createCollection("keyStore");
    db.keyStore.insert( { xDellAuthKey: "x-dell-auth-key", secretKey: "EDGEX_SECRET_KEY" } );
    //Create Service Mapping
    db.createCollection("serviceMapping");
    db.serviceMapping.insert( { serviceName: "coredata", serviceUrl: "http://localhost:48080/" });
    db.serviceMapping.insert( { serviceName: "metadata", serviceUrl: "http://localhost:48081/" });
    db.serviceMapping.insert( { serviceName: "command", serviceUrl: "http://localhost:48082/" });
    db.serviceMapping.insert( { serviceName: "rules", serviceUrl: "http://localhost:48084/" });
    db.serviceMapping.insert( { serviceName: "notifications", serviceUrl: "http://localhost:48060/" });
    db.serviceMapping.insert( { serviceName: "logging", serviceUrl: "http://localhost:48061/" });
    
    db=db.getSiblingDB('metadata');
    db.createUser({ user: "$METADATA",
    pwd: "$METADATAPASSWD",
    roles: [
        { role: "readWrite", db: "metadata" }
    ]
    });
    db.createCollection("addressable");
    db.addressable.createIndex({name: 1}, {unique: true});
    db.createCollection("command");
    db.createCollection("device");
    db.device.createIndex({name: 1}, {unique: true});
    db.createCollection("deviceManager");
    db.deviceManager.createIndex({name: 1}, {unique: true});
    db.createCollection("deviceProfile");
    db.deviceProfile.createIndex({name: 1}, {unique: true});
    db.createCollection("deviceReport");
    db.deviceReport.createIndex({name: 1}, {unique: true});
    db.createCollection("deviceService");
    db.deviceService.createIndex({name: 1}, {unique: true});
    db.createCollection("provisionWatcher");
    db.provisionWatcher.createIndex({name: 1}, {unique: true});
    db.createCollection("schedule");
    db.schedule.createIndex({name: 1}, {unique: true});
    db.createCollection("scheduleEvent");
    db.scheduleEvent.createIndex({name: 1}, {unique: true});

    db=db.getSiblingDB('coredata');
    db.createUser({ user: "$COREDATA",
    pwd: "$COREDATAPASSWD",
    roles: [
        { role: "readWrite", db: "coredata" }
    ]
    });
    db.createCollection("event");
    db.createCollection("reading");
    db.createCollection("valueDescriptor");
    db.valueDescriptor.createIndex({name: 1}, {unique: true});

    db=db.getSiblingDB('rules_engine_db');
    db.createUser({ user: "$RULESENGINE",
    pwd: "$RULESENGINEPASSWD",
    roles: [
        { role: "readWrite", db: "rules_engine_db" }
    ]
    });

    db=db.getSiblingDB('notifications');
    db.createUser({ user: "$NOTIFICATIONS",
    pwd: "$NOTIFICATIONSPASSWD",
    roles: [
        { role: "readWrite", db: "notifications" }
    ]
    });
    db.createCollection("notification");
    db.createCollection("transmission");
    db.createCollection("subscription");
    db.notification.createIndex({slug: 1}, {unique: true});
    db.subscription.createIndex({slug: 1}, {unique: true});

    db=db.getSiblingDB('scheduler');
    db.createUser({ user: "$SCHEDULER",
    pwd: "$SCHEDULERPASSWD",
    roles: [
        { role: "readWrite", db: "scheduler" }
    ]
    });
    db.createCollection("interval");
    db.createCollection("intervalAction");
    db.interval.createIndex({name: 1}, {unique: true});
    db.intervalAction.createIndex({name: 1}, {unique: true});

    db=db.getSiblingDB('logging');
    db.createUser({ user: "$LOGGING",
    pwd: "$LOGGINGPASSWD",
    roles: [
        { role: "readWrite", db: "logging" }
    ]
    });
    db.createCollection("logEntry");
EOJS