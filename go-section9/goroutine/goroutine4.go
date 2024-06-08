// 고루틴(Goroutine)기초(4)
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//고루틴(Goroutine)
	//클로저 사용 예제

	//예제1
	runtime.GOMAXPROCS(1) // CPU를 하나만 사용

	s := "Goroutine Closure : "

	for i := 0; i < 1000; i++ { //고루틴 1000개 실행
		go func(n int) {
			fmt.Println(s, n, " - ", time.Now())
		}(i) //반복문 클로저는 일반적으로 즉시 실행 But 고루틴 클로저는 가장 나중에 실행(반복문 종료 후)
	}
	/*
		//실행 결과 비교
		for i := 0; i < 1000; i++ {
			go func() {
				fmt.Println(s, n, " - ", time.Now())
			}()
		}
	*/

	time.Sleep(5 * time.Second)
}
