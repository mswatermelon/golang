package configuration

import (
	"fmt"
	"net/url"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/namsral/flag"
)

type Config struct {
	Port        int    `json:"Port"`
	DbUrl       string `json:"DbUrl"`
	JaegerUrl   string `json:"JaegerUrl"`
	SentryUrl   string `json:"SentryUrl"`
	KafkaBroker string `json:"KafkaBroker"`
	AppId       string `json:"AppId"`
	AppKey      string `json:"AppKey"`
}

func (config *Config) validate() error {
	if err := validateURI(config.DbUrl); err != nil {
		return fmt.Errorf("Error while parsing\"DbUrl\"", err)
	}
	if err := validateURI(config.JaegerUrl); err != nil {
		return fmt.Errorf("Error while parsing\"JaegerUrl\"", err)
	}
	if err := validateURI(config.SentryUrl); err != nil {
		return fmt.Errorf("Error while parsing\"SentryUrl\"", err)
	}
	if err := validatePort(config.Port); err != nil {
		return fmt.Errorf("Error while parsing \"Port\"", err)
	}

	return nil
}

func (config *Config) fillConfig() error {
	flag.IntVar(&config.Port, "port", 0, "Port")
	flag.StringVar(&config.DbUrl, "db_url", "", "Database URL")
	flag.StringVar(&config.JaegerUrl, "jaeger_url", "", "Jaeger URL")
	flag.StringVar(&config.SentryUrl, "sentry_url", "", "Sentry URL")
	flag.StringVar(&config.KafkaBroker, "kafka_broker", "", "Kafka broker")
	flag.StringVar(&config.AppId, "app_id", "", "Application identificator")
	flag.StringVar(&config.AppKey, "app_key", "", "Application key")
	flag.Parse()
	if err := config.validate(); err != nil {
		return err
	}

	return nil
}

func loadFile(fileName string) error {
	if err := godotenv.Load(fileName); err != nil {
		return fmt.Errorf("Error while loading file", fileName, err)
	}

	return nil
}

func validateURI(value string) error {
	if _, err := url.ParseRequestURI(value); err != nil {
		return err
	}

	return nil
}

func validatePort(value int) error {
	if value < 1 || value > 65535 {
		return fmt.Errorf("%s not within range of valid ports (1-65535)", value)
	}

	return nil
}

func LoadConfig(fileName string) (config Config, err error) {
	if err = loadFile(fileName); err != nil {
		return config, err
	}

	if err := config.fillConfig(); err != nil {
		fmt.Println("Error happened while loading config", err)
		return config, err
	}

	return config, nil
}

func LoadEnvFiles(pattern string) (config Config, err error) {
	matches, err := filepath.Glob(pattern)

	if err != nil {
		fmt.Println("Error while finding enironment files")
		return config, err
	}

	if err := godotenv.Overload(matches...); err != nil {
		fmt.Println("Error while reading enironment files")
		return config, err
	}

	if err := config.fillConfig(); err != nil {
		fmt.Println("Error happened while loading config", err)
		return config, err
	}

	return config, nil
}
