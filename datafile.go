package rugdb

import "os"

type dataFile struct {
	path string
	file *os.File
}

func NewDataFile(path string) (*dataFile, error) {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}

	return &dataFile{
		path: path,
		file: file,
	}, err
}
