package main

func toIntArrBasic(text string) []int {
	var res []int
	for _, char := range text {
		res = append(res, int(char)-65)
	}
	return res
}

func toStringBasic(arr []int) string {
	var res string = ""
	for _, char := range arr {
		res += string(char + 65)
	}
	return res
}

func freq(arr string) (map[string]int, int) {
	var res map[string]int = make(map[string]int)
	for _, char := range arr {
		res[string(char)] += 1
	}
	return res, len(arr)
}

func rel_freq(arr string) (map[string]float32, int) {
	freqMap, count := freq(arr)
	var relFreqMap = make(map[string]float32)
	for key, value := range freqMap {
		relFreqMap[key] = float32(value) / float32(count)
	}
	return relFreqMap, count
}
