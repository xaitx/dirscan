package dirscan

import (
	"flag"

	"github.com/xaitx/dirscan/config"
)

func Parse() *config.Config {
	Config := config.Config{}
	flag.StringVar(&Config.Url, "u", "", "target url")
	flag.StringVar(&Config.UrlFile, "uf", "", "target url file")
	flag.StringVar(&Config.Proxy, "p", "", "proxy")
	flag.StringVar(&Config.Method, "m", "GET", "request method")
	flag.IntVar(&Config.Timeout, "t", 10, "timeout")
	flag.IntVar(&Config.Threads, "n", 10, "number of threads")
	flag.StringVar(&Config.DictPath, "d", "", "dictionary path")
	flag.StringVar(&Config.Code, "c", "200", "status code")
	flag.StringVar(&Config.Length, "l", "", "length")
	flag.StringVar(&Config.Output, "o", "", "output file")
	flag.Parse()
	return &Config
}
