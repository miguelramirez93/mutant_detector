package drivers

import (
	"errors"
	"mutant_detector/db/contracts"
	"mutant_detector/db/gormprovider/config"
	"mutant_detector/domain/dna"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type gormPostgres struct{}

// NewGormPostgres creates new instance of gorm orm with postgres driver
func NewGormPostgres() contracts.DbProvider {
	return &gormPostgres{}
}

func (gi *gormPostgres) Connect() (interface{}, error) {
	dsn := config.PostgresGetConnectionString()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}

func (gi *gormPostgres) Migrate(connInstance interface{}) error {
	db, ok := connInstance.(*gorm.DB)
	if !ok {
		return errors.New("conn-instance-parse-error")
	}

	db.AutoMigrate(&dna.DnaReport{})

	return nil
}
