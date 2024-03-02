package db

import (
	"encoding/binary"
	"time"

	"github.com/boltdb/bolt"
)

var taskBucket = []byte("tasks")

var db *bolt.DB

type Task struct {
	Id    int
	Value string
}

type BoltDB struct {
	*bolt.DB
}

func InitDB(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}

	// opening a rw tx ;all tx are processed int he closure(kinda like callbacks in js) inside the func args
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})
}

func (bd *BoltDB) CreateTask(task string) (int, error) {
	var id uint64
	err := bd.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		id64, _ := b.NextSequence()
		id = id64
		key := itob(int(id64))
		err := b.Put(key, []byte(task))
		return err
	})
	if err != nil {
		return -1, err
	}
	return int(id), err

}

func (bd *BoltDB) AllTasks() ([]Task, error) {
	var tasks []Task
	err := bd.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket(taskBucket)

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			id := btoi(k)
			tasks = append(tasks, Task{
				Id:    id,
				Value: string(v),
			})
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (bd *BoltDB) DeleteTask(key int) error {
	return bd.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		err := b.Delete(itob(key))
		return err
	})
}

// util func
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}

// getters for DB
func GetDB() Database {
	if db == nil {
		panic("database not initialized")
	}
	return &BoltDB{db}
}
