package queueutils

import "time"

// QueueJobWithTimmer queue jobs with a basic fifo queue drive by time intervals
func QueueJobWithTimmer(millisecondsInterval int, job func()) {
	timer1 := time.NewTimer(time.Duration(millisecondsInterval) * time.Millisecond)
	timer2 := time.NewTimer(time.Duration(millisecondsInterval) * time.Millisecond)
	<-timer2.C
	go func() {
		<-timer1.C
		job()
	}()
}
