package dirscan

import (
	"flag"
	"os"

	"github.com/xaitx/dirscan/config"
)

func Parse() *config.Config {
	Config := config.Config{}
	var help bool
	flag.StringVar(&Config.Url, "url", "", "扫描的URL")
	flag.StringVar(&Config.UrlFile, "urlfile", "", "扫描的URL文件")
	flag.StringVar(&Config.Words, "words", "", "字典文件")
	flag.IntVar(&Config.Threads, "threads", 10, "线程数")
	flag.IntVar(&Config.Timeout, "timeout", 5, "超时时间")
	flag.StringVar(&Config.Proxy, "proxy", "", "代理")
	flag.BoolVar(&help, "h", false, "帮助信息")
	flag.Usage = func() {
		flag.PrintDefaults()
	}
	flag.Parse()
	if help {
		flag.Usage()
		os.Exit(0)
	}
	return &Config
}
