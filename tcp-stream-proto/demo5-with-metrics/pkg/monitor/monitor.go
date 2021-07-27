package monitor

import (
	"expvar"
	"time"
)

var (
	SubmitInTotal *expvar.Int
	submitInRate  *expvar.Int
)

func init() {
	// register statistics index
	SubmitInTotal = expvar.NewInt("submitInTotal")
	submitInRate = expvar.NewInt("submitInRate")

	go func() {
		var lastSubmitInTotal int64

		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				newSubmitInTotal := SubmitInTotal.Value()
				submitInRate.Set(newSubmitInTotal - lastSubmitInTotal)
				lastSubmitInTotal = newSubmitInTotal
			}
		}
	}()
}
