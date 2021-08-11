package utils

func UniqInts(array []int) []int {
	result := []int{}
	set := make(map[int]bool)
	for _, item := range array {
		if _, ok := set[item]; !ok {
			set[item] = true
			result = append(result, item)
		}
	}
	return result
}
