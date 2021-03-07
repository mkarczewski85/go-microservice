package controller

import (
	"context"
	"net/http"
	"words-microservice/db"
	"words-microservice/model"
	"words-microservice/service"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

var Datastore *db.MongoDatastore

// POST
func SearchForWordsHandler(c *gin.Context) {
	var input model.WordsSearchInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	findQuery := service.GenerateSearchCriteria(input)
	foundDocuments := findMatchedWords(&findQuery)
	c.JSON(http.StatusOK, gin.H{"result": foundDocuments})
}

func findMatchedWords(filters *bson.M) []model.WordSearchOutput {
	var result []model.WordSearchOutput
	collection := Datastore.GetWordsCollection()
	cur, err := collection.Find(context.TODO(), filters)
	if err != nil {
		log.Fatal("Error during finding filtered documents", err)
	}

	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		var entry model.WordSearchOutput
		err = cur.Decode(&entry)
		if err != nil {
			log.Fatal("Error during decoding the document", err)
		}
		result = append(result, entry)
	}
	return result
}
