package config

// Parameter configuration
type Config struct {
	Url      string // target url
	UrlFile  string // target url file
	Proxy    string // proxy
	Method   string // request method
	Timeout  int    // timeout
	Threads  int    // number of threads
	DictPath string // dictionary path
	Code     string // status code
	Length   string // length
	Output   string // output file
}
