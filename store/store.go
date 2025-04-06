package store

import (
	// "errors"
	// "fmt"
	sqllightstore "photo-saver/store/sql-light-store"
	// "strings"
)

type DataStore interface {
	Open() error
	Close() error
	// SaveUser(username string) (int64, error)
	// SavePhoto(userID int64, filePath string) (int64, error)
	// GetPhotos(userID int64) ([]string, error)
}

func NewDataStore(config *Config) DataStore {
	// fmt.Println(config.DatabaseUrl)
	return sqllightstore.New(config.DatabaseUrl)

}
