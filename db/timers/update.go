package timers

import (
	"time"
)

func StartWeekUpdate() int64 {
	now := time.Now().UnixNano() / int64(time.Millisecond)
	nowUnix := time.Unix(now/1000, 0).UTC()

	daysUntilNextMonday := (int(time.Monday) - int(nowUnix.Weekday()) + 7) % 7
	if daysUntilNextMonday == 0 {
		daysUntilNextMonday = 7
	}

	nextMonday := nowUnix.AddDate(0, 0, daysUntilNextMonday)
	nextMondayMidnight := time.Date(nextMonday.Year(), nextMonday.Month(), nextMonday.Day(), 0, 0, 0, 0, nextMonday.Location())

	return nextMondayMidnight.UnixMilli()
}
