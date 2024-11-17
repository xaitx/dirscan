package dirscan

import (
	"context"
	"errors"
	"net"
	"net/http"
	"net/url"
	"time"

	p "golang.org/x/net/proxy"
)

// 请求函数
func Request(url string, method string, headers map[string]string, timout time.Duration, proxy string) (*http.Response, error) {
	// 创建请求
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	// 设置请求头
	for k, v := range headers {
		request.Header.Set(k, v)
	}
	// 创建客户端
	client := &http.Client{
		Timeout: timout,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	client.Transport = &http.Transport{}

	// 设置代理
	if proxy != "" {
		err = SetProxy(client, proxy)
		if err != nil {
			return nil, err
		}
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// 代理处理函数
func SetProxy(client *http.Client, proxy string) error {
	parsedUrl, err := url.Parse(proxy)
	if err != nil {
		return err
	}
	if parsedUrl.Scheme == "socks5" {
		dialer, err := p.SOCKS5("tcp", parsedUrl.Host, nil, p.Direct)
		if err != nil {
			return err
		}
		transport := client.Transport.(*http.Transport)
		transport.DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
			return dialer.Dial(network, addr)
		}
	} else if parsedUrl.Scheme == "http" || parsedUrl.Scheme == "https" {
		transport := client.Transport.(*http.Transport)
		transport.Proxy = http.ProxyURL(parsedUrl)
	} else {
		return errors.New("unknown proxy scheme")
	}
	return nil
}
