package main

import (
	"runtime"
	"sync"
)

func main() {

	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(7)
kek:
	for i := 0; i < 7; i++ {
		//go func(i int) {
		//	defer wg.Done()
		//	fmt.Println(i)
		//
		//}(i)
		switch 1 {
		case 2:
		case 1:
			break kek
		default:

		}
	}

	//wg.Wait()

}
