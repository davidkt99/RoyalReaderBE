package jobs

import (
	"time"

	"github.com/go-co-op/gocron"
)

func CronJobsSetup() {
	s := gocron.NewScheduler(time.UTC)

	// s.Every(1).Day().At("7:00").Do(func() {
	// 	NewChaptersJob()
	// })
	s.Every(10).Second().Do(func() {
		NewChaptersJob()
	})

	s.StartAsync()
}
