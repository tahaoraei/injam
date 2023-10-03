package websocketserver

import (
	"context"
	"fmt"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/labstack/echo/v4"
	"injam/service/authservice"
	"log/slog"
)

func (h Handler) getLocation(c echo.Context) error {
	conn, _, _, err := ws.UpgradeHTTP(c.Request(), c.Response())
	if err != nil {
		return err
	}
	u := c.Get("my_claims").(*authservice.Claims)
	go func() {
		defer conn.Close()

		for {
			msg, _, err := wsutil.ReadClientData(conn)
			if err != nil {
			}
			if err = h.redisAdapter.Client().Publish(context.Background(), fmt.Sprintf("taha:%d", u.UserID), string(msg[:])).Err(); err != nil {
				slog.Error("cant save data in memory")
			}
		}
	}()
	return nil
}
