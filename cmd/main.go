package main

import (
	"github.com/xaitx/dirscan/dirscan"
	"github.com/xaitx/logs"
)

func main() {
	config := dirscan.Parse()
	logs.Info("Start scanning...")
	logs.Print(config)
}
