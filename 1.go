package main

import "fmt"

/*
   恢复二叉搜索树
*/
func main() {
	r := &TreeNode{Val: 3}
	r.Left = &TreeNode{Val: 1}
	r.Right = &TreeNode{Val: 2}
	recoverTree(r)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	res   []int
	pre   *TreeNode
	count int
)

/**
 *
 * @param root TreeNode类
 * @return void
 */
func recoverTree(root *TreeNode) {
	res = make([]int, 2)
	pre = nil
	count = 0
	// write code here
	findError(root)
	fmt.Println(res)
	recover(root)
	fmt.Println(root)
}

func findError(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil {
		findError(root.Left)
	}
	val := root.Val
	if pre != nil && pre.Val > val {
		count++
		if count == 1 {
			res[0] = val
			res[1] = pre.Val
		} else {
			res[0] = val
		}
	}
	pre = root
	if root.Right != nil {
		findError(root.Right)
	}
}

func recover(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil {
		recover(root.Left)
	}
	if root.Val == res[0] {
		root.Val = res[1]
	} else if root.Val == res[1] {
		root.Val = res[0]
	}
	if root.Right != nil {
		recover(root.Right)
	}
}
