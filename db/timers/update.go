package timers

import (
	"time"
)

func StartWeekUpdate() int64 {
	now := time.Now()
	daysUntilMonday := (8 - int(now.Weekday())) % 7
	if daysUntilMonday == 0 {
		daysUntilMonday = 7
	}

	return now.AddDate(0, 0, daysUntilMonday).Sub(now).Milliseconds()
}
