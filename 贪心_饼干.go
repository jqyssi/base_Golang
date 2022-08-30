package main

import (
	"fmt"
	"sort"
)

func main() {

	var p = []int {1,2,4,6}
	var c = []int {1,2,3,5}
	fmt.Println("吃饱的孩子：",findContentChildren(p , c))

}
/**
 * @Description :LeetCode贪心算法 https://leetcode.cn/problems/assign-cookies/
 * @Author : ranksin
 * @Date : 2022/5/14 5:51 下午
 */
//排序后，局部最优
func findContentChildren(g []int, s []int) int {
	sort.Ints(g) //数组排序之后，数组变为排序之后的
	sort.Ints(s)

	// 给孩子和饼干一个索引
	child := 0
	for sIdx := 0; child < len(g) && sIdx < len(s); sIdx++ {
		if s[sIdx] >= g[child] {//如果饼干的大小大于或等于孩子的为空则给与，否则不给予，继续寻找选一个饼干是否符合
			child++
		}
	}

	return child
}
