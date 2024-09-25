package dirscan

import (
	"bufio"
	"os"
)

// Dict is a dictionary
type dict struct {
	FilePaht string   //file path
	ReadAll  bool     //Whether to load all dictionaries at once
	File     *os.File //
}

// NewDict creates a new dictionary
func NewDict(filePath string, readAll bool) (*dict, error) {
	d := &dict{
		FilePaht: filePath,
		ReadAll:  readAll,
	}
	var err error
	d.File, err = os.Open(filePath)
	if err != nil {
		return nil, err
	}
	return d, nil
}

// Close closes the dictionary
func (d *dict) Close() error {
	return d.File.Close()
}

// Read reads a dictionary
func (d *dict) ReadLine() (string, error) {
	buf := bufio.NewReader(d.File)
	data, _, err := buf.ReadLine()
	if err != nil {
		d.Close()
		return "", err
	}
	return string(data), nil
}
