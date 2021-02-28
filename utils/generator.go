package utils

import (
	"math"
	"strconv"
)

func GenerateSearchCriteria(letters []rune) []string {
	if letters == nil {
		return nil
	}
	powerSet := powerSet(&letters)
	sortedStrings := sortLettersAndConvert(&powerSet)
	return removeDuplicates(sortedStrings)
}

func powerSet(letters *[]rune) [][]rune {
	powerSet := [][]rune{[]rune{}}
	length := len(*letters)
	total := math.Pow(2, float64(length))
	for i := 0; i < int(total); i++ {
		var tmpSet []rune
		num := strconv.FormatInt(int64(i), 2)
		for len(num) < length {
			num = "0" + num
		}
		for j := 0; j < len(num); j++ {
			if num[j] == '1' {
				tmpSet = append(tmpSet, (*letters)[j])
			}
		}
		powerSet = append(powerSet, tmpSet)
	}
	return powerSet
}

func sortLettersAndConvert(powerSet *[][]rune) []string {
	var result []string
	for _, s := range *powerSet {
		SortLetters(&s)
		result = append(result, string(s))
	}
	return result[1:]
}

func removeDuplicates(powerSet []string) []string {
	seen := make(map[string]struct{}, len(powerSet))
	j := 0
	for _, s := range powerSet {
		if _, ok := seen[s]; ok {
			continue
		}
		seen[s] = struct{}{}
		powerSet[j] = s
		j++
	}
	return powerSet[:j]
}
