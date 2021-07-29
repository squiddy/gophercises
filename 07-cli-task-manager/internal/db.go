package internal

import (
	"encoding/binary"
	"fmt"
	"time"

	"github.com/boltdb/bolt"
)

type Task struct {
	Created   time.Time
	Completed time.Time
	Title     string
}

func NewTask(title string) Task {
	return Task{
		Created:   time.Now(),
		Completed: time.Time{},
		Title:     title,
	}
}

func (t *Task) Complete() {
	t.Completed = time.Now()
}

func OpenDb(path string) (*bolt.DB, error) {
	db, err := bolt.Open("/tmp/my.db", 0600, nil)
	if err != nil {
		return nil, err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("todos"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}

// We store IDs as BigEndian because that gives us proper ordering of the keys.

func IDtoB(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
}

func BtoID(v []byte) uint64 {
	return binary.BigEndian.Uint64(v)
}
