package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type User struct {
	gorm.Model
	Username string
}

type Keijo struct {
	Kulli int
}

func main() {

	dsn := "host=192.168.37.133 user=postgres password=supersecretpassword dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("db failure")
	}

	db.AutoMigrate(&User{})

	u := User{
		Username: "kekkonen",
	}
	db.Create(&u)

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hi")
	})
	e.GET("/api/user/:id", func(c echo.Context) error {
		u := User{Username: c.Param("id")}
		db.Where(&u).First(&u)
		return c.String(http.StatusOK, strconv.FormatInt(int64(u.ID), 10))
	})

	e.POST("/api/user", func(c echo.Context) error {
		jsonBody := make(map[string]interface{})
		err := json.NewDecoder(c.Request().Body).Decode(&jsonBody)
		if err != nil {
			log.Error("error from POST /api/user")
			return c.String(http.StatusInternalServerError, "error")
		}
		js, err := json.Marshal(jsonBody)
		if err != nil {
			log.Error("error from POST /api/user")
			return c.String(http.StatusInternalServerError, "error")
		}
		var u User
		json.Unmarshal(js, &u)
		db.Create(&u)
		return c.String(http.StatusOK, u.Username)
	})

	e.Logger.Fatal(e.Start(":8080"))

}
