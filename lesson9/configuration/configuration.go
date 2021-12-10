package configuration

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/joho/godotenv"
	"github.com/namsral/flag"
)

type Config struct {
	Port        int    `json:"Port" yaml:"Port"`
	DbUrl       string `json:"DbUrl" yaml:"DbUrl"`
	JaegerUrl   string `json:"JaegerUrl" yaml:"JaegerUrl"`
	SentryUrl   string `json:"SentryUrl" yaml:"SentryUrl"`
	KafkaBroker string `json:"KafkaBroker" yaml:"KafkaBroker"`
	AppId       string `json:"AppId" yaml:"AppId"`
	AppKey      string `json:"AppKey" yaml:"AppKey"`
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

func loadConfigFromEnv(fileName string) (config Config, err error) {
	if err = loadFile(fileName); err != nil {
		return config, err
	}

	if err := config.fillConfig(); err != nil {
		return config, fmt.Errorf("Error happened while loading config", err)
	}

	return config, nil
}

func loadConfigFromJson(fileName string) (config Config, err error) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return config, fmt.Errorf("Error happened while loading config", err)
	}

	if err = json.Unmarshal(content, &config); err != nil {
		return config, fmt.Errorf("Error happened while unmarshal config", err)
	}

	return config, nil
}

func loadConfigFromYaml(fileName string) (config Config, err error) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return config, fmt.Errorf("Error happened while loading config", err)
	}

	if err = yaml.Unmarshal(content, &config); err != nil {
		return config, fmt.Errorf("Error happened while unmarshal config", err)
	}

	return config, nil
}

func LoadConfig(fileName string) (config Config, err error) {
	if strings.Contains(fileName, "env") {
		return loadConfigFromEnv(fileName)
	}
	if strings.Contains(fileName, "json") {
		return loadConfigFromJson(fileName)
	}
	if strings.Contains(fileName, "yaml") {
		return loadConfigFromYaml(fileName)

	}

	return config, fmt.Errorf("Unknown file type")
}
