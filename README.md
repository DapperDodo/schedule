# schedule

`run func X every Y at Z`

go get -u github.com/DapperDodo/schedule

// run DoIt() every night at 03:00
Run(EveryDay, OnHour, "03", DoIt)

// run DoIt() every minute, 30 seconds into the minute
Run(EveryMinute, OnSecond, "30", DoIt)

Block()



import (
	"github.com/DapperDodo/schedule"
)

func main() {

	schedule.Run(schedule.EveryMinute, schedule.OnSecond, "00", importer.Import)

	// start `serving`
	schedule.Block()
}
