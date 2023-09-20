package main

import (
	"log"
	fiberSwagger "github.com/swaggo/fiber-swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/test-fiber/database"
	"github.com/test-fiber/routes"
	_ "github.com/test-fiber/docs"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to my awesome API")
}


func SetupRoutes() *fiber.App{
	app:= fiber.New()
	app.Get("/swagger/*",fiberSwagger.WrapHandler)

	app.Get("/",welcome)
	app.Post("/api/user/signin",routes.CreateUser)
	app.Get("/api/users",routes.GetUsers)
	app.Get("/api/user/:id",routes.GetUser)
	app.Put("/api/user/update/:id",routes.UpdateUser)
	app.Delete("/api/user/unsubscribe/:id",routes.DeleteUser)

	log.Fatal(app.Listen(":3000"))
	return app
}


// @title Fiber Swagger Example API
// @version 2.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /
// @schemes http
func main() {
	database.ConnectDB()	
	SetupRoutes()
}

