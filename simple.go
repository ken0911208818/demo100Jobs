package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	countChan := make(chan int, 1)
	for i := 0; i < 100; i++ {
		go func(value int, count chan int) {
			time.Sleep(time.Duration(rand.Int31n(1000)) * time.Millisecond)
			if value == 66 {
				countChan <- value
			} else {
				log.Println(fmt.Sprintf("The job is finish id : %d", value))
			}
		}(i, countChan)
	}
	for {
		select {
		case c := <-countChan:
			log.Println(fmt.Sprintf("The job is fail id : %d", c))
			goto end
		case <-time.After(500 * time.Millisecond):
			log.Println("Timeout")
			goto end
		}
	}
end:
}
