package state

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/boltdb/bolt"
)

const stateKey = "state"
const defaultDatabaseFile = "unisonStates.db"

// singleton pattern to handle a key value database for keeping track of bots current state relative to guild ID.
// https://github.com/boltdb/bolt

type singleton struct {
	db *bolt.DB
}

var instance *singleton
var once sync.Once

// DefaultState bot state for new newly added guilds
var DefaultState = Normal

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
	return GetDatabaseInstance(defaultDatabaseFile)
}

// GetGuildValue retrieves a value using guildID and a key
// bucket == GuildID
func GetGuildValue(bucket, key string) ([]byte, error) {
	db, err := GetInstance()
	if err != nil {
		return nil, err
	}

	var val []byte
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return fmt.Errorf("Bucket " + bucket + " not found!")
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

	err = db.Update(func(tx *bolt.Tx) error {
		b, e := tx.CreateBucketIfNotExists([]byte(bucket))
		if e != nil {
			return fmt.Errorf("create bucket: %s", e)
		}
		return b.Put([]byte(key), val)
	})

	return err
}

// GetGuildState returns the state of guild
func GetGuildState(guildID string /*discordgo uses strings for ID...*/) (Type, error) {

	val, err := GetGuildValue(guildID, stateKey)
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
	err := SetGuildValue(guildID, stateKey, []byte(ToStr(state)))

	return err
}
