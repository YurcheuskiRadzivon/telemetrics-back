package middleware

import (
	"fmt"
	"strconv"

	"github.com/YurcheuskiRadzivon/telemetrics-back/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

func buildRequestMessage(c *fiber.Ctx) string {
	result := fmt.Sprintf("%s - %s %s - %s %s",
		c.IP(),
		c.Method(),
		c.OriginalURL(),
		strconv.Itoa(c.Response().StatusCode()),
		strconv.Itoa(len(c.Response().Body())),
	)

	return result
}

func Logger(lgr *logger.Logger) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		err := c.Next()

		lgr.InfoLogger.Println(buildRequestMessage(c))

		return err
	}
}
