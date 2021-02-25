package db

import (
	"context"
	"sync"

	log "github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"words-microservice/config"
)

const CONNECTED = "Successfully connected to database: %v"
const FAILED_TO_CONNECT = "Failed to connect to database: %v"

type MongoDatastore struct {
	Database *mongo.Database
	Session  *mongo.Client
	Logger   *log.Logger
	Config   *config.GeneralConfig
}

func NewDatastore(config *config.GeneralConfig, logger *log.Logger) *MongoDatastore {
	var mongoDataStore *MongoDatastore
	db, session := connect(config, logger)
	if db != nil && session != nil {
		mongoDataStore = new(MongoDatastore)
		mongoDataStore.Database = db
		mongoDataStore.Logger = logger
		mongoDataStore.Session = session
		mongoDataStore.Config = config
		return mongoDataStore
	}
	logger.Fatalf(FAILED_TO_CONNECT, config.Database.DatabaseName)

	return nil
}

func (d MongoDatastore) GetWordsCollection() *mongo.Collection {
	return d.Database.Collection(d.Config.Database.CollectionName)
}

func connect(generalConfig *config.GeneralConfig, logger *log.Logger) (a *mongo.Database, b *mongo.Client) {
	var connectOnce sync.Once
	var db *mongo.Database
	var session *mongo.Client
	connectOnce.Do(func() {
		db, session = connectToMongo(generalConfig, logger)
	})

	return db, session
}

func connectToMongo(generalConfig *config.GeneralConfig, logger *log.Logger) (a *mongo.Database, b *mongo.Client) {
	var err error
	session, err := mongo.NewClient(options.Client().ApplyURI(generalConfig.Database.ConnectionUri))
	if err != nil {
		logger.Fatal(err)
	}
	session.Connect(context.TODO())
	if err != nil {
		logger.Fatal(err)
	}
	var DB = session.Database(generalConfig.Database.DatabaseName)
	logger.Info(CONNECTED, generalConfig.Database.DatabaseName)

	return DB, session
}
