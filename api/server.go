package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4/middleware"
	// "gorm.io/driver/mysql"
	// "gorm.io/gorm"
)

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

func login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Throws unauthorized error
	if username != "admin" || password != "admin" {
		return echo.ErrUnauthorized
	}

	// Set custom claims
	claims := &jwtCustomClaims{
		"Administrator",
		true,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

// func restricted(c echo.Context) error {
// 	user := c.Get("user").(*jwt.Token)
// 	claims := user.Claims.(*jwtCustomClaims)
// 	name := claims.Name
// 	return c.String(http.StatusOK, "Welcome "+name+"!")
// }

// type Book struct {
// 	gorm.Model
// 	name  string
// 	price uint
// 	isbn  string
// }

func main() {
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// Login route
	e.POST("/login", login)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "books")
	})

	e.GET("/books", getBooks)
	e.GET("/books/:id", getBook)

	r := e.Group("/admin")
	config := middleware.JWTConfig{
		Claims:     &jwtCustomClaims{},
		SigningKey: []byte("secret"),
	}
	r.Use(middleware.JWTWithConfig(config))
	// r.GET("", restricted)
	r.POST("/books", saveBook)
	r.PUT("/books/:id", updateBook)
	r.DELETE("/books/:id", deleteBook)

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
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	token := claims.Name
	name := c.FormValue("name")
	price := c.FormValue("price")
	isbn := c.FormValue("isbn")
	return c.String(http.StatusOK, "token: "+token+", name: "+name+", price: "+price+", isbn: "+isbn)
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
