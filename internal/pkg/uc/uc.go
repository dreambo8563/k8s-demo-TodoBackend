package uc

import (
	"encoding/json"
	"fmt"

	bolt "go.etcd.io/bbolt"
	"vincent.com/todo/internal/domain/model"
)

//DB - global instantce
var db *bolt.DB

//USERBUCKET -
const USERBUCKET = "User"

//Client -
type Client struct {
	DB *bolt.DB
}

//NewDB -
func NewDB() *Client {
	fmt.Println("NewDB", db)
	if db != nil {
		return &Client{
			DB: db,
		}
	}
	var err error
	db, err = bolt.Open("uc.db", 0666, nil)
	if err != nil {
		fmt.Println("NewDB err", err)
		return nil
	}
	return &Client{
		DB: db,
	}
}

//Destroy - Destroy the instantce
func Destroy() {
	if db != nil {
		db.Close()
	}
}

// Save -
func (c *Client) Save(u *model.User) error {
	return c.DB.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(USERBUCKET))
		if err != nil {
			return err
		}
		buf, err := json.Marshal(u)
		if err != nil {
			return err
		}
		err = b.Put([]byte(u.ID), []byte(buf))
		if err != nil {
			return err
		}
		return nil
	})
}

// Get -
func (c *Client) Get(u *model.User) error {
	return c.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(USERBUCKET))
		v := b.Get([]byte(u.ID))
		err := json.Unmarshal(v, u)
		if err != nil {
			return err
		}
		return nil
	})
}
