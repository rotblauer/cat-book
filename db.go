package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/rotblauer/trackpoints/trackPoint"
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

// itob returns an 8-byte big endian representation of v.
func itob(v int64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func storeCat(cat CatPic) error {

	var err error

	go func() {
		GetDB().Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(trackKey))

			// id, _ := b.NextSequence()
			// trackPoint.ID = int(id)
			cat.TrackPoint.ID = cat.TrackPoint.Time.UnixNano() //dunno if can really get nanoy, or if will just *1000.

			if exists := b.Get(itob(cat.TrackPoint.ID)); exists != nil {
				// make sure it's ours
				var existingTrackpoint trackPoint.TrackPoint
				e := json.Unmarshal(exists, &existingTrackpoint)
				if e != nil {
					fmt.Println("Checking on an existing trackpoint and got an error with one of the existing trackpoints unmarshaling.")
				}
				if existingTrackpoint.Name == cat.TrackPoint.Name {
					fmt.Println("Got that trackpoint already. Breaking.")
					return nil
				}
			}

			trackPointJSON, err := json.Marshal(CatPic{})
			if err != nil {
				return err
			}
			err = b.Put(itob(cat.TrackPoint.ID), trackPointJSON)
			if err != nil {
				fmt.Println("Didn't save post trackPoint in bolt.", err)
				return err
			}
			fmt.Println("Saved trackpoint: ", cat.TrackPoint.ID)
			return nil
		})
	}()

	if err != nil {
		fmt.Println(err)
	}
	return err
}
