package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// set tampilan html ke dalam router
	router.LoadHTMLGlob("./templates/*")

	var library []Book

	router.GET("/", func(c *gin.Context) {
		// kembalikan tampilan index.html
		c.HTML(http.StatusOK, "index.html", gin.H{"library": library})
	})

	router.POST("/add", func(c *gin.Context) {
		// kerjakan createBook() dan kembalikan nilai book & err
		book, err := createBook(c.PostForm("title"), c.PostForm("author"), c.PostForm("isbn"))

		// cek apakah terdapat error
		if err != nil {
			// kembalikan status http 400 apabila terdapat error
			c.HTML(http.StatusBadRequest, "error.html", gin.H{"error": err.Error()})
			return
		}

		// tambahkan book ke dalam library apabila tidak terdapat error
		library = append(library, book)

		// arahkan kembali ke halaman awal
		c.Redirect(http.StatusSeeOther, "/")
	})

	// jalankan router pada port 8080
	router.Run(":8080")
}

// struktur objek Book
type Book struct {
	Title  string
	Author string
	ISBN   string
}

func validateISBN(isbn string) error {
	// cek apakah isbn terdiri dari 13 karakter
	if len(isbn) != 13 {
		// kembalikan string error apabila tidak memenuhi syarat
		return errors.New("ISBN must be 13 characters long")
	}

	// kembalikan error nil (mungkin berarti null) apabila memenuhi syarat
	return nil
}

func createBook(title, author, isbn string) (Book, error) {
	// kerjakan pengecekan nilai isbn dan kembalikan nilai err
	err := validateISBN(isbn)

	// cek apakah terdapat error
	if err != nil {
		// kembalikan book kosong dan err apabila terdapat error
		return Book{}, err
	}

	// kembalikan book baru apabila tidak terdapat error
	return Book{Title: title, Author: author, ISBN: isbn}, nil
}
