package agileful

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"testing"
)

const srvURL = "http://127.0.0.1:%d/article"


func runServer(ctx context.Context, t *testing.T, app *fiber.App, port int) {
	if err := app.Listen(fmt.Sprintf(":%d", port)); err != nil {
		t.Fatal(err)
	}
	<-ctx.Done()
	app.Shutdown()
}