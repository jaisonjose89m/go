package concurrency

import (
	"log"
	"sync"
	"time"
)

func ScheduleInvoker() func() {
	ticker := time.NewTicker(5 * time.Second)
	done := make(chan struct{})
	cancel := func() {
		ticker.Stop()
		close(done)
	}

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				//log.Printf("Triggering at time: %v\n", t)
				log.Println("Scheduler commenced")
				go process(t.String())
				log.Println("Scheduler Completed")
			}
		}
	}()

	return cancel
}

func process(prefix string) {
	wg := sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		go func() {
			consume(prefix + "message" + (string(i)))
			wg.Done()
		}()
		wg.Add(1)
	}
	wg.Wait()
}

func consume(name string) {
	log.Printf("Consuming message: %s", name)
	time.Sleep(10 * time.Second)
	log.Printf("Consumed message: %s", name)
}
