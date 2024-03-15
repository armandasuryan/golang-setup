package routes

import (
	"sygap-cmdb/controller"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, controller controller.Icontroller) {
	app.Get("api/sygap/cmdb/ciclass-list", controller.GetCiClassList)
}
