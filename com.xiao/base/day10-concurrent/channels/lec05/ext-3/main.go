package main

import (
	"fmt"
	"time"
)

/**
 * @Author safoti
 * @Date Created in 2022/7/28
 * @Description  closing a channel
 **/


func main() {
	// get a notification form a channel
	// ----
	fmt.Printf("Message 1 at: %v\n", time.Now())
	alarm := notifyAfter(1 * time.Second)
	<-alarm
	fmt.Printf("Message 2 at: %v\n", time.Now())
}
func notifyAfter(delay time.Duration) (out chan time.Time) {
	out = make(chan time.Time)
	go func() {
		time.Sleep(delay)
		out <- time.Now()
		close(out)
	}()
	return out
}