package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

const (
	ErrorModel = "!!!Error"
	ErrorStyle = "-->"
)

type Config struct {
	ProjectHost string
	ProjectPort string
	ServiceName string
	ProjectMode string // debug, test, release
	Version     string

	PostgresHost           string
	PostgresPort           string
	PostgresUser           string
	PostgresPassword       string
	PostgresDatabase       string
	PostgresSSLMode        string
	PostgresMaxConnections int32
	PostgresURL            string

	RefreshTokenSecret      string
	RefreshTokenExpiryHour  int
	AccessTokenSecret       string
	AccessTokenExpiryMinute int

	PasswordSalt string
	Environment  string

	AWSAccessKeyId     string
	AWSEndpointUrlS3   string
	AWSEndpointUrlIAM  string
	AWSRegion          string
	AWSSecretAccessKey string
	AWSImageBucketName string
	AWSVideoBucketName string
	AWSFileBucketName  string
	AWSRetryMax        int
}

// Load ...
func Load() *Config {
	envFileName := cast.ToString(getOrReturnDefaultValue("ENV_FILE_PATH", ".env"))

	if err := godotenv.Load(envFileName); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.ServiceName = cast.ToString(getOrReturnDefaultValue("SERVICE_NAME", "lnk"))
	config.ProjectMode = cast.ToString(getOrReturnDefaultValue("PROJECT_MODE", ProjectModeDevelopment))
	config.PasswordSalt = cast.ToString(getOrReturnDefaultValue("PASSWORD_SALT", "secret"))
	config.ProjectHost = cast.ToString(getOrReturnDefaultValue("PROJECT_HOST", "0.0.0.0"))
	config.ProjectPort = cast.ToString(getOrReturnDefaultValue("PROJECT_PORT", "8080"))
	config.Version = cast.ToString(getOrReturnDefaultValue("VERSION", "1.0.1"))
	config.Environment = cast.ToString(getOrReturnDefaultValue("ENVIRONMENT", ProjectModeDevelopment))

	config.PostgresHost = cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST", "0.0.0.0"))
	config.PostgresPort = cast.ToString(getOrReturnDefaultValue("POSTGRES_PORT", "5432"))
	config.PostgresUser = cast.ToString(getOrReturnDefaultValue("POSTGRES_USER", "lnk"))
	config.PostgresPassword = cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD", "secret123"))
	config.PostgresSSLMode = cast.ToString(getOrReturnDefaultValue("POSTGRES_SSL_MODE", "disable"))
	config.PostgresDatabase = cast.ToString(getOrReturnDefaultValue("POSTGRES_DATABASE", "lnk"))
	config.PostgresMaxConnections = cast.ToInt32(getOrReturnDefaultValue("POSTGRES_MAX_CONNECTIONS", 50))
	config.PostgresURL = cast.ToString(getOrReturnDefaultValue("POSTGRES_URL", "smth"))

	config.RefreshTokenSecret = cast.ToString(getOrReturnDefaultValue("REFRESH_TOKEN_SECRET", "dfb344389hf834ht49483u9r3u49503u590u34059i34905iu9t8h9430t943jt"))
	config.RefreshTokenExpiryHour = cast.ToInt(getOrReturnDefaultValue("REFRESH_TOKEN_EXPIRY_HOUR", 120))
	config.AccessTokenSecret = cast.ToString(getOrReturnDefaultValue("ACCESS_TOKEN_SECRET", "uh34980j09t89345y3748ty3o48958f3uk4590348590u4tijg49gk304okr43"))
	config.AccessTokenExpiryMinute = cast.ToInt(getOrReturnDefaultValue("ACCESS_TOKEN_EXPIRY_MINUTE", 15))

	config.AWSAccessKeyId = cast.ToString(getOrReturnDefaultValue("AWS_ACCESS_KEY_ID", "123"))
	config.AWSEndpointUrlS3 = cast.ToString(getOrReturnDefaultValue("AWS_ENDPOINT_URL_S3", "123"))
	config.AWSEndpointUrlIAM = cast.ToString(getOrReturnDefaultValue("AWS_ENDPOINT_URL_IAM", "123"))
	config.AWSRegion = cast.ToString(getOrReturnDefaultValue("AWS_REGION", "123"))
	config.AWSSecretAccessKey = cast.ToString(getOrReturnDefaultValue("AWS_SECRET_ACCESS_KEY", "123"))
	config.AWSImageBucketName = cast.ToString(getOrReturnDefaultValue("AWS_IMAGE_BUCKET_NAME", "image"))
	config.AWSVideoBucketName = cast.ToString(getOrReturnDefaultValue("AWS_VIDEO_BUCKET_NAME", "videos"))
	config.AWSFileBucketName = cast.ToString(getOrReturnDefaultValue("AWS_FILE_BUCKET_NAME", "files"))
	config.AWSRetryMax = cast.ToInt(getOrReturnDefaultValue("AWS_RETRY_MAX", 5))

	return &config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
