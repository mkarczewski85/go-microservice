package utils

import (
	"words-microservice/model"

	"go.mongodb.org/mongo-driver/bson"
)

func MapToOutput(documents *[]bson.M) []model.WordSearchOutput {
	var result []model.WordSearchOutput
	for _, d := range *documents {
		var model model.WordSearchOutput
		bsonBytes, _ := bson.Marshal(d)
		bson.Unmarshal(bsonBytes, &model)
		result = append(result, model)
	}
	return result
}
