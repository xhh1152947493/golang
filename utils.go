package main

func MaxInt(arr []int) (int, int) {
	/*
	 */
	maxVal := arr[0]
	maxIndex := 0

	for i := 1; i < len(arr); i++ {
		if maxVal < arr[i] {
			maxVal = arr[i]
			maxIndex = i
		}
	}
	return maxIndex, maxVal
}
