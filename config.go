package fibervhost

import (
	"github.com/gofiber/fiber/v2"
)

type Config struct {
	Next                 func(c *fiber.Ctx) bool
	Hostname             string
	Handler              func(c *fiber.Ctx) error
	HostnameRegexpString string
}

var ConfigDefault = Config{
	Next:     nil,
	Hostname: "vhost.local",
	Handler: func(c *fiber.Ctx) error {
		return nil
	},
	HostnameRegexpString: "",
}

func configDefault(config ...Config) Config {
	if len(config) < 1 {
		return ConfigDefault
	}

	cfg := config[0]

	if cfg.Next == nil {
		cfg.Next = ConfigDefault.Next
	}
	if cfg.Hostname == "" {
		cfg.Hostname = ConfigDefault.Hostname
	}
	if cfg.Handler == nil {
		cfg.Handler = ConfigDefault.Handler
	}
	if cfg.HostnameRegexpString == "" {
		cfg.HostnameRegexpString = ConfigDefault.HostnameRegexpString
	}

	return cfg
}
