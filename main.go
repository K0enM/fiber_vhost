package fibervhost

import (
	"github.com/gofiber/fiber/v2"
)

type Vhost struct {
	Host string
	Hostname string
	Length int
}

func New(config ...Config) func(c *fiber.Ctx) error {
	
	cfg := configDefault(config...)
	

	return func(c *fiber.Ctx) error {
		if cfg.Next != nil && cfg.Next(c) {
			return c.Next()
		}
		
		if c.Hostname() == cfg.Hostname {
			vh := Vhost{
				Host: c.Hostname(),
				Hostname: cfg.Hostname,
				Length: len(cfg.Hostname),
			}
			c.Locals("vhost", vh)
		} else {
			return c.Next()
		}

		cfg.Handler(c)
		return nil
	}
}
