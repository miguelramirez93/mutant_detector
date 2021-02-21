package main

import (
	"mutant_detector/db"
	server "mutant_detector/server/http"
	queueutils "mutant_detector/shared/utils/queue_utils"
)

func main() {
	dbProvider := "gorm"
	providerDriver := "postgres"
	queueutils.WaitForJobQueue()
	connInstance, err := db.GetDBProviderConnInstance(dbProvider, providerDriver)
	if err != nil {
		panic(err.Error())
	}
	db.ProviderMigrate(dbProvider, providerDriver, connInstance)
	server.HTTPInitServer(connInstance)
}
