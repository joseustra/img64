package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/boltdb/bolt"
	"github.com/ustrajunior/minion"
)

// Context holds application context
type Context struct {
	db *bolt.DB
}

func openDB() *bolt.DB {
	db, err := bolt.Open(os.Getenv("DBNAME")+".db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := db.Begin(true)
	if err != nil {
		return nil
	}
	defer tx.Rollback()

	_, err = tx.CreateBucketIfNotExists([]byte("images"))
	if err != nil {
		log.Fatal(err)
	}

	if err = tx.Commit(); err != nil {
		log.Fatal(err)
	}

	return db
}

func main() {
	opts := minion.Options{
		Cors: []string{os.Getenv("CORS")},
		UnauthenticatedRoutes: []string{"*"},
	}
	m := minion.New(opts)

	ctx := &Context{db: openDB()}

	m.Get("/", ctx.ImageHandler)

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Panic("You need set a env var PORT with the desire port to run")
	}
	m.Run(port)
}

type Image struct {
	URL     string
	Content string
}

// ImageHandler grab the url download the image, convert to base64 and return
func (ctx *Context) ImageHandler(c *minion.Context) {
	image := &Image{URL: c.ByQuery("url")}
	if len(image.URL) == 0 {
		c.Text(http.StatusBadRequest, "")
		return
	}

	ctx.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("images"))
		v := b.Get([]byte(image.URL))
		if len(v) > 0 {
			image.Content = string(v)
		}
		return nil
	})

	if len(image.Content) == 0 {
		resp, err := http.Get(image.URL)

		if err != nil {
			c.Text(http.StatusBadRequest, err.Error())
			return
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err.Error())
		}

		image.Content = base64.StdEncoding.EncodeToString(body)

		ctx.db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("images"))
			err := b.Put([]byte(image.URL), []byte(image.Content))
			return err
		})
	}

	c.Text(http.StatusOK, image.Content)
}
