package rugdb

type DB struct {
	file     *dataFile
	settings Settings
}

func Open(path string) (*DB, error) {
	return OpenWithSettings(path, Settings{})
}

func OpenWithSettings(path string, settings Settings) (*DB, error) {
	dataFile, err := NewDataFile(path)
	if err != nil {
		return nil, err
	}

	settings = fillDefaults(Settings{})

	return &DB{
		file:     dataFile,
		settings: settings,
	}, nil
}
