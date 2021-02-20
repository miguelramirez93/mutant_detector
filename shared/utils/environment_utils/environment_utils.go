package environmentutils

import (
	"fmt"
	"os"
)

// GetOsEnv get environment variable from OS, if it is not set will panic.
func GetOsEnv(envVarName string) string {
	envVar, ok := os.LookupEnv(envVarName)
	if !ok {
		formatedMessage := fmt.Sprintf("Missing %s env var", envVarName)
		panic(formatedMessage)
	}
	return envVar
}
