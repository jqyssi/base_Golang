package main

import "fmt"

func main() {

	a := "A"
	fmt.Println(&a)

	a = "b"
	fmt.Println(&a)

	s := []byte{1} // 分配存储1数组的内存空间，s结构体的array指针指向这个数组。
	fmt.Println(s)
	s = []byte{2}
	fmt.Println(s)
}
