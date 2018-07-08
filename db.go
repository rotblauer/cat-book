package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"os"
	"path"
)

var (
	db         *bolt.DB
	trackKey   = "cats"
	allBuckets = []string{trackKey, "cats"}
)

// GetDB is db getter.
func GetDB() *bolt.DB {
	return db
}

func initBuckets(buckets []string) error {
	err := GetDB().Update(func(tx *bolt.Tx) error {
		var e error
		for _, buck := range buckets {
			_, e = tx.CreateBucketIfNotExists([]byte(buck))
			if e != nil {
				return e
			}
		}
		return e
	})
	return err
}

// InitBoltDB sets up initial stuff, like the file and necesary buckets
func InitBoltDB() error {
	var err error
	os.Mkdir("db", 0777)
	db, err = bolt.Open(path.Join("db", "cats.db"), 0666, nil)
	if err != nil {
		fmt.Println("Could not initialize Bolt database. ", err)
		return err
	}

	err = initBuckets(allBuckets)
	if err != nil {
		fmt.Println("Err initing buckets.", err)
	}
	return err
}
