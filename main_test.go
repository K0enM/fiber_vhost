package fibervhost

import (
	"github.com/gofiber/fiber/v2"
	"net/http/httptest"
	"testing"
)

// go test -run Test_Vhost_Match
func Test_Vhost_Match(t *testing.T) {
	want := "example.com"

	app := fiber.New()
	app.Use(New(Config{
		Hostname: want,
		Handler: func(c *fiber.Ctx) error {
			got := ToVhostStruct(c.Locals("vhost"))
			if !(got.Host == want) {
				t.Error("Error: incorrect match, host does not match hostname")
			}
			return nil
	}, }))
	
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("test")
	})

	req := httptest.NewRequest("GET", "http://" + want, nil)
	app.Test(req)

}

// go test -run Test_Vhost_No_Match
func Test_Vhost_No_Match(t *testing.T) {
	want := "test.com"

	app := fiber.New()
	app.Use(New(Config{
		Hostname: want,
		Handler: func(c *fiber.Ctx) error {
			t.Error("Error: match occurred with different host & hostname")
			return nil
		},
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("test")
	})

	req := httptest.NewRequest("GET", "http://example.com", nil)
	app.Test(req)
}

// go test -run Test_VHost_Next_Skip
func Test_VHost_Next_Skip(t *testing.T) {
	want := "example.com"

	app := fiber.New()
	app.Use(New(Config{
		Next: func(c *fiber.Ctx) bool {
			if c.Get("X-test-skip") == "yes" {
				return true
			} 

			return false
		},
		Hostname: want,
		Handler: func(c *fiber.Ctx) error {
			t.Error("Error: did not skip when Next returned true")
			return nil
	}, }))
	
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("test")
	})

	req := httptest.NewRequest("GET", "http://" + want, nil)
	req.Header.Add("X-test-skip", "yes")
	app.Test(req)
}
