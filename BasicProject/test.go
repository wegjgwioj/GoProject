package main

import "fmt"

func main() {
	fmt.Println("=== 测试 1: Map ===")
	var m map[string]int
	// 1. 读 nil map (绝对安全，返回 0)
	fmt.Println("读 nil map:", m["age"]) 
	
	// 2. 写 nil map (会崩溃，所以我用 recover 捉住它给你看)
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("写 nil map 导致:", r)
			}
		}()
		m["age"] = 18 // 💣 只有 make 后才能写！
	}()

	fmt.Println("\n=== 测试 2: Slice ===")
	s1 := []int{1, 2, 3}
	s2 := s1          // 此时 s1, s2 共享底层数组
	s2[0] = 999       // 修改 s2，s1 也变了！(s1[0] 现在是 999)
	
	s2 = append(s2, 4) // s2 容量不够，搬家了，从此 s1, s2 无关
	s2[0] = 888        // 修改新家，不影响旧家
	
	fmt.Println("s1[0] 最终是:", s1[0]) // 结果是 999

	fmt.Println("\n=== 测试 3: Defer ===")
	testDefer()
}
//defer 的执行顺序是后进先出（LIFO）的。 
func testDefer() {
	defer fmt.Println("A (第一个defer，最后执行)")
	fmt.Println("B (正常逻辑)")
	defer fmt.Println("C (第二个defer，最先执行)")
}