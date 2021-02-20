package config

import (
	"fmt"
	environmentutils "mutant_detector/shared/utils/environment_utils"
)

// PostgresGetConnectionString get connection string formated from env vars for postgresDB
func PostgresGetConnectionString() string {
	host := environmentutils.GetOsEnv("DB_HOST")
	user := environmentutils.GetOsEnv("DB_USER")
	pass := environmentutils.GetOsEnv("DB_PASS")
	dbName := environmentutils.GetOsEnv("DB_NAME")
	port := environmentutils.GetOsEnv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, pass, dbName, port)

	return dsn
}
