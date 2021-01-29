package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

const (
	Job     = 100 //平行處理數量
	Timeout = 40  // 毫秒
)

func main() {
	outChan := make(chan string, 100)
	errChan := make(chan error)
	finishChan := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(Job)
	for i := 0; i < 100; i++ {
		go func(value int, group *sync.WaitGroup, out chan string, err chan error) {
			time.Sleep(time.Duration(rand.Int63n(1000)) * time.Millisecond)
			if value == 77 {
				err <- errors.New(fmt.Sprintf("The job is fail id: %d", value))
			} else {
				out <- fmt.Sprintf("The job is finished id : %d", value)
			}
			wg.Done()
		}(i, &wg, outChan, errChan)
	}

	go func() {
		wg.Wait() // 做完100個才會執行close
		close(finishChan)
	}()
Loop:
	for {
		select {
		case out := <-outChan:
			log.Println(out)
		case err := <-errChan:
			log.Println(err)
			break Loop
		case <-finishChan:
			break Loop
		case <-time.After(Timeout * time.Millisecond):
			log.Println("Timeout")
			break Loop
		}
	}
}
