package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	// "gorm.io/driver/mysql"
	// "gorm.io/gorm"
)

// type Book struct {
// 	gorm.Model
// 	name  string
// 	price uint
// 	isbn  string
// }

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "books")
	})

	e.GET("/books", getBooks)
	e.GET("/books/:id", getBook)
	e.POST("/books", saveBook)
	e.PUT("/books/:id", updateBook)
	e.DELETE("/books/:id", deleteBook)
	e.Logger.Fatal(e.Start(":1323"))
}

func getBooks(c echo.Context) error {

	// dsn := "root@tcp(127.0.0.1:3306)/books?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// if err != nil {
	// 	panic("failed to connect database")
	// }

	// var book []Book
	// result := db.Find(&book)
	page := c.QueryParam("page")
	per_page := c.QueryParam("per_page")
	return c.String(http.StatusOK, "list buku page: "+page+", per_page: "+per_page)
}

func getBook(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "buku id: "+id)
}

func saveBook(c echo.Context) error {
	name := c.FormValue("name")
	price := c.FormValue("price")
	isbn := c.FormValue("isbn")
	return c.String(http.StatusOK, "name: "+name+", price: "+price+", isbn: "+isbn)
}

func updateBook(c echo.Context) error {
	id := c.Param("id")
	name := c.FormValue("name")
	price := c.FormValue("price")
	isbn := c.FormValue("isbn")
	return c.String(http.StatusOK, "id: "+id+", name: "+name+", price: "+price+", isbn: "+isbn)
}

func deleteBook(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "id: "+id)
}
