# schedule

`run func X every Y at Z`

Usage:

	go get -u github.com/DapperDodo/schedule


schedule examples:

	// run hello() every 21st of the month
	Run(EveryMonth, OnDay, "21", hello)

	// run hello() every night at 03:00
	Run(EveryDay, OnHour, "03", hello)

	// run hello() every hour, at 15 minutes before the top of the hour
	Run(EveryHour, OnMinute, "45", hello)
	
	// run hello() every minute, 30 seconds into the minute
	Run(EveryMinute, OnSecond, "30", hello)
	
	Block()

example main.go:

	import (
		"fmt"
		
		"github.com/DapperDodo/schedule"
	)
	
	func main() {
	
		schedule.Run(schedule.EveryMinute, schedule.OnSecond, "00", hello)
	
		// start `serving`
		schedule.Block()
	}
	
	func hello() {
		fmt.Println("hello world!")
	}
