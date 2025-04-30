package main

import (
	"clinic-app/config"
	"clinic-app/controllers"
	"clinic-app/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	db := config.InitDB()

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		return controllers.Login(c)
	})
	app.Post("/register", func(c *fiber.Ctx) error {
		return controllers.Register(c)
	})

	api := app.Group("/api", middleware.JWTProtected())

	api.Post("/patients", middleware.RoleRequired("receptionist"), func(c *fiber.Ctx) error {
		return controllers.CreatePatient(c)
	})
	api.Get("/patients", middleware.RoleRequired("receptionist"), func(c *fiber.Ctx) error {
		return controllers.GetPatients(c)
	})
	api.Get("/patients/:id", middleware.RoleRequired("receptionist"), func(c *fiber.Ctx) error {
		return controllers.GetPatientByID(c)
	})
	api.Put("/patients/:id", middleware.RoleRequired("receptionist", "doctor"), func(c *fiber.Ctx) error {
		return controllers.UpdatePatient(c)
	})
	api.Delete("/patients/:id", middleware.RoleRequired("receptionist"), func(c *fiber.Ctx) error {
		return controllers.DeletePatient(c)
	})

	app.Listen(":3000")
}
