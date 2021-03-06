package main

import "fmt"

/**
 * @Author safoti
 * @Date Created in 2022/7/28
 * @Description
 **/
func main() {
	 var ch chan int  //int 类型的channel
	fmt.Printf("ch: %v, len: %v, cap: %v\n", ch, len(ch), cap(ch))
	// creating channel with capacity (buffered channels)
	// ----
	ch = make(chan int,2)
	//有地址了
	
	fmt.Printf("ch: %v, len: %v, cap: %v\n", ch, len(ch), cap(ch))

	// sending data to buffered channel
	// ----
	ch <- 2
	ch <- 4
	fmt.Printf("ch len: %v, cap: %v\n", len(ch), cap(ch))

	// receiving data from an empty channel
	// ----
	fmt.Printf("1st value from ch: %v, len: %v\n", <-ch, len(ch))
	fmt.Printf("2nd value from ch: %v, len: %v\n", <-ch, len(ch))
	// next command will fail, since the channel is empty
	fmt.Printf("3rd value from ch: %v, len: %v\n", <-ch, len(ch))

}
