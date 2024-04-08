package main

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//求二叉树中最大路径和

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32
	var maxGain func(node *TreeNode) int
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		//左边能提供给上面多大？最小为0
		leftGain := max(0, maxGain(node.Left))
		//右边能提供给上面多大？最小为0
		rightGain := max(0, maxGain(node.Right))
		//计算最大路径和，无非自己，自己和左边最大，自己和右边最大，自己和左和右
		//因为左右都做了非负处理，所以就只需要全加起来
		priceNewPath := node.Val + leftGain + rightGain
		//记录最大路径和
		maxSum = max(maxSum, priceNewPath)
		//递归嘛，返回的肯定是要给上层的
		//能提供给上层什么，只能一条路，而且必须得走该节点
		return node.Val + max(leftGain, rightGain)
	}
	maxGain(root)
	return maxSum
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {

}
