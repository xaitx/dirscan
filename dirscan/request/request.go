package request

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/proxy"
)

type Request struct {
	Method string // request method
	Proxy  string // proxy
}

// send request
func (r *Request) Request(url string) (*http.Response, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	// add proxy
	err := r.formatProxy(tr)
	if err != nil {
		return nil, err
	}
	// create client
	client := &http.Client{Transport: tr}
	// creare request
	req, err := http.NewRequest(r.Method, url, nil)
	if err != nil {
		return nil, err
	}
	// send request
	return client.Do(req)
}

// format proxy
func (r *Request) formatProxy(tr *http.Transport) error {
	if r.Proxy == "" {
		return nil
	} else if strings.HasPrefix(r.Proxy, "http://") {
		urlProxy, err := url.Parse(r.Proxy)
		if err != nil {
			return err
		}
		tr.Proxy = http.ProxyURL(urlProxy)
	} else if strings.HasPrefix(r.Proxy, "socks5://") {
		dialIp := r.Proxy[10:]
		dialer, err := proxy.SOCKS5("tcp", dialIp, nil, proxy.Direct)
		if err != nil {
			return err
		}
		perHost := proxy.NewPerHost(dialer, nil)
		// tr.Dial = dialer.Dial
		tr.DialContext = perHost.DialContext

	}
	return nil
}
