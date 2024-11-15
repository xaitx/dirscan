package dirscan_test

import (
	"io"
	"testing"

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
	response, err := dirscan.Request("https://ifconfig.me/ip", "GET", nil, "")
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
	response, err = dirscan.Request("https://ifconfig.me/ip", "GET", nil, "socks5://127.0.0.1:7890")
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
