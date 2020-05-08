package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

type Timestamp time.Time

func main() {
	e := echo.New()
	e.GET("/request", func(c echo.Context) error {
		req := c.Request()
		t := time.Now()

		format := `
			<code>
				Protocol: %s<br>
				Host: %s<br>
				Remote Address: %s<br>
				Method: %s<br>
				Path: %s<br>
				Time: %s<br>
			</code>
		`

		return c.HTML(http.StatusOK, fmt.Sprintf(format, req.Proto, req.Host, req.RemoteAddr, req.Method, req.URL.Path, t))
	})

	e.GET("/user", GetUser)

	e.Logger.Fatal(e.StartTLS(":1323", "cert.pem", "key.pem"))

}
