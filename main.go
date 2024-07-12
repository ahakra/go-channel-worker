package main

import (
	"log"
	"runtime"
	"sync"

	filereader "go-channel-worker/file-reader"
	"go-channel-worker/kafka"
	"go-channel-worker/workers"
)

func main() {

	runtime.GOMAXPROCS(1)

	//-------------
	location := "file.txt"
	urls := []string{"0.0.0.0:9092"}
	topic := "kafkaTopic"
	workersNeeded := 50

	//---------------------------
	var wg sync.WaitGroup

	c := make(chan string)
	quit := make(chan int)

	//-------------------------

	kf, err := kafka.NewKafkaProducer(urls, topic)
	if err != nil {
		log.Fatal(err)
	}

	workers := workers.NewWorkers(c, quit, kf, &wg, workersNeeded)
	workers.StartWorkers()

	newfileReader := filereader.NewFileReader(location, c)
	newfileReader.InitReader()
	err = newfileReader.ReadFileToChannel()
	if err != nil {
		log.Fatal(err)
	}

	workers.WaitForWorkersToFinish()

}
