package configuration

import (
	"fmt"
	"net/url"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/namsral/flag"
)

type Config struct {
	port        int
	dbUrl       string
	jaegerUrl   string
	sentryUrl   string
	kafkaBroker string
	appId       string
	appKey      string
}

func loadFile(fileName string) error {
	if err := godotenv.Load(fileName); err != nil {
		fmt.Println("Error while loading file", fileName)
		return err
	}

	return nil
}

func validateURI(value string) error {
	if _, err := url.ParseRequestURI(value); err != nil {
		return err
	}

	return nil
}

func validate(config *Config) {
	if err := validateURI(config.dbUrl); err != nil {
		fmt.Println("Error while parsing\"", err, "\"")
	}
	if err := validateURI(config.jaegerUrl); err != nil {
		fmt.Println("Error while parsing\"", err, "\"")
	}
	if err := validateURI(config.sentryUrl); err != nil {
		fmt.Println("Error while parsing\"", err, "\"")
	}

	for _, b := range strconv.Itoa(config.port) {
		if b < '0' || b > '9' {
			fmt.Println("Error while parsing \"port\"")
		}
	}
}

func parseConfig(config *Config) {
	flag.IntVar(&config.port, "port", 0, "Port")
	flag.StringVar(&config.dbUrl, "db_url", "", "Database URL")
	flag.StringVar(&config.jaegerUrl, "jaeger_url", "", "Jaeger URL")
	flag.StringVar(&config.sentryUrl, "sentry_url", "", "Sentry URL")
	flag.StringVar(&config.kafkaBroker, "kafka_broker", "", "Kafka broker")
	flag.StringVar(&config.appId, "app_id", "", "Application identificator")
	flag.StringVar(&config.appKey, "app_key", "", "Application key")
	flag.Parse()
	validate(config)
}

func LoadConfig(fileName string) (config Config, err error) {
	if err = loadFile(fileName); err != nil {
		return config, err
	}

	parseConfig(&config)

	return config, nil
}

func LoadEnvFiles(fileName string) (config Config, err error) {
	if err = loadFile(fileName); err != nil {
		return config, err
	}

	matches, err := filepath.Glob("*.env")

	if err != nil {
		fmt.Println("Error while finding enironment files")
		return config, err
	}

	if err := godotenv.Overload(matches...); err != nil {
		fmt.Println("Error while reading enironment files")
		return config, err
	}

	parseConfig(&config)

	return config, nil
}
