package main

import "fmt"

// https://leetcode-cn.com/problems/binary-tree-level-order-traversal/

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return [][]int{}
	}
	var stack []*TreeNode
	stack = append(stack, root)
	for len(stack) > 0 {
		var levelValList []int
		var levelNodeList []*TreeNode
		for _, node := range stack {
			levelValList = append(levelValList, node.Val)
			if node.Left != nil {
				levelNodeList = append(levelNodeList, node.Left)
			}
			if node.Right != nil {
				levelNodeList = append(levelNodeList, node.Right)
			}
		}
		res = append(res, levelValList)
		stack = levelNodeList
	}
	fmt.Println(res)
	return res
}
