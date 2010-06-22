package main

import "flag"

var ShowFPS bool
var CmdNum int

func init() {
	flag.BoolVar(&ShowFPS, "fps", false, "Put FPS to std every 5 secs.")
	flag.IntVar(&CmdNum, "cmd", -1, "Command number to run.")
}
