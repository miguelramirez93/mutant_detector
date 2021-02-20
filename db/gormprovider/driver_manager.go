package gormprovider

import (
	"mutant_detector/db/contracts"
	"mutant_detector/db/gormprovider/drivers"
)

type gormDriverManager struct{}

// NewGormDriverManager returns a instance of driver manager for GORM orm.
func NewGormDriverManager() contracts.DbDriverManager {
	return &gormDriverManager{}
}

func (dm *gormDriverManager) GetProvider(driverName string) contracts.DbProvider {
	switch driverName {
	case "postgres":
		return drivers.NewGormPostgres()
	default:
		return nil
	}
}
