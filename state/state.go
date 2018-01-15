package state

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/boltdb/bolt"
	"github.com/s1kx/unison/constant"
)

// DatabaseError acts as an adapter for bolt errors.
type DatabaseError error

// BucketError adds context to a transaction error.
type BucketError struct {
	Bucket string
	DatabaseError
}

func (e BucketError) Error() string {
	return fmt.Sprintf("bucket `%s`: %s", e.Bucket, e.DatabaseError)
}

// singleton pattern to handle a key value database for keeping track of bots current state relative to guild ID.
// https://github.com/boltdb/bolt
type singleton struct {
	db *bolt.DB
}

var instance *singleton
var once sync.Once

// DefaultState bot state for new newly added guilds
// var DefaultState = Normal

// GetDatabaseInstance get a bolt database instance.
// TODO: defer close?
func GetDatabaseInstance(file string) (*bolt.DB, error) {
	var err error
	once.Do(func() {
		db, e := bolt.Open(file, 0600, &bolt.Options{Timeout: 1 * time.Second})
		if e == nil {
			instance = &singleton{db: db}
		} else {
			err = e
		}
	})

	// bolt database instance.
	return instance.db, err
}

// GetInstance get a bolt database instance using the default database file
func GetInstance() (*bolt.DB, error) {
	return GetDatabaseInstance(constant.DefaultDatabaseFile)
}

// GetGuildValue retrieves a value using guildID and a key
// bucket == GuildID
func GetGuildValue(bucket, key string) ([]byte, error) {
	var val []byte
	db, err := GetInstance()
	if err != nil {
		return val, err
	}

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return &BucketError{bucket, bolt.ErrBucketNotFound}
		}

		val = b.Get([]byte(key))
		return nil
	})

	return val, err
}

// SetGuildValue sets a value using guildID, key and a value
// bucket == GuildID
func SetGuildValue(bucket, key string, val []byte) error {
	db, err := GetInstance()
	if err != nil {
		return err
	}

	if err := db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return &BucketError{bucket, err}
		}
		return b.Put([]byte(key), val)
	}); err != nil {
		return err
	}

	return nil
}

// GetGuildState returns the state of guild
func GetGuildState(guildID string /*discordgo uses strings for ID...*/) (Type, error) {
	val, err := GetGuildValue(guildID, constant.StateKey)
	if err != nil {
		return 0, err
	}

	// convert []byte into uint8 / aka state.Type
	i, err := strconv.ParseInt(string(val), 10, 16)
	if err != nil {
		return 0, err
	}

	return Type(i), nil
}

// SetGuildState updates the guild state in database
func SetGuildState(guildID string, state Type) error {
	err := SetGuildValue(guildID, constant.StateKey, []byte(ToStr(state)))

	return err
}
