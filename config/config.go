package config

type Config struct {
	Url     string // 扫描的URL
	UrlFile string // 扫描的URL文件
	Words   string // 字典文件
	Threads int    // 线程数
	Timeout int    // 超时时间
	Proxy   string // 代理
}
