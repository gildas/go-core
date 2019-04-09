package core

import (
	"time"
)

// ExecEvery runs a func as a go routine regularly
// returns 3 chans:
// - a chan to stop the time loop, pipe true
// - a chan to force the job execution on demand
// - a chan to change the ticker interval
// the invoked func can modify the interval by piping a new Duration to its given chan
// the func is executed once before the ticker starts
func ExecEvery(job func(tick int64, at time.Time, changeme chan time.Duration), every time.Duration) (chan bool, chan bool, chan time.Duration) {
	changeme := make(chan time.Duration)
	pingme   := make(chan bool)
	stopme   := make(chan bool)

	go func() {
		tick := int64(0)
		go job(tick, time.Now(), changeme)
		ticker := time.NewTicker(every)
		for {
			select {
			case newtick := <-changeme:
				if newtick > 0 * time.Second {
					ticker.Stop()
					ticker = time.NewTicker(newtick)
				}
			case <- pingme:
				go job(tick, time.Now(), changeme)
			case now := <- ticker.C:
				tick++
				go job(tick, now, changeme)
			case <- stopme:
				ticker.Stop()
				return
			}
		}
	}()
	return stopme, pingme, changeme
}