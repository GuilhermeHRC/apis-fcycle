package main

import "github.com/GuilhermeHRC/apis-fcycle/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}
