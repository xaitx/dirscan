package main

import (
	"github.com/xaitx/dirscan/dirscan"
	"github.com/xaitx/logs"
)

func main() {
	logs.SetFlags(0)
	config := dirscan.Parse()
	err := dirscan.Start(config)
	if err != nil {
		logs.Fatal(err)
	}

}
