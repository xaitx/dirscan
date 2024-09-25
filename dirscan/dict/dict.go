package dict

import "os"

// Dict is a dictionary
type Dict struct {
	FilePaht string //file path
	ReadAll  bool   //Whether to load all dictionaries at once
}

// read all dictionary
func (d *Dict) ReadAllDict() {
	// read all dictionary
	os.OpenFile(d.FilePaht, os.O_RDONLY, 0666)
}
