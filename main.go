package main

import (
	"flag"
	"fmt"
	"github.com/51idc/transfer/g"
	"github.com/51idc/transfer/http"
	"github.com/51idc/transfer/proc"
	"github.com/51idc/transfer/receiver"
	"github.com/51idc/transfer/sender"
	"os"
)

func main() {
	cfg := flag.String("c", "cfg.json", "configuration file")
	version := flag.Bool("v", false, "show version")
	versionGit := flag.Bool("vg", false, "show version")
	flag.Parse()

	if *version {
		fmt.Println(g.VERSION)
		os.Exit(0)
	}
	if *versionGit {
		fmt.Println(g.VERSION, g.COMMIT)
		os.Exit(0)
	}

	// global config
	g.ParseConfig(*cfg)

	//init seelog
	g.InitSeeLog()

	// proc
	proc.Start()

	sender.Start()
	receiver.Start()

	// http
	http.Start()

	select {}
}