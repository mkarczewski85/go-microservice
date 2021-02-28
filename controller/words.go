package controller

import (
	"context"
	"net/http"
	"words-microservice/db"
	"words-microservice/model"
	"words-microservice/utils"

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

	powerSet := utils.GenerateSearchCriteria([]rune(input.Letters))
	findQuery := convertToQuery(&powerSet)

	result := findMatchedWords(&findQuery)
	c.JSON(http.StatusOK, gin.H{"result": result})
}

func findMatchedWords(filters *bson.M) []bson.M {
	var result []bson.M
	collection := Datastore.GetWordsCollection()
	cur, err := collection.Find(context.TODO(), filters)
	if err != nil {
		log.Fatal("Error during finding filtered documents", err)
	}

	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		var entry bson.M
		err = cur.Decode(&entry)
		if err != nil {
			log.Fatal("Error during decoding the document", err)
		}
		result = append(result, entry)
	}
	return result
}

func convertToQuery(powerSet *[]string) bson.M {
	var arr bson.A
	for _, s := range *powerSet {
		arr = append(arr, s)
	}
	findQuery := bson.M{"sortedLetters": bson.M{"$in": arr}}
	return findQuery
}
