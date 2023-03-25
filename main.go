package main

import (
	"github.com/davidkt99/RoyalReaderBE/db"
	"github.com/davidkt99/RoyalReaderBE/jobs"
	"github.com/davidkt99/RoyalReaderBE/services"
	_ "github.com/lib/pq"
)

func main() {
	//*	Database Start
	db.DBSetup()
	defer db.DBShutDown()

	//* CronJobs Start
	jobs.CronJobsSetup()

	//*	Services Start
	services.StartServices()

}
