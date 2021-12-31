package routes

import (
	"challenge_fullstack/src/controllers"
	"challenge_fullstack/src/utilities"
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

//load routes
func LoadRoutes() {
	app := fiber.New()

	app.Use(cors.New())

	// app.Get("/dashboard", monitor.New())

	api := app.Group("/api")

	v0 := api.Group("/v0")

	developers := v0.Group("/developers")

	developers.Post("", controllers.CreateDevelopers)
	developers.Put(utilities.ParamByID, controllers.UpdateDevelopers)
	developers.Delete(utilities.ParamByID, controllers.DeleteDevelopers)
	developers.Get("", controllers.QueryDevelopers)
	developers.Get(utilities.ParamByID, controllers.QueryByIDDevelopers)

	levels := v0.Group("/levels")

	levels.Post("", controllers.CreateLevels)
	levels.Put(utilities.ParamByID, controllers.UpdateLevels)
	levels.Delete(utilities.ParamByID, controllers.DeleteLevels)
	levels.Get("", controllers.QueryLevels)
	levels.Get(utilities.ParamByID, controllers.QueryByIDLevels)
	var err error

	port, err := strconv.Atoi(os.Getenv("PORT_API"))
	if err != nil {
		log.Fatal("Invalid API port: " + err.Error())
	}

	err = app.Listen(":" + strconv.Itoa(port))
	if err != nil {
		log.Fatal(err.Error())
	}

}
