package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-clean-arch-pg/logic"
	"net/http"
)

//instead of jwt
//https://docs.gofiber.io/api/middleware/basicauth

// http://127.0.0.1:8080/api/fileContent/pdf/1
func FileContentController(app fiber.Router, logic logic.FileContentLogic) {
	app.Get("/text/:id", GetBookTextContent(logic))
	app.Get("/pdf/:id", GetBookUrl(logic))
}

func GetBookTextContent(logic logic.FileContentLogic) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		valid, catId := isInt(id)
		if valid {
			fetched, err := logic.GetBookTextContent(catId)
			if err != nil {
				c.Status(http.StatusInternalServerError)
				return c.JSON(nil)
			}
			return c.JSON(fetched)
		}

		c.Status(http.StatusNotFound)
		return c.JSON(nil)
	}
}
func GetBookUrl(logic logic.FileContentLogic) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		valid, bookId := isInt(id)
		if valid {
			//param ok
			fetched, err := logic.GetBookUrl(bookId)
			if err != nil {
				c.Status(http.StatusInternalServerError)
				return c.JSON(nil)
			}
			return c.JSON(fetched)
		}
		c.Status(http.StatusNotFound)
		return c.JSON(nil)
	}
}
