package websocketserver

import (
	"context"
	"fmt"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/labstack/echo/v4"
	"log/slog"
	"strings"
)

func (h Handler) getLocation(c echo.Context) error {
	conn, _, _, err := ws.UpgradeHTTP(c.Request(), c.Response())
	if err != nil {
		return err
	}
	u := strings.Replace(c.Request().Header.Get("Authorization"), "Basic ", "", 1)

	go func() {
		defer conn.Close()

		for {
			msg, _, err := wsutil.ReadClientData(conn)
			if err != nil {
			}
			if err = h.redisAdapter.Client().Publish(context.Background(), fmt.Sprintf("taha:%s", u), string(msg[:])).Err(); err != nil {
				slog.Error("cant save data in memory")
			}
		}
	}()
	return nil
}
