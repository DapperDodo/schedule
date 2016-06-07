package schedule

import (
	"fmt"
	"time"
)

const (
	EveryMonth  = "2006-01"
	EveryDay    = "2006-01-02"
	EveryHour   = "2006-01-02T15"
	EveryMinute = "2006-01-02T15:04"

	OnDay    = "02"
	OnHour   = "15"
	OnMinute = "04"
	OnSecond = "05"
)

/*
	// run DoIt() every night at 03:00
	Run(EveryDay, OnHour, "03", DoIt)

	// run DoIt() every minute, 30 seconds into the minute
	Run(EveryMinute, OnSecond, "30", DoIt)

	// start `serving`
	Block()
*/
func Run(periodformat string, momentformat string, moment string, action func()) {

	last := time.Now().Format(periodformat)
	fmt.Println("\nScheduler : `last` initialized to", last)

	go func() {

		for {

			time.Sleep(time.Second)

			t := time.Now()

			// skip while in last period
			curr := t.Format(periodformat)
			if curr == last {
				continue
			}

			// skip until we reach the right moment
			if moment != t.Format(momentformat) {
				continue
			}

			// go go go!
			last = curr
			action()
		}
	}()
}

func Block() {
	select {}
}
