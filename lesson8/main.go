package main

import (
	"fmt"
	"os"

	"github.com/mswatermelon/lesson8/configuration"
)

func main() {
	config, err := configuration.LoadConfig("configuration.env")
	if err != nil {
		fmt.Println("Error happend during reading config:", err)
		os.Exit(0)
	}
	fmt.Printf("%+v\n", config)
}
