//채널(Channel) 심화(5)

package main

import "fmt"
import "time"

func main() {
	//채널(Channel)

	//예제1
	ch1 := make(chan int)
	ch2 := make(chan string)

	go func() {
		for {
			num := <-ch1 //값 수신
			fmt.Println("ch1 : ", num)
			time.Sleep(250 * time.Millisecond)
		}
	}()

	go func() {
		for {
			ch2 <- "Golang Hi!" //값 발신
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for {
			select {
			case ch1 <- 777: //발신용도
			case str := <-ch2:
				fmt.Println("ch2 : ", str)
			}
		}
	}()
	time.Sleep(7 * time.Second)
}
