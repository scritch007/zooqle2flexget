package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"encoding/xml"

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

		z := ZXML{}
		f := FXML{}
		zDec := xml.NewDecoder(res.Body)
		fEnc := xml.NewEncoder(c.Response())

		err = zDec.Decode(&z)
		if err != nil {
			return err
		}

		f.Channel.Items = make([]FItem, len(z.Channel.Items))
		for i, zi := range z.Channel.Items {
			f.Channel.Items[i] = FItem{
				Title: zi.Title,
				Link:  zi.Link,
				GUID:  zi.GUID.URL,
			}
		}
		err = fEnc.Encode(&f)
		return err
	})
	e.Logger.Fatal(e.Start(":" + port))
}
