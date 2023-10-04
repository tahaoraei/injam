package websocketserver

import (
	"context"
	"fmt"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/labstack/echo/v4"
	"injam/service/authservice"
	"log/slog"
	"time"
)

func (h Handler) getLocation(c echo.Context) error {
	conn, _, _, err := ws.UpgradeHTTP(c.Request(), c.Response())
	if err != nil {
		return err
	}
	u := c.Get("my_claims").(*authservice.Claims)
	var locations []string
	stop := make(chan struct{})
	t := time.NewTicker(time.Second * 10)
	go func() {
		defer t.Stop()
		for {
			select {
			case <-t.C:
				fmt.Println("locations: ", locations)
				locations = nil
			case <-stop:
				fmt.Println("Here")
				t.Stop()
				break
			}
		}
	}()

	go func() {
		defer conn.Close()

		for {
			msg, _, err := wsutil.ReadClientData(conn)
			locations = append(locations, string(msg))
			fmt.Println("shity: ", time.Now(), string(msg))
			if err != nil {
				fmt.Println("websocket close")
				stop <- struct{}{}
				break
			}
			if err = h.redisAdapter.Client().Publish(context.Background(), fmt.Sprintf("taha:%d", u.UserID), string(msg[:])).Err(); err != nil {
				slog.Error("cant save data in memory")
			}
		}
	}()
	return nil
}
