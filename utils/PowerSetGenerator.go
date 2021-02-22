package main

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

func sortLettersAndConvert(sa *[][]rune) []string {
	var result []string
	for _, s := range *sa {
		SortLetters(&s)
		result = append(result, string(s))
	}
	return result
}

func removeDuplicates(s []string) []string {
	seen := make(map[string]struct{}, len(s))
	j := 0
	for _, v := range s {
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		s[j] = v
		j++
	}
	return s[:j]
}
