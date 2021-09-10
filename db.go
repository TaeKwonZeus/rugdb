package rugdb

type DB struct {
	file *dataFile
}

func Open(path string) (*DB, error) {
	dataFile, err := NewDataFile(path)
	if err != nil {
		return nil, err
	}

	return &DB{file: dataFile}, nil
}
