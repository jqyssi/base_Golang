package main

import (
	"fmt"
	"sort"
)

func main() {
	peo:=[]int{3,5,3,4}
	boatNeed := numRescueBoats(peo,5)
	fmt.Println("需要船：",boatNeed)
}

func numRescueBoats(people []int, limit int) (ans int) {
	sort.Ints(people)
	//排序之后 最轻的下标为light  最重的下标为 heavy
	light, heavy := 0, len(people)-1
	for light <= heavy {
		//如果最轻和最重不能乘一个船 则移动最重下标
		if people[light]+people[heavy] > limit {
			heavy--
			//如果最轻和最重能乘一个船 则移动柜最轻的下标和最重的下标
		} else {
			light++
			heavy--
		}
		ans++
	}
	return ans
}

