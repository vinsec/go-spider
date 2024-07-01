package util

import (
	"encoding/json"
	"errors"
	"os"
	"sync"
)

// depth of start seed
// 种子起始深度
const (
	SEED_START_DEPTH = 0
)

type seedData []string

func GetSeedFromFile(path string, isDownload bool) (map[string]bool, error) {
	if !IsFileExist(path) {
		return nil, errors.New("file not exist")
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	var decode seedData
	err = decoder.Decode(&decode)
	if err != nil {
		return nil, err
	}

	result := make(map[string]bool)
	var wg sync.WaitGroup
	for _, url := range decode {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			result[url] = isDownload
		}(url)
	}
	wg.Wait()

	return result, nil
}
