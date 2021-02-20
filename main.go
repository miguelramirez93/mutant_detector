package main

import (
	"mutant_detector/db"
	server "mutant_detector/server/http"
)

func main() {
	dbProvider := "gorm"
	providerDriver := "postgres"
	connInstance, err := db.GetDBProviderConnInstance(dbProvider, providerDriver)
	if err != nil {
		panic(err.Error())
	}
	db.ProviderMigrate(dbProvider, providerDriver, connInstance)
	server.HTTPInitServer(connInstance)
}
