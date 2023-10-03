package websocketserver

import (
	"context"
	"fmt"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/labstack/echo/v4"
)

func (h Handler) shareLocation(c echo.Context) error {
	conn, _, _, err := ws.UpgradeHTTP(c.Request(), c.Response())
	if err != nil {
		return err
	}

	pub_id := c.Param("id")
	subscriber := h.redisAdapter.Client().Subscribe(context.Background(), fmt.Sprintf("taha%s", pub_id))

	go func() {
		defer conn.Close()

		for {

			// fetch from redis
			msg, err := subscriber.ReceiveMessage(context.Background())
			if err != nil {
				//slog.Error("problem in getting data from redis: ", err)
			}

			err = wsutil.WriteServerMessage(conn, ws.OpText, []byte(msg.Payload))
			if err != nil {
			}
		}
	}()
	return nil
}
