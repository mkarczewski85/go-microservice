package service

import (
	"fmt"
	"words-microservice/model"
	"words-microservice/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GenerateSearchCriteria(input model.WordsSearchInput) bson.M {
	switch input.Mode {
	case 1:
		return searchAllWords([]rune(input.Letters))
	case 2:
		startingLetter := []rune(input.AuxiliaryLetter)[0]
		letters := []rune(input.Letters)
		letters = append(letters, startingLetter)
		return searchAllWordsStartingWith(letters, startingLetter)
	case 3:
		endingLetter := []rune(input.AuxiliaryLetter)[0]
		letters := []rune(input.Letters)
		letters = append(letters, endingLetter)
		return searchAllWordsEndingWith(letters, endingLetter)
	default:
		panic("bad input")
	}
}

func searchAllWords(letters []rune) bson.M {
	powerSet := utils.GetPowerset([]rune(letters))
	var arr bson.A
	for _, s := range powerSet {
		arr = append(arr, s)
	}
	findQuery := bson.M{"sortedLetters": bson.M{"$in": arr}}
	return findQuery
}

func searchAllWordsStartingWith(letters []rune, startingLetter rune) bson.M {
	powerSet := utils.GetPowerset([]rune(letters))
	inQuery := inQuery(powerSet)
	startingWithQuery := bson.M{"word": regexStartingWith(startingLetter)}
	sortedLettersQuery := bson.M{"sortedLetters": bson.M{"$in": inQuery}}
	findQuery := bson.M{"$and": bson.A{sortedLettersQuery, startingWithQuery}}
	fmt.Println(findQuery)
	return findQuery
}

func searchAllWordsEndingWith(letters []rune, endingLetter rune) bson.M {
	powerSet := utils.GetPowerset([]rune(letters))
	inQuery := inQuery(powerSet)
	endingWithQuery := bson.M{"word": regexEndingWith(endingLetter)}
	sortedLettersQuery := bson.M{"sortedLetters": bson.M{"$in": inQuery}}
	findQuery := bson.M{"$and": bson.A{sortedLettersQuery, endingWithQuery}}
	fmt.Println(findQuery)
	return findQuery
}

func inQuery(powerSet []string) bson.A {
	var inQuery bson.A
	for _, s := range powerSet {
		inQuery = append(inQuery, s)
	}
	return inQuery
}

func regexStartingWith(letter rune) primitive.Regex {
	regex := fmt.Sprintf("^%s", string(letter))
	return primitive.Regex{Pattern: regex}
}

func regexEndingWith(letter rune) primitive.Regex {
	regex := fmt.Sprintf("%s$", string(letter))
	return primitive.Regex{Pattern: regex}
}
