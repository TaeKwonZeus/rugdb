package rugdb

type dataFile struct {

}

func NewDataFile(path string) (*dataFile, error) {
	return &dataFile{}, nil
}