package core

import (
	"time"
)

// ExecEvery runs a func as a go routine regularly
// returns a chan, pipe "true" to stop the time loop
// the invoked func can modify the interval by piping a new Duration to its given chan
// the func is executed once before the ticker starts
func ExecEvery(job func(tick int64, at time.Time, changeme chan time.Duration), every time.Duration) chan bool {
	changeme := make(chan time.Duration)
	stopme   := make(chan bool)

	go func() {
		tick   := int64(0)
		go job(tick, time.Now(), changeme)
		ticker := time.NewTicker(every)
		for {
			select {
			case newtick := <-changeme:
					ticker.Stop()
					ticker = time.NewTicker(newtick)
			case now := <- ticker.C:
				tick++
				go job(tick, now, changeme)
			case <- stopme:
				ticker.Stop()
				return
			}
		}
	}()
	return stopme
}