package main

import (
	"encoding/base64"
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
		UnauthenticatedRoutes: minion.AllRoutes,
	}
	m := minion.Classic(opts)

	ctx := &Context{db: openDB()}

	m.Get("/b", ctx.EncodedImageHandler)
	m.Get("/", ctx.ImageHandler)

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Panic("You need set a env var PORT with the desire port to run")
	}
	m.Run(port)
}

// Image representation of an image
type Image struct {
	URL     string
	Content string
}

// EncodedImageHandler grab the url download the image, convert to base64 and return
func (ctx *Context) EncodedImageHandler(c *minion.Context) {
	image := &Image{URL: c.Req.URL.Query().Get("url")}
	if len(image.URL) == 0 {
		c.Text(http.StatusBadRequest, "")
		return
	}

	err := getImageContent(ctx, image)
	if err != nil {
		c.Text(http.StatusBadRequest, "")
		return
	}

	c.Text(http.StatusOK, image.Content)
}

// ImageHandler grab the url download the image, convert to base64 to cache but returns the image file
func (ctx *Context) ImageHandler(c *minion.Context) {
	if len(c.Req.URL.Query().Get("url")) == 0 {
		c.Text(http.StatusBadRequest, "")
		return
	}

	image := &Image{URL: c.Req.URL.Query().Get("url")}
	err := getImageContent(ctx, image)
	if err != nil {
		c.Text(http.StatusBadRequest, "")
		return
	}

	decodedImg, err := base64.StdEncoding.DecodeString(image.Content)
	if err != nil {
		c.Text(http.StatusBadRequest, "")
	}

	c.Text(http.StatusOK, string(decodedImg))
}

func getImageContent(ctx *Context, image *Image) error {
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
			// c.Text(http.StatusBadRequest, err.Error())
			return err
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		image.Content = base64.StdEncoding.EncodeToString(body)

		ctx.db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("images"))
			err := b.Put([]byte(image.URL), []byte(image.Content))
			return err
		})
	}

	return nil
}
