package middleware

import (
	"fmt"
	"runtime/debug"

	"github.com/YurcheuskiRadzivon/telemetrics-back/pkg/logger"
	"github.com/gofiber/fiber/v2"
	fiberRecover "github.com/gofiber/fiber/v2/middleware/recover"
)

func buildPanicMessage(c *fiber.Ctx, err interface{}) string {
	result := fmt.Sprintf("%s - %s %s PANIC DETECTED: %v\n%s\n",
		c.IP(),
		c.Method(),
		c.OriginalURL(),
		err,
		debug.Stack(),
	)

	return result
}

func logPanic(lgr *logger.Logger) func(c *fiber.Ctx, err interface{}) {
	return func(c *fiber.Ctx, err interface{}) {
		lgr.ErrorLogger.Println(buildPanicMessage(c, err))
	}
}

func Recovery(lgr *logger.Logger) func(c *fiber.Ctx) error {
	return fiberRecover.New(fiberRecover.Config{
		EnableStackTrace:  true,
		StackTraceHandler: logPanic(lgr),
	})
}
