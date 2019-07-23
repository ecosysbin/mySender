package util

import (
	"flag"
	"fmt"
)

var Input = struct {
	threads string
	dtime   string
}{}

var HelpInfo = `
	-s threads
	-d delay time
	./mySender -s 1 -d 1
`

func init() {
	fmt.Println("init")
	flag.Parse()
	// for test flag args
	args := flag.Args()
	// for test range
	for _, s := range args {
		fmt.Println(s)
	}
	if len(args) == 2 {
		Input.threads = args[0]
		Input.dtime = args[1]
	}
	fmt.Println(Input)
}
