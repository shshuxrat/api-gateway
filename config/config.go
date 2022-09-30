package config

import (
	"os"

	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	Environment string // develop, staging, production

	PositionServiceHost string
	PositionServicePort int

	ProfessionServiceHost string
	ProfessionServicePort int

	MinioAccessKeyID string
	MinioSecretKey   string
	MinioEndpoint    string
	MinioBucketName  string
	MinioLocation    string
	MinioHost        string

	LogLevel string
	HttpPort string
}

// Load loads environment vars and inflates Config
func Load() Config {
	config := Config{}

	config.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))

	config.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	config.HttpPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8091"))

	config.PositionServiceHost = cast.ToString(getOrReturnDefault("POSITION_SERVICE_HOST", "localhost"))
	config.PositionServicePort = cast.ToInt(getOrReturnDefault("POSITION_SERVICE_PORT", 9102))

	config.ProfessionServiceHost = cast.ToString(getOrReturnDefault("PROFESSION_SERVICE_HOST", "localhost"))
	config.ProfessionServicePort = cast.ToInt(getOrReturnDefault("PROFESSION_SERVICE_PORT", 9103))

	/*config.MinioEndpoint = cast.ToString(getOrReturnDefault("MINIO_ENDPOINT", "test.cdn.urecruit.udevs.io"))
	config.MinioAccessKeyID = cast.ToString(getOrReturnDefault("MINIO_ACCESS_KEY_ID", "2R5YabYDYwesXPDPprWc6DpbczCsXL97"))
	config.MinioSecretKey = cast.ToString(getOrReturnDefault("MINIO_SECRET_KEY_ID", "Ps5Che6XtJ6JmvsFXrXUH3tnhxwnZNYh"))
	config.MinioBucketName = cast.ToString(getOrReturnDefault("MINIO_BACKET_NAME", "photos"))
	config.MinioLocation = cast.ToString(getOrReturnDefault("MINIO_LOCATION", "us-east-1"))
	config.MinioHost = cast.ToString(getOrReturnDefault("MINIO_HOST", "test.cdn.urecruit.udevs.io"))
	*/
	return config
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
