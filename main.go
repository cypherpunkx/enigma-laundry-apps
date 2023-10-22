package main

import (
	"fmt"

	"enigmacamp.com/enigma-laundry-apps/config"
	"enigmacamp.com/enigma-laundry-apps/delivery"
)

func init() {
	config.InitiliazeConfig()
}

func main() {
	fmt.Println(config.Cfg)
	fmt.Println(config.Cfg.Token)
	fmt.Println(config.TokenCfg)
	fmt.Println(config.TokenCfg)
	delivery.Server().Run()
}
