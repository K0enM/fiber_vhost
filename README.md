# fiber_vhost
Vhost (Virtual host) middleware for [Fiber](https://github.com/gofiber/fiber) that enables the use of virtual hosts based on the Host Header.

### Table of Contents
- [Signatures](#signatures)
- [Examples](#examples)
- [Config](#config)
- [Default Config](#default-config)
- [TODO](#todo)

### Signatures
```go
func New(config ...Config) func(c *fiber.Ctx) error
```

## Examples
First ensure that the appropiate packages are imported
```go
import (
	"github.com/gofiber/fiber/v2"
	"github.com/K0enM/fiber_vhost"
)
```

#### **Initialization / Default Config**
```go
// Default middleware config
app.Use(fiber_vhost.New())
```

#### **Matching example.com and define basic Handler function**
```go
app.Use(fiber_vhost.New(fiber_vhost.Config{
	Hostname: "example.com",
	Handler: func(c *fiber.Ctx) error {
		return c.SendString("Inside the Vhost Handler")
	},
}))
```

#### **Define Next function to decide if to skip this middleware**
```go
app.Use(fiber_vhost.New(fiber_vhost.Config{
	Next: func(c *fiber.Ctx) bool {
		if c.Get("X-Skip-Vhost") == "true" {
			return true
		}

		return false
	},	
	Hostname: "example.com",
	Handler: func(c *fiber.Ctx) error {
		return c.SendString("Inside the Vhost Handler")
	},
}))
```

### Config
```go
type Config struct {
	Next func(c *fiber.Ctx) bool

	Hostname string

	Handler func(c *fiber.Ctx) error
}
```

### Default Config
```go
var ConfigDefault = Config{
	Next: nil,
	Hostname: "vhost.local",
	Handler: func(c *fiber.Ctx) error {
		return nil
	},
}
```

### TODO
- Comment the code
- Document the data added to `fiber.Ctx.Locals()`
- Match based on regex, not pure string match
