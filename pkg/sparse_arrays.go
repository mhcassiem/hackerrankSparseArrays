package pkg

func MatchingStrings(strings []string, queries []string) []int32 {

	var result []int32

	for _, qValue := range queries {
		var counts int32 = 0
		for _, dValue := range strings {
			if dValue == qValue {
				counts++
			}
		}
		result = append(result, counts)
	}

	return result
}
