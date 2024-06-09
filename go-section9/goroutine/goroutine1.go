// 고루틴(Goroutine) 기초 (1)
package main

import "fmt"
import "time"

func exe1() {
	fmt.Println("exe1 func start", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("exe1 func end", time.Now())
}

func exe2() {
	fmt.Println("exe2 func start", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("exe2 func end", time.Now())
}

func exe3() {
	fmt.Println("exe3 func start", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("exe3 func end", time.Now())
}

func main() {
	exe1() // 가장 먼저 실행 (일반적인 실행 흐름)

	fmt.Println("Main Routine Start : ", time.Now())
	go exe2()
	go exe3()
	fmt.Scanln()
	fmt.Println("Main Routine End : ", time.Now())
}

/*
고루틴(Goroutine)

- 별도의 흐름을 가지고 실행된다. (운영체제가 알아서 수행함, Main 쓰레드와 별개의 흐름을 가진다)
- 타 언어의 쓰레드(Thread)와 비슷한 기능을 한다.
- 생성 방법이 매우 간단하고, 리소스를 매우 적게 사용한다.
- 수많은 고루틴을 동시에 생성하고 실행할 수 있다.
- 비동기적 함수 루틴을 실행하며, 매우 적은 용량을 차지한다.
- 채널을 통해 고루틴 간에 통신한다.
- 공유 메모리를 사용할 때 정확한 동기화 코딩이 필요하다.
- 싱글 루틴에 비해 항상 빠른 처리 결과를 보장하지는 않는다.
*/

/*
멀티 쓰레드 장점과 단점

- 장점 : 응답성 향상, 자원 공유를 효율적으로 활용 및 사용, 작업이 분리되어 코드가 간결해진다.
- 단점 : 구현하기 어렵다, 테스트 및 디버깅이 어렵다, 전체 프로세스의 사이드 이펙트가 발생할 가능성이 높아진다, 성능 저하가 발생할 수 있다, 동기화 코딩에 대한 이해 필요, 데드락...
*/
