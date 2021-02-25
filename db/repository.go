package db

import (
	"context"
	"words-microservice/model"
	"words-microservices/config"

	"go.mongodb.org/mongo-driver/bson"
)

func FindMatchedWords(dbConnection *MongoDatastore, filters bson.M, config *config.GeneralConfig) []*model.Word {
	var words []*model.Word
	collection := dbConnection.GetWordsCollection()
	cur, err := collection.Find(context.TODO(), filters)
	if err != nil {
		dbConnection.Logger.Fatal("Error during finding filtered documents", err)
	}
	for cur.Next(context.TODO()) {
		var word model.Word
		err = cur.Decode(&word)
		if err != nil {
			dbConnection.Logger.Fatal("Error during decoding the document", err)
		}
		words = append(words, &word)
	}
	return words
}
