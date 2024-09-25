package main

import (
	"github.com/xaitx/dirscan/dirscan"
)

func main() {
	// Config := Parse()
	config := dirscan.Parse()
	dirscan.Scan(config)
}
