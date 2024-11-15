package dirscan

import (
	"bufio"
	"os"
)

// 读取字典文件
func ReadDictFile(filename string) ([]string, error) {
	file, error := os.Open(filename)
	if error != nil {
		return nil, error
	}
	defer file.Close()

	var dict []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		dict = append(dict, scanner.Text())
	}

	if error := scanner.Err(); error != nil {
		return nil, error
	}

	// 对dict去除空白字符和重复项
	dict = RemoveDuplicateAndEmpty(dict)
	return dict, nil
}

// 去除重复项和空白字符
func RemoveDuplicateAndEmpty(s []string) []string {
	m := make(map[string]bool)
	var r []string
	for _, v := range s {
		if v != "" && !m[v] {
			m[v] = true
			r = append(r, v)
		}
	}
	return r
}
