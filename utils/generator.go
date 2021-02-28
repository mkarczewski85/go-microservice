package utils

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
	for _, letter := range *letters {
		var tmpSlice [][]rune
		for _, tmp := range powerSet {
			tmpSlice = append(tmpSlice, append(tmp, letter))
		}
		powerSet = append(powerSet, tmpSlice...)
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
