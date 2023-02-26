package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	HTTPPort int
	HTTPHost string

	PostgresHost     string
	PostgresPort     int
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string

	SecretKey string

	PostgresMaxConnections  int
	PostgresConnMaxIdleTime int // in minutes

	RedisHost       string
	RedisPort       int
	RedisDB         int
	RedisPassword   string
	RedisPoolSize   int
	RedisExpiryTime int

	AccessTokenName string
	AccessTokenTTL  int // in minutes

	RefreshTokenName string
	RefreshTokenTTL  int // in minutes

}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}

func NewConfig() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	cnf := Config{}
	cnf.HTTPPort = getOrReturnDefaultValue("HTTP_PORT", 8080).(int)
	cnf.HTTPHost = getOrReturnDefaultValue("HTTP_HOST", "localhost").(string)

	cnf.PostgresHost = getOrReturnDefaultValue("POSTGRES_HOST", "localhost").(string)
	cnf.PostgresPort = getOrReturnDefaultValue("POSTGRES_PORT", 5432).(int)
	cnf.PostgresUser = getOrReturnDefaultValue("POSTGRES_USER", "postgres").(string)
	cnf.PostgresDatabase = getOrReturnDefaultValue("POSTGRES_DATABASE", "postgres").(string)
	cnf.PostgresPassword = getOrReturnDefaultValue("POSTGRES_PASSWORD", "postgres").(string)
	cnf.PostgresMaxConnections = getOrReturnDefaultValue("POSTGRES_MAX_CONNECTIONS", 10).(int)
	cnf.PostgresConnMaxIdleTime = getOrReturnDefaultValue("POSTGRES_CONN_MAX_IDLE_TIME", 10).(int)

	cnf.SecretKey = getOrReturnDefaultValue("SECRET_KEY0", "secret").(string)

	cnf.RedisHost = getOrReturnDefaultValue("REDIS_HOST", "localhost").(string)
	cnf.RedisPort = getOrReturnDefaultValue("REDIS_PORT", 6379).(int)
	cnf.RedisDB = getOrReturnDefaultValue("REDIS_DB", 0).(int)
	cnf.RedisPassword = getOrReturnDefaultValue("REDIS_PASSWORD", "").(string)
	cnf.RedisPoolSize = getOrReturnDefaultValue("REDIS_POOL_SIZE", 10).(int)
	cnf.RedisExpiryTime = getOrReturnDefaultValue("REDIS_EXPIRY_TIME", 10).(int)

}
