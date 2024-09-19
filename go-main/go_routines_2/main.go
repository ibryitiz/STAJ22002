// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {

// 	wg := sync.WaitGroup{}

// 	wg.Add(2)

// 	go func() {
// 		defer wg.Done()
// 		fmt.Println("f1")
// 	}()

// 	go func() {
// 		defer wg.Done()
// 		fmt.Println("f2")
// 	}()

// 	wg.Wait() // Blocked

// 	fmt.Println("End of the main")

// }

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	startTime := time.Now()

	wg := sync.WaitGroup{}

	wg.Add(3)
	go func() {
		defer wg.Done()
		fmt.Println("f1")
		time.Sleep(3 * time.Second)
	}()
	go func() {
		defer wg.Done()
		fmt.Println("f2")
		time.Sleep(3 * time.Second)
	}()
	go func() {
		defer wg.Done()
		fmt.Println("f3")
		time.Sleep(3 * time.Second)
	}()

	wg.Wait()
	fmt.Println("Passed time : ", time.Now().Sub(startTime))
}
