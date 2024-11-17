package dirscan

import (
	"flag"
	"os"

	"github.com/xaitx/dirscan/config"
	"github.com/xaitx/logs"
)

func Parse() *config.Config {
	Config := config.Config{}
	var help bool
	flag.StringVar(&Config.Url, "u", "", "扫描的URL")
	flag.StringVar(&Config.UrlFile, "uf", "", "扫描的URL文件")
	flag.StringVar(&Config.Words, "w", "", "字典文件")
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
	// 检查字段
	if Config.Url == "" && Config.UrlFile == "" {
		logs.Info("url和urlfile不能同时为空")
		flag.Usage()
		os.Exit(0)
	}
	if Config.Words == "" {
		logs.Info("目录字典不能为空")
		flag.Usage()
		os.Exit(0)
	}

	return &Config
}
