package main

import (
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func hello(c echo.Context) error {
	conn, _, _, err := ws.UpgradeHTTP(c.Request(), c.Response())

	if err != nil {
		// handle error
		return err
	}
	go func() {
		defer conn.Close()

		for {
			msg, op, err := wsutil.ReadClientData(conn)
			if err != nil {
				// handle error
			}
			u := c.Request().Header.Get("Authorization")
			//for key, values := range c.Request().Header {
			//	fmt.Println(key)
			//	for _, value := range values {
			//		fmt.Println(value)
			//	}
			//}
			msg1 := string(msg[:]) + u
			err = wsutil.WriteServerMessage(conn, op, []byte(msg1))
			if err != nil {
				// handle error
			}
		}
	}()
	return nil
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/ws", hello)
	e.Logger.Fatal(e.Start(":8089"))
}
