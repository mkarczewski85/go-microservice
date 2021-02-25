package service

import (
	"words-microservice/model"
	"words-microservice/utils"
)

// TODO
func FindMatchedWords(input *model.WordsSearchInput) []model.Word {
	powerSet := utils.GenerateSearchCriteria([]rune(input.Letters))
	return nil
}
