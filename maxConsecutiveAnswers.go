package LeetCode

func maxConsecutiveAnswers(answerKey string, k int) int {
	return max(maxConsectiveChar(answerKey, k, 'F'), maxConsectiveChar(answerKey, k, 'T'))
}

func maxConsectiveChar(answerKey string, k int, flag byte) (ans int) {
	num, left := 0, 0

	for right := range answerKey {
		if answerKey[right] == flag {
			num += 1
		}
		for num > k {
			if answerKey[left] == flag {
				num -= 1
			}
			left += 1
		}
		ans = max(ans, right-left+1)
	}
	return ans
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
