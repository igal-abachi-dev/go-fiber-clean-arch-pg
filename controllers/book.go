package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-clean-arch-pg/logic"
	"net/http"
)

//instead of jwt
//https://docs.gofiber.io/api/middleware/basicauth

//http://127.0.0.1:8080/api/book/1
func BookController(app fiber.Router, logic logic.BookLogic) {
	app.Get("/category/:id", GetBooksByCategory(logic))
	app.Get("/:id", GetBookById(logic))
}

func GetBooksByCategory(logic logic.BookLogic) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		valid, catId := isInt(id)
		if valid {
			fetched, err := logic.GetBooksByCategory(catId)
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
func GetBookById(logic logic.BookLogic) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		valid, bookId := isInt(id)
		if valid {
			//param ok
			fetched, err := logic.GetBookById(bookId)
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

//https://curatedgo.com/r/gojay-is-a-francoispqtgojay/index.html
/*
gojay.MarshalJSONObject()

func (m *DSTopic) MarshalJSONObject(enc *gojay.Encoder) {
	enc.AddIntKey("id", m.Id)
	enc.AddStringKey("slug", m.Slug)
}

func (m *DSTopics) MarshalJSONArray(enc *gojay.Encoder) {
	for _, e := range *m {
		enc.AddObject(e)
	}
}
func (m *DSTopics) IsNil() bool {
	return m == nil
}
*/
