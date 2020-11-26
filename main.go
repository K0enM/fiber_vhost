package fibervhost

import (
	"github.com/gofiber/fiber/v2"
)

type Vhost struct {
	Host                 string
	Hostname             string
	HostnameRegexpString string
}

func New(config ...Config) fiber.Handler {

	cfg := configDefault(config...)

	return func(c *fiber.Ctx) error {
		if cfg.Next != nil && cfg.Next(c) {
			return c.Next()
		}

		if cfg.HostnameRegexpString != "" {
			re, _ := compile_regexp(cfg.HostnameRegexpString)
			if match(re, c.Hostname()) {
				vh := Vhost{
					Host:                 c.Hostname(),
					Hostname:             cfg.Hostname,
					HostnameRegexpString: cfg.HostnameRegexpString,
				}
				c.Locals("vhost", vh)
			}
		} else if re, _ := string_to_regexp(cfg.Hostname); match(re, c.Hostname()) {
			vh := Vhost{
				Host:                 c.Hostname(),
				Hostname:             cfg.Hostname,
				HostnameRegexpString: cfg.HostnameRegexpString,
			}
			c.Locals("vhost", vh)
		} else {
			return c.Next()
		}
		return cfg.Handler(c)
	}
}

func ToVhostStruct(val interface{}) Vhost {
	return val.(Vhost)
}
