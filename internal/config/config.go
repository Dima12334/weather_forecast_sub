package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
)

const (
	defaultHTTPPort       = "8080"
	defaultMigrationsPath = "file://migrations"
)

type (
	Config struct {
		HTTP       HTTPConfig
		Logger     LoggerConfig
		DB         DatabaseConfig
		ThirdParty ThirdPartyConfig
		SMTP       SMTPConfig
		Email      EmailConfig
	}

	HTTPConfig struct {
		Host   string `mapstructure:"host"`
		Port   string `mapstructure:"port"`
		Domain string
	}

	ThirdPartyConfig struct {
		WeatherAPIKey string
	}

	LoggerConfig struct {
		LoggerEnv string
	}

	DatabaseConfig struct {
		Host           string
		Port           string
		User           string
		Password       string
		DBName         string
		SSLMode        string
		DSN            string
		MigrationsPath string `mapstructure:"migrationsPath"`
	}

	SMTPConfig struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		From     string `mapstructure:"from"`
		FromName string `mapstructure:"from_name"`
		Pass     string
	}

	EmailConfig struct {
		Templates EmailTemplates
		Subjects  EmailSubjects
	}

	EmailTemplates struct {
		Confirmation          string `mapstructure:"confirmation_email"`
		WeatherForecastDaily  string `mapstructure:"weather_forecast_daily"`
		WeatherForecastHourly string `mapstructure:"weather_forecast_hourly"`
	}

	EmailSubjects struct {
		Confirmation    string `mapstructure:"confirmation_email"`
		WeatherForecast string `mapstructure:"weather_forecast"`
	}
)

func Init(configDir string) (*Config, error) {
	populateDefaults()

	if err := parseConfigFile(configDir); err != nil {
		return nil, err
	}

	var cfg Config
	if err := unmarshalConfig(&cfg); err != nil {
		return nil, err
	}

	setFormEnv(&cfg)

	cfg.HTTP.Domain = fmt.Sprintf("%s:%s", cfg.HTTP.Host, cfg.HTTP.Port)

	cfg.DB.DSN = fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.DBName,
		cfg.DB.SSLMode,
	)

	return &cfg, nil
}

func unmarshalConfig(cfg *Config) error {
	if err := viper.UnmarshalKey("http_server", &cfg.HTTP); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("db", &cfg.DB); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("smtp", &cfg.SMTP); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("email", &cfg.Email); err != nil {
		return err
	}
	return nil
}

func parseConfigFile(configDir string) error {
	viper.SetConfigName("main")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configDir)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

func setFormEnv(cfg *Config) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file")
	}

	cfg.Logger.LoggerEnv = os.Getenv("LOGG_ENV")

	cfg.ThirdParty.WeatherAPIKey = os.Getenv("WEATHER_API_KEY")

	cfg.SMTP.Pass = os.Getenv("SMTP_PASSWORD")

	cfg.DB.Host = os.Getenv("DB_HOST")
	cfg.DB.Port = os.Getenv("DB_PORT")
	cfg.DB.User = os.Getenv("DB_USER")
	cfg.DB.Password = os.Getenv("DB_PASSWORD")
	cfg.DB.DBName = os.Getenv("DB_NAME")
	cfg.DB.SSLMode = os.Getenv("DB_SSLMODE")
}

func populateDefaults() {
	viper.SetDefault("http_server.port", defaultHTTPPort)
	viper.SetDefault("db.migrationsPath", defaultMigrationsPath)
}
