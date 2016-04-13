# schedule

`run func X every Y at Z`

Schedule is a REALLY REALLY SIMPLE scheduler for automating jobs.
Simply wrap your job inside a func(){}
Schedule the runs
Block until shutdown


installation:

	go get -u github.com/DapperDodo/schedule


basic example main.go:

	import (
		"fmt"
		
		"github.com/DapperDodo/schedule"
	)
	
	func main() {

		// schedule one or more runs	
		schedule.Run(schedule.EveryMinute, schedule.OnSecond, "00", hello)
	
		// start `serving`
		schedule.Block()
	}
	
	func hello() {
		fmt.Println("hello world!")
	}


variations:

	// run hello() every 21st of the month
	schedule.Run(EveryMonth, OnDay, "21", hello)

	// run hello() every night at 03:00
	schedule.Run(EveryDay, OnHour, "03", hello)

	// run hello() every hour, at 15 minutes before the top of the hour
	schedule.Run(EveryHour, OnMinute, "45", hello)
	
	// run hello() every minute, 30 seconds into the minute
	schedule.Run(EveryMinute, OnSecond, "30", hello)
