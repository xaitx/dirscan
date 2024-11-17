package dirscan

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/xaitx/dirscan/config"
	"github.com/xaitx/logs"
)

var (
	UserAgents = []string{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3",
		"Mozilla/5.0 (Windows NT 10.0; WOW64; rv:45.0) Gecko/20100101 Firefox/45.0",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/602.3.12 (KHTML, like Gecko) Version/10.1.2 Safari/602.3.12",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_1 like Mac OS X) AppleWebKit/603.1.30 (KHTML, like Gecko) Version/10.0 Mobile/14E304 Safari/602.1",
		"Mozilla/5.0 (Linux; Android 7.0; Nexus 5X Build/NBD90W) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Mobile Safari/537.36",
		"Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0; AS; rv:11.0) like Gecko",
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/44.0.2403.157 Safari/537.36",
		"Mozilla/5.0 (iPad; CPU OS 10_3_2 like Mac OS X) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.0 Mobile/14F89 Safari/602.1",
	}
)

// Scan函数
func Scan(url string, method string, timeout time.Duration, proxy string) (string, error) {
	// 根据UserAgents随机选择一个UserAgent
	userAgent := UserAgents[rand.Intn(len(UserAgents))]

	headers := map[string]string{
		"User-Agent": userAgent,
	}
	response, err := Request(url, method, headers, timeout, proxy)
	if err != nil {
		return "", err
	}
	// 判断相应状态码
	switch response.StatusCode {
	case 301, 302:
		redirecturl, err := GetRedirectUrl(url, response, proxy)
		if err != nil {
			return "", err
		}

		return logs.Magenta(fmt.Sprintf("[%d] %s -> %s", response.StatusCode, url, redirecturl)), nil
	case 403:
		return logs.Red(fmt.Sprintf("[%d] %s", response.StatusCode, url)), nil
	case 404:
		return "", fmt.Errorf("[%d] %s", response.StatusCode, url)
	default:
		// 处理其他状态码
		return logs.Green(fmt.Sprintf("[%d] %s", response.StatusCode, url)), nil
	}
}

// 获取重定向地址
func GetRedirectUrl(requestUrl string, response *http.Response, proxy string) (string, error) {
	// 获取重定向地址
	location := response.Header.Get("Location")
	if location != "" {
		// 处理location可能的情况，比如相对路径，绝对路径，其他链接跳转
		if location[0] == '/' {
			parsedRequestUrl, err := url.Parse(requestUrl)
			if err != nil {
				return "", err
			}
			// 拼接重定向地址
			redirectUrl := parsedRequestUrl.Scheme + "://" + parsedRequestUrl.Host + location
			return redirectUrl, nil
		} else if strings.HasPrefix(location, "http://") || strings.HasPrefix(location, "https://") {
			return location, nil
		} else {
			// 判断url结尾是/还是有文件名，如歌是/则直接拼接重定向，否则截取url最后一个/之前的部分再拼接
			if strings.HasSuffix(requestUrl, "/") {
				return requestUrl + location, nil
			} else {
				lastSlashIndex := strings.LastIndex(requestUrl, "/")
				return requestUrl[:lastSlashIndex] + location, nil
			}
		}
	}
	return "", errors.New("no location header")
}

// 开始扫描
func Start(cfg *config.Config) error {
	var wg sync.WaitGroup

	outputChan := make(chan string)
	dictChan := make(chan string)

	// 读取字典
	dict, err := ReadDictFile(cfg.Words)
	if err != nil {
		return err
	}

	// 创建线程池
	for i := 0; i < cfg.Threads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for word := range dictChan {
				url := cfg.Url + "/" + word
				str, err := Scan(url, "GET", time.Duration(cfg.Timeout)*time.Second, cfg.Proxy)
				if err != nil {
					continue
				}
				outputChan <- str
			}
		}()
	}
	go func() {
		for _, word := range dict {
			dictChan <- word
		}
		close(dictChan)
	}()

	go func() {
		wg.Wait()
		close(outputChan)
	}()

	for str := range outputChan {
		logs.Print(str)
	}
	return nil
}
