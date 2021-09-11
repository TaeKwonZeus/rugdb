package rugdb

import (
	"bytes"
	"encoding/binary"
	"errors"
	"os"
)

type dataFile struct {
	db   *DB
	path string
	file *os.File
}

func (d *dataFile) close() error {
	return d.file.Close()
}

func (d *dataFile) getPage(id pgid) ([]byte, error) {
	buf := make([]byte, d.db.settings.PageSize)

	n, err := d.file.ReadAt(buf, int64(id)*int64(d.db.settings.PageSize))
	if err != nil {
		d.db.Close()
		return nil, err
	}
	if n != d.db.settings.PageSize {
		return nil, errors.New("invalid page")
	}

	return buf, nil
}

func (d *dataFile) valid() (bool, error) {
	meta, err := d.getPage(1)
	if err != nil {
		return false, err
	}

	mag := make([]byte, 5)
	binary.LittleEndian.PutUint32(mag, magic)

	if !bytes.Equal(meta[:6], mag) {
		return false, nil
	}

	return true, nil
}

func newDataFile(path string, db *DB) (*dataFile, error) {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}

	return &dataFile{
		db:   db,
		path: path,
		file: file,
	}, nil
}
