package util

import (
	"flag"
	"fmt"
	"os"
)

var Input = struct {
	Thread string
	Connections string
	Header string
	Usl string
	Dtime string
}{}

const (
	HelpInfo = `
 helpInfo
	-t Threads
	-c Connections
	-H Header
	-usl Usl
	-d DelayTime
	for example
	  ./mySender -usl www.baidu.com  -t 10 -c 10 -d 10
	`
)

func init() {
	help := false
	flag.BoolVar(&help, "help", false, "help Info")
	flag.StringVar(&Input.Thread, "t", "", "thread")
	flag.StringVar(&Input.Connections, "c", "", "Connections")
	flag.StringVar(&Input.Header, "H", "", "Header")
	flag.StringVar(&Input.Usl, "usl", "", "Usl")
	flag.StringVar(&Input.Dtime, "d", "", "DelayTime")
	flag.Parse()

	if help || Input.Usl == ""{
		fmt.Println(HelpInfo)
		os.Exit(0)
	}
}
