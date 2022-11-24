package main

import (
	"drone/config"
	"flag"
	"fmt"
	"log"

	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "user.yaml", "the config file")

func main() {
	fmt.Printf("Hello World!\n")
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	log.Printf("Starting userapi at %s:%d...\n", c.Host, c.Port)
}
