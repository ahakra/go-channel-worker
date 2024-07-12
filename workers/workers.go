package workers

import (
	"fmt"
	"sync"
)

type Producer interface {
	ProduceMessage(message string)
}

type Workers struct {
	C             chan string
	Quit          chan int
	WG            *sync.WaitGroup
	Kafkaproducer Producer
	workersNeeded int
}

func NewWorkers(c chan string, quit chan int, k Producer, wg *sync.WaitGroup, workersNeeded int) *Workers {
	return &Workers{
		C:             c,
		Quit:          quit,
		WG:            wg,
		Kafkaproducer: k,
		workersNeeded: workersNeeded,
	}

}

func (w *Workers) worker(i int) {
	fmt.Printf("Wroker #%d started\n", i)
	defer w.WG.Done()
	for {
		select {
		case msg, ok := <-w.C:
			if ok {
				w.Kafkaproducer.ProduceMessage(msg)
			} else {
				fmt.Println("No more work, quitting.")
				return
			}
		case <-w.Quit:
			fmt.Println("quit")
			return
		}
	}
}

func (w *Workers) StartWorkers() {
	for i := 0; i < w.workersNeeded; i++ {
		w.WG.Add(1)
		go w.worker(i)
	}

}

func (w *Workers) WaitForWorkersToFinish() {
	close(w.Quit)
	close(w.C)
	w.WG.Wait()
}
