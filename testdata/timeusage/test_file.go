package timeusage

import (
	"time"
)

func timeUsageAnalyzerData() {
	time.Sleep(1 * time.Nanosecond) //want "time-Sleep"
	time.NewTimer(1 * time.Nanosecond)
}
