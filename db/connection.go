package db

import (
	"context"
	"sync"

	"words-microservice/config"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const CONNECTED = "Successfully connected to database: %s"
const FAILED_TO_CONNECT = "Failed to connect to database: %s"

type MongoDatastore struct {
	Database *mongo.Database
	Session  *mongo.Client
	Config   *config.Config
}

func NewDatastore(config *config.Config) *MongoDatastore {
	var mongoDataStore *MongoDatastore
	db, session := connect(config)
	if db != nil && session != nil {
		mongoDataStore = new(MongoDatastore)
		mongoDataStore.Database = db
		mongoDataStore.Session = session
		mongoDataStore.Config = config
		return mongoDataStore
	}
	log.Fatalf(FAILED_TO_CONNECT, config.Database.DatabaseName)

	return nil
}

func (d MongoDatastore) GetWordsCollection() *mongo.Collection {
	return d.Database.Collection(d.Config.Database.CollectionName)
}

func connect(generalConfig *config.Config) (a *mongo.Database, b *mongo.Client) {
	var connectOnce sync.Once
	var db *mongo.Database
	var session *mongo.Client
	connectOnce.Do(func() {
		db, session = connectToMongo(generalConfig)
	})

	return db, session
}

func connectToMongo(generalConfig *config.Config) (a *mongo.Database, b *mongo.Client) {
	var err error
	session, err := mongo.NewClient(options.Client().ApplyURI(generalConfig.Database.ConnectionUri))
	if err != nil {
		log.Fatal(err)
	}
	session.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	var DB = session.Database(generalConfig.Database.DatabaseName)
	e := session.Ping(context.TODO(), nil)
	if e != nil {
		log.Fatal("Connection ping failed")
	}
	log.Info(CONNECTED, generalConfig.Database.DatabaseName)

	return DB, session
}
