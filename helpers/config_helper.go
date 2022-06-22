package helpers

import (
	"os"
	"path/filepath"
)

// GetDataDir -
func GetDataDir() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	folder := os.Getenv("DATA_DIR")
	if folder == "" {
		folder = "data"
	}

	mydir := filepath.Join(wd, folder)

	// Create the folder if it doesn't exist
	err = os.MkdirAll(mydir, os.ModePerm)
	if err != nil {
		return "", err
	}

	return mydir, nil
}

// GetDSN -
func GetDSN() string {
	dsn := os.Getenv("DSN")
	if dsn == "" {
		dir, _ := GetDataDir()
		dsn = filepath.Join(dir, "nid-native-registry.db")
	}

	return dsn
}
