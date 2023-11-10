package main

func kthSmallest(root *TreeNode, k int) int {
	cc := 1

	var zx func(root *TreeNode, k int) int
	zx = func(root *TreeNode, k int) int {
		if root == nil {
			return -1
		}
		var ans int
		ans = zx(root.Left, k)
		if ans != -1 {
			return ans
		}
		if cc == k {
			return root.Val
		}
		cc++
		ans = zx(root.Right, k)
		if ans != -1 {
			return ans
		}
		return -1
	}

	a := zx(root, k)
	return a
}
