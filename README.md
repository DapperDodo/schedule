# schedule

`run func X every Y at Z`

Usage:

go get -u github.com/DapperDodo/schedule


schedule examples:

	// run DoIt() every night at 03:00
	Run(EveryDay, OnHour, "03", DoIt)
	
	// run DoIt() every minute, 30 seconds into the minute
	Run(EveryMinute, OnSecond, "30", DoIt)
	
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
