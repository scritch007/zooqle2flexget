package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Provide port to run on")
		os.Exit(1)
	}
	port := os.Args[1]

	e := echo.New()
	e.GET("/*", func(c echo.Context) error {

		res, err := http.Get(fmt.Sprintf("https://zooqle.com%s?%s", c.Request().URL.Path, c.Request().URL.RawQuery))
		if err != nil {
			log.Fatal(err)
		}
		robots, err := ioutil.ReadAll(res.Body)

		c.String(200, string(robots))
		return err
	})
	e.Logger.Fatal(e.Start(":" + port))
}
