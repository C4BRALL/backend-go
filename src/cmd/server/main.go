package main

import (
	"fmt"

	"github.com/backend/src/configs"
)

func main() {
	config, _ := configs.LoadConfig("../../../")
	fmt.Println(config)
}
