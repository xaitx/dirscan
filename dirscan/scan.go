package dirscan

import (
	"github.com/xaitx/dirscan/config"
)

func Scan(config *config.Config) {
	// read dictionary
	dict, err := NewDict(config.DictPath, true)
	if err != nil {
		println(err.Error())
		return
	}
	defer dict.Close()
	// read dictionary
	for {
		line, err := dict.ReadLine()
		if err != nil {
			break
		}
		println(line)
	}
}
