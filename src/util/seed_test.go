package util

import "testing"

func TestGetSeedFromFile(t *testing.T) {
	dummyPath := "./seedFile"
	_, err := GetSeedFromFile(dummyPath, false)
	if err == nil || err.Error() != "file not exist" {
		t.Error("test GetSeedFromFile() failed")
	}

	realPath := "../../data/seed"
	data, err := GetSeedFromFile(realPath, false)

	var isSeedExist bool
	for u, _ := range data {
		if u == "https://china.caixin.com/" {
			isSeedExist = true
			break
		}
	}
	if !isSeedExist {
		t.Error("test GetSeedFromFile() failed ")
	}

}
