package main

import "fmt"

func maxProfit(prices []int) int {
	//如果少于1天 无法售出
	if len(prices) < 1 {
		return 0
	}
	//定义状态
	min, maxProfit := prices[0], 0

	for i := 1; i < len(prices); i++ {
		//状态转换  依次用每天的价格减去最低价格  依次和之前的最大利润比较如果有更大例如则换最大利润
		if prices[i]-min > maxProfit {
			maxProfit = prices[i] - min
		}
		//状态转换  先设置第1天为价格最低  依次遍历如果有低价就设置为最低价
		if prices[i] < min {
			min = prices[i]
		}
	}
	//返回最大利润
	return maxProfit
}

func main() {
	price := []int{7,1,9,3,6,4}
	max :=maxProfit(price)
	fmt.Println("最大利润：",max)

}
