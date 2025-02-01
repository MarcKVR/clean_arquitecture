package config

import "github.com/spf13/viper"

type Config struct {
	ServerPort            string
	ExchangeAPIURL        string
	TimeZone              string
	DBDriver              string
	DBHost                string
	DBPort                string
	DBUser                string
	DBPassword            string
	DBName                string
	ExchangeRateAPIAccess string
	AWSAccessKeyID        string
	AWSSecretAccessKey    string
	AWSRegion             string
	S3BucketName          string
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	return &Config{
		ServerPort:            viper.GetString("SERVER_PORT"),
		ExchangeAPIURL:        viper.GetString("EXCHANGE_API_URL"),
		TimeZone:              viper.GetString("TIMEZONE"),
		DBDriver:              viper.GetString("DB_DRIVER"),
		DBHost:                viper.GetString("DB_HOST"),
		DBPort:                viper.GetString("DB_PORT"),
		DBUser:                viper.GetString("DB_USER"),
		DBPassword:            viper.GetString("DB_PASSWORD"),
		DBName:                viper.GetString("DB_NAME"),
		ExchangeRateAPIAccess: viper.GetString("EXCHANGERATE_API_ACCESS"),
		AWSAccessKeyID:        viper.GetString("AWS_ACCESS_KEY_ID"),
		AWSSecretAccessKey:    viper.GetString("AWS_SECRET_ACCESS_KEY"),
		AWSRegion:             viper.GetString("AWS_REGION"),
		S3BucketName:          viper.GetString("S3_BUCKET_NAME"),
	}, nil
}
