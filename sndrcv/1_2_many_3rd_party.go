package sndrcv

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

// close signal is made by third party
// extra signal channel to notify the sender to close data channel
func One2Many() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	const (
		Max             = 1000000
		NumReceivers    = 100
		NumThirdParties = 15
	)
	var (
		dataCh  = make(chan int)
		closing = make(chan struct{})
		closed  = make(chan struct{})
	)
	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(NumReceivers)

	stop := func() {
		select {
		case closing <- struct{}{}:
			<-closed
		case <-closed:
		}
	}

	// third party
	for i := 0; i < NumThirdParties; i++ {
		go func() {
			r := 1 + rand.Intn(3)
			time.Sleep(time.Duration(r) * time.Second)
			stop()
		}()
	}

	// the sender
	go func() {
		defer func() {
			close(closed)
			close(dataCh)
		}()

		for {
			select {
			case <-closing:
				return
			default:
			}

			select {
			case <-closing:
				return
			case dataCh <- rand.Intn(Max):
			}
		}
	}()

	// receivers
	for i := 0; i < NumReceivers; i++ {
		go func() {
			defer wgReceivers.Done()

			for value := range dataCh {
				log.Println(value)
			}
		}()
	}

	wgReceivers.Wait()
}
