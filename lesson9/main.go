package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/mswatermelon/lesson8/configuration"
)

func main() {
	config, err := configuration.LoadConfig("configuration.json")
	if err != nil {
		fmt.Println("Error happend during reading config:", err)
		os.Exit(0)
	}
	preparedConfig, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		fmt.Println("Error happend during printing config:", err)
		os.Exit(0)
	}
	fmt.Println(string(preparedConfig))
}
