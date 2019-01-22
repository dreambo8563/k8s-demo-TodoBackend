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
// TODO: need to seperate the model in db and service
func (c *Client) Save(u *model.User) error {
	return c.DB.Update(func(tx *bolt.Tx) error {
		b, err := getUserBucket(tx)
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

// GetByID -
// TODO: need to seperate the model in db and service
func (c *Client) GetByID(u *model.User) error {
	return c.DB.View(func(tx *bolt.Tx) error {
		b, err := getUserBucket(tx)
		if err != nil {
			return err
		}
		v := b.Get([]byte(u.ID))
		err = json.Unmarshal(v, u)
		if err != nil {
			return err
		}
		return nil
	})
}

// GetByName -
func (c *Client) GetByName(name string) (bool, error) {
	var exist bool
	if err := c.DB.Update(func(tx *bolt.Tx) error {
		b, err := getUserBucket(tx)
		if err != nil {
			return err
		}
		u := &model.User{}
		return b.ForEach(func(k, v []byte) error {
			err := json.Unmarshal(v, u)
			if err != nil {
				return err
			}
			if name == u.Name {
				exist = true
			}
			fmt.Println("GetByName", u)
			return nil
		})
	}); err != nil {
		return false, err
	}
	return exist, nil
}
func getUserBucket(tx *bolt.Tx) (*bolt.Bucket, error) {
	return tx.CreateBucketIfNotExists([]byte(USERBUCKET))
}
