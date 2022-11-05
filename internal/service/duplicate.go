package service

import (
	"crypto/sha1"
	"encoding/base64"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Duplicate struct {
	FilsMap map[string]int
}

func New() *Duplicate {
	fileMap := make(map[string]int)
	return &Duplicate{
		FilsMap: fileMap,
	}
}

func (d *Duplicate) ReadFiles(path string) ([]byte, error) {
	dat, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return dat, nil
}

func (d *Duplicate) createHash(hashBytes []byte) string {
	hasher := sha1.New()
	hasher.Write(hashBytes)
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	return sha
}

func (d *Duplicate) GetDuplicates(dirPath string) (int, error) {
	duplicateCount := 0
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return 0, err
	}

	for _, file := range files {
		data, err := d.ReadFiles(filepath.Join(dirPath, file.Name()))
		if err != nil {
			return duplicateCount, err
		}

		hash := d.createHash(data)

		_, ok := d.FilsMap[hash]
		if ok {
			duplicateCount++
		} else {
			d.FilsMap[hash] = int(file.Size())
		}
	}

	return duplicateCount, nil
}
