package queueutils

import (
	"log"
	"time"
)

var jobQueue = make(chan func(), 1000)

// QueueJob queue jobs with a basic fifo queue
func QueueJob(job func()) {
	jobQueue <- job
}

// WaitForJobQueue init the job queue bufered chan
func WaitForJobQueue() {
	log.Println("Init job queue")
	go func() {
		for {
			timer1 := time.NewTimer(time.Duration(250) * time.Millisecond)
			<-timer1.C
			select {
			case job := <-jobQueue:
				job()
			}

		}
	}()
}
