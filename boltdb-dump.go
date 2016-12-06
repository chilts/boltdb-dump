package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/boltdb/bolt"
)

func check(err error) {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(2)
	}
}

func dumpCursor(tx *bolt.Tx, c *bolt.Cursor, indent int) {
	for k, v := c.First(); k != nil; k, v = c.Next() {
		if v == nil {
			fmt.Printf(strings.Repeat("  ", indent-1)+"[%s]\n", k)
			newBucket := c.Bucket().Bucket(k)
			if newBucket == nil {
				// from the top-level cursor and not a cursor from a bucket
				newBucket = tx.Bucket(k)
			}
			newCursor := newBucket.Cursor()
			dumpCursor(tx, newCursor, indent+1)
		} else {
			fmt.Printf(strings.Repeat("  ", indent-1)+"%s\n", k)
			fmt.Printf(strings.Repeat("  ", indent-1)+"  %s\n", v)
		}
	}
}

// Dump everything in the database.
func dump(db *bolt.DB) error {
	return db.View(func(tx *bolt.Tx) error {
		c := tx.Cursor()
		dumpCursor(tx, c, 1)
		return nil
	})
}

func main() {
	// the first arg is the database file
	filename := os.Args[1]

	// open this file
	opts := bolt.Options{
		ReadOnly: true,
		Timeout:  1 * time.Second,
	}
	db, err := bolt.Open(filename, 0600, &opts)
	check(err)
	defer db.Close()

	// dump all keys
	dump(db)
}
