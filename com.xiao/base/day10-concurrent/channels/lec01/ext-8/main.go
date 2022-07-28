package main

import "fmt"

/**
 * @Author safoti
 * @Date Created in 2022/7/28
 * @Description    超出容量的缓存
 **/
func main() {
	 var ch chan int  //int 类型的channel
	fmt.Printf("ch: %v, len: %v, cap: %v\n", ch, len(ch), cap(ch))
	// creating channel with capacity (buffered channels)
	// ----
	ch = make(chan int,1)
	//有地址了
	
	fmt.Printf("ch: %v, len: %v, cap: %v\n", ch, len(ch), cap(ch))

	// sending data to buffered channel
	// ----
	ch <- 2
	ch <- 4

}
