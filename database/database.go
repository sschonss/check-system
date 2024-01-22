package database

import (
	"encoding/json"
	"log"

	"main/models"

	"github.com/boltdb/bolt"
)

var bucketName = []byte("system_info")

// InitDB initializes the BoltDB database.
func InitDB(dbPath string) (*bolt.DB, error) {
	db, err := bolt.Open(dbPath, 0600, nil)
	if err != nil {
		return nil, err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(bucketName)
		return err
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}

// SaveSystemInfo saves system information to the database.
func SaveSystemInfo(db *bolt.DB, data models.SystemInfo) {
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)
		encoded, err := json.Marshal(data)
		if err != nil {
			return err
		}

		return b.Put([]byte(data.Timestamp), encoded)
	})
	if err != nil {
		log.Printf("Error saving to database: %v", err)
	}
}
