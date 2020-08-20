package main

func largestArea(samples [][]int32) int32 {
	var n int
	if n = len(samples); n == 0 {
		return 0
	}
	m := int32(len(samples[0]))

	// make result matrix
	var sum [][]int32
	for i := 0; i < n; i++ {
		tmp := make([]int32, m)
		sum = append(sum, tmp)
	}

	for j := int32(0); j < m; j++ {
		var add int32 = 0
		for i := 0; i < n; i++ {
			if samples[i][j] == 1 {
				add += 1
			} else {
				add = 0
			}
			sum[i][j] = add
		}
	}

	var largestArea int32
	for i := 0; i < n; i++ {
		var cLen int32
		var count int32
		for j := int32(0); j < m; j++ {
			if sum[i][j] != cLen {
				count = 1
				cLen = sum[i][j]
			} else {
				count++
			}
			if count >= cLen && cLen > largestArea {
				largestArea = cLen
			}
		}
	}
	return largestArea
}