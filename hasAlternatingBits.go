package LeetCode

// https://leetcode-cn.com/problems/binary-number-with-alternating-bits/

func hasAlternatingBits(n int) bool {
	m := n ^ (n >> 1)
	return m&(m+1) == 0
}
