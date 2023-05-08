package jobs

import (
	"time"

	"github.com/go-co-op/gocron"
)

func CronJobsSetup() {
	s := gocron.NewScheduler(time.UTC)

	// * for production
	s.Every(1).Day().At("7:00").Do(func() {
		NewChaptersJob()
	})

	//* for testing cronjobs
	// s.Every(10).Second().Do(func() {
	// 	NewChaptersJob()
	// })

	s.StartAsync()
}
