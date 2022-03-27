package Array

// https://leetcode-cn.com/problems/find-missing-observations/

func missingRolls(rolls []int, mean int, n int) []int {
	nList := make([]int, n)
	result := mean * (len(rolls) + n)

	for _, v := range rolls {
		result -= v
	}

	if result > 6*n || result < n {
		return nil
	}
	quotient, reminder := result/n, result%n
	for i, _ := range nList {
		if reminder > 0 {
			nList[i] = quotient + 1
			reminder--
		} else {
			nList[i] = quotient
		}
	}
	return nList
}
