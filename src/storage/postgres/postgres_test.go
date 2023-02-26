package postgres_test

import (
	"log"
	"os"
	"testing"

	"github.com/SuyunovJasurbek/url_shorting/config"
	"github.com/SuyunovJasurbek/url_shorting/src/storage"
	"github.com/SuyunovJasurbek/url_shorting/src/storage/postgres"
)

var cfg config.Config = config.Config{
	PostgresHost:     "localhost",
	PostgresPort:     "5437",
	PostgresUser:     "postgres",
	PostgresPassword: "postgres",
	PostgresDatabase: "najottalim",
}

var strg storage.StorageI

func TestMain(m *testing.M) {
	psdb, err := postgres.NewPostgres(&cfg)
	if err != nil {
		log.Printf("Error while connecting to postgres_test: %v", err)
		panic(err)
	}
	strg = psdb
	os.Exit(m.Run())
}
