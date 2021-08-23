package concurrency

import (
	"fmt"
	"time"
	"sync"
)

func compute(value int, wg *sync.WaitGroup) {
	for i := 1; i <= value; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
	wg.Done()
}

func GoRoutineMain() {
	fmt.Println("Concurrency with GoRoutines!")
	var wg sync.WaitGroup
	wg.Add(2)
	go compute(5, &wg)
	go compute(5, &wg)
	wg.Wait();

}