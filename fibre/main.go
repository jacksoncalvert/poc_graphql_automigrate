package main

import (
	"bitbucket.com/jcalvert/or-service/fibre/api"
	"bitbucket.com/jcalvert/or-service/fibre/psql_orm"

	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		res := api.Query()
		new_res, _ := json.Marshal(res)
		return c.SendString(string(new_res))
	})

	psql_orm.MakeGorm()
	app.Listen(":3000")
}
