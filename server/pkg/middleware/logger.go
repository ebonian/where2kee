package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func RequestLogger(logger *logrus.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Locals("logger", logger)

		start := time.Now()
		c.Next()

		latency := time.Since(start)
		statusCode := c.Response().StatusCode()
		method := c.Method()
		path := c.Path()

		logger.Infof("%s - %s %s - %d - %v", c.IP(), method, path, statusCode, latency)
		return nil
	}
}
