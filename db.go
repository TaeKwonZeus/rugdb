package rugdb

const magic = "RugDB"

type DB struct {
	file     *dataFile
	settings Settings
}

func Open(path string) (*DB, error) {
	return OpenWithSettings(path, Settings{})
}

func OpenWithSettings(path string, settings Settings) (*DB, error) {
	db := &DB{}
	var err error

	db.settings = fillDefaults(settings)

	db.file, err = newDataFile(path, db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (d *DB) Close() error {
	return d.file.close()
}
