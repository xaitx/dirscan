package dirscan_test

import (
	"io"
	"testing"
	"time"

	"github.com/xaitx/dirscan/dirscan"
	"github.com/xaitx/logs"
)

func TestReadDictFile(t *testing.T) {
	data, err := dirscan.ReadDictFile("../dict/test.txt")
	if err != nil {
		t.Error(err)
	}
	logs.Print(data)
}

// 测试请求函数
func TestRequest(t *testing.T) {
	logs.PrintColor(logs.Blue, "不使用代理请求")
	response, err := dirscan.Request("https://ifconfig.me/ip", "GET", nil, 5*time.Second, "")
	if err != nil {
		t.Error(err)
	}
	// 读取响应内容
	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
	}
	logs.Info(string(body))

	logs.PrintColor(logs.Blue, "使用代理请求")
	response, err = dirscan.Request("https://ifconfig.me/ip", "GET", nil, 5*time.Second, "socks5://127.0.0.1:7890")
	if err != nil {
		t.Error(err)
	}
	// 读取响应内容
	body, err = io.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
	}
	logs.Info(string(body))
}

// 测试扫描函数
func TestScan(t *testing.T) {
	str, err := dirscan.Scan("https://httpstat.us", "GET", 5*time.Second, "")
	if err != nil {
		logs.Error(err)
	}
	logs.Info(str)

	str, err = dirscan.Scan("https://httpstat.us/301", "GET", 5*time.Second, "")
	if err != nil {
		logs.Error(err)
	}
	logs.Info(str)

	str, err = dirscan.Scan("https://httpstat.us/302", "GET", 5*time.Second, "")
	if err != nil {
		logs.Error(err)
	}
	logs.Info(str)

	str, err = dirscan.Scan("https://httpstat.us/404", "GET", 5*time.Second, "")
	if err != nil {
		logs.Error(err)
	} else {
		logs.Info(str)
	}
	str, err = dirscan.Scan("https://httpstat.us/500", "GET", 5*time.Second, "")
	if err != nil {
		logs.Error(err)
	}
	logs.Info(str)

	str, err = dirscan.Scan("https://httpstat.us/503", "GET", 5*time.Second, "")
	if err != nil {
		logs.Error(err)
	}
	logs.Info(str)
}
