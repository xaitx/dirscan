package dirscan_test

import (
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
	response, err := dirscan.Request("http://www.baidu.com", "GET", nil, "")
	if err != nil {
		t.Error(err)
	}
	logs.Print(response)
}
