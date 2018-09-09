package main

import (
	"flag"

	conf "github.com/BlockchainLSB/pinocchio/config"
	"github.com/BlockchainLSB/stove/config"
	"github.com/BlockchainLSB/stove/server"
)

func main() {
	f := flag.String("f", "config.yml", "Config file")
	c := &conf.PinocchioConfig{}
	_ = config.LoadConfig(*f, c)

	s := server.New()
	s.Port = c.Servers[0].Port
	s.Static("/", "public")
	s.File("/", "public/index.html")

	_ = s.Start()
}
