package flags

import (
	"flag"
)

var (
	RunAddress           string
	DataBaseURI          string
	AccrualSystemAddress string
)

const (
	defaultServerAddress = "localhost:8000"
	defaultDatabaseURI   = "postgres://postgres:qwerty1234@localhost:5432/postgres"
)

func InitFlags() {
	flag.StringVar(&RunAddress, "a", defaultServerAddress, "server address to listen on")
	flag.StringVar(&DataBaseURI, "d", defaultDatabaseURI, "database URL")
	flag.StringVar(&AccrualSystemAddress, "r", "", "accrual system address")

	flag.Parse()
}
