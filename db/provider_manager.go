package db

import (
	"errors"
	"mutant_detector/db/contracts"
	gormprovider "mutant_detector/db/gormprovider"
)

// GetDBProviderConnInstance get the provider connection instance by provider and driver name.
func GetDBProviderConnInstance(providerName string, driverName string) (interface{}, error) {
	var provider contracts.DbProvider
	switch providerName {
	case "gorm":
		provider = gormprovider.NewGormDriverManager().GetProvider(driverName)
	default:
		return nil, errors.New("db-provider-not-found")
	}

	return provider.Connect()
}

// ProviderMigrate do the priovider right migration to DB.
func ProviderMigrate(providerName, driverName string, connInstance interface{}) error {
	var provider contracts.DbProvider
	switch providerName {
	case "gorm":
		provider = gormprovider.NewGormDriverManager().GetProvider(driverName)
	default:
		return errors.New("db-provider-not-found")
	}

	return provider.Migrate(connInstance)
}
