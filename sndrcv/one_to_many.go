package sndrcv

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const ff = 5

func OneToMany() {
	maxReceivers := 5

	dataCh := make(chan int)
	stopCh := make(chan struct{})
	wg := sync.WaitGroup{}

	sender := func() {
		counter := 1000
		for {
			doWork()

			select {
			case dataCh <- counter:
				counter++
			case <-stopCh:
				fmt.Println("stop send")
				close(dataCh)
				return
			}
		}
	}

	receiver := func(i int) {
		for v := range dataCh {
			process(i, v)
		}

		fmt.Printf("rcv #%d stopped\n", i)
		wg.Done()
	}

	go sender()
	wg.Add(maxReceivers)
	for i := 0; i < maxReceivers; i++ {
		go receiver(i + 1)
	}

	time.Sleep(4 * time.Second)
	stopCh <- struct{}{}

	wg.Wait()
}

func doWork() {
	t := rand.Int() % 500
	time.Sleep(time.Duration(300+t) * time.Millisecond)
}

func process(i, v int) {
	t := rand.Int() % 1000
	time.Sleep(time.Duration(500+t) * time.Millisecond)
	fmt.Printf("#%d => %d\n", i, v)
}
