package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"encoding/xml"

	"github.com/labstack/echo"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Provide port to run on and server basename")
		os.Exit(1)
	}
	port := os.Args[1]

	e := echo.New()
	e.GET("/rss/*", func(c echo.Context) error {

		res, err := http.Get(fmt.Sprintf("https://zooqle.com%s?%s", c.Request().URL.Path, c.Request().URL.RawQuery))
		if err != nil {
			log.Fatal(err)
		}

		z := ZXML{}
		f := FXML{}
		zDec := xml.NewDecoder(res.Body)
		fEnc := xml.NewEncoder(c.Response())
		c.Response().Write([]byte(xml.Header))

		err = zDec.Decode(&z)
		if err != nil {
			return err
		}

		f.Channel.Items = make([]FItem, len(z.Channel.Items))
		for i, zi := range z.Channel.Items {
			f.Channel.Items[i] = FItem{
				Title: strings.Replace(zi.Title, "–", "-", -1),
				Link:  zi.Link,
				GUID:  strings.Replace(zi.GUID.URL, "https://zooqle.com", os.Args[2], -1),
			}
		}
		err = fEnc.Encode(&f)
		return err
	})
	e.GET("/download/*", func(c echo.Context) error {
		res, err := http.Get(fmt.Sprintf("https://zooqle.com%s?%s", c.Request().URL.Path, c.Request().URL.RawQuery))
		if err != nil {
			log.Fatal(err)
		}

		_, err = io.Copy(c.Response(), res.Body)
		return err
	})
	e.Logger.Fatal(e.Start(":" + port))
}
