package configs

import (
	"github.com/cucumberjaye/gophermart/pkg/flags"
	"os"
)

var (
	SigningKey           string
	RunAddress           string
	DataBaseURI          string
	AccrualSystemAddress string
)

const defaultSigningKey = "qwerty1234"

func InitConfigs() error {
	flags.InitFlags()

	SigningKey = lookUpOrSetDefault("SIGNING_KEY", defaultSigningKey)
	RunAddress = lookUpOrSetDefault("RUN_ADDRESS", flags.RunAddress)
	DataBaseURI = lookUpOrSetDefault("DATABASE_URI", flags.DataBaseURI)
	AccrualSystemAddress = lookUpOrSetDefault("ACCRUAL_SYSTEM_ADDRESS", flags.AccrualSystemAddress)

	return nil
}

func lookUpOrSetDefault(name, defaultValue string) string {
	out, ok := os.LookupEnv(name)
	if !ok {
		return defaultValue
	}

	return out
}
