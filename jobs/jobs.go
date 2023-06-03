package jobs

import (
	"time"

	"github.com/davidkt99/RoyalReaderBE/db"
	"github.com/go-co-op/gocron"
)

func CronJobsSetup() {
	s := gocron.NewScheduler(time.UTC)

	s.Every(1).Sunday().At("6:59").Do(func() {
		println("At 6:59 TruncateWeeklyChapterUpdates")
		db.TruncateWeeklyChapterUpdates()
	})

	s.Every(1).Day().At("7:00").Do(func() {
		println("At 7:00 NewChaptersJob")
		NewChaptersJob()
	})

	s.Every(1).Day().At("19:00").Do(func() {
		println("At 19:00 NewChaptersJob")
		NewChaptersJob()
	})

	//* for testing cronjobs
	// s.Every(10).Second().Do(func() {
	// 	NewChaptersJob()
	// })

	// s.Every(10).Second().Do(func() {
	// 	db.TruncateWeeklyChapterUpdates()
	// })

	s.StartAsync()
}
