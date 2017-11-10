package main

import (
	"flag"
	"fmt"
	"github.com/anchnet/transfer/g"
	"github.com/anchnet/transfer/http"
	"github.com/anchnet/transfer/proc"
	"github.com/anchnet/transfer/receiver"
	"github.com/anchnet/transfer/sender"
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

	if  !g.CheckMac(){
		fmt.Println("No move without permission")
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