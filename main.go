package main

import (
	"flag"
	"fmt"

	"github.com/hytzongxuan/Codeforces-Hacker/app"
)

func version() {
	fmt.Println("Codeforces Hacker v0.3 by hytzongxuan")
}

var configFilePath string
var remoteServerURL string

func main() {
	version()

	flag.StringVar(&configFilePath, "c", "./config.json", "config file")
	flag.StringVar(&remoteServerURL, "s", "https://codeforces.com", "remote server URL")

	flag.Parse()

	app.Load(configFilePath, remoteServerURL)
}
