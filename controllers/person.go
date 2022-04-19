package controllers

import (
	"fiber-rest-api/logic"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func PersonController(app fiber.Router, logic logic.PersonLogic) {
	app.Get("/" /*,middleware.JWTProtected()*/, GetPersons(logic))
	app.Get("/:id" /*,middleware.JWTProtected()*/, GetPerson(logic))
	//
	//app.Post("/person", AddBook(service))
	//app.Put("/person", UpdateBook(service))
	//app.Delete("/person", RemoveBook(service))
}

// GetBooks is handler/controller which lists all Books from the BookShop
func GetPersons(logic logic.PersonLogic) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fetched, err := logic.GetAll()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(nil)
		}
		return c.JSON(fetched)
	}
}
func GetPerson(logic logic.PersonLogic) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		fetched, err := logic.GetItem(id)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(nil)
		}
		return c.JSON(fetched)
	}
}

/*
// AddBook is handler/controller which creates Books in the BookShop
func AddBook(service book.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Book
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.BookErrorResponse(err))
		}
		if requestBody.Author == "" || requestBody.Title == "" {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.BookErrorResponse(errors.New(
				"Please specify title and author")))
		}
		result, err := service.InsertBook(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.BookErrorResponse(err))
		}
		return c.JSON(presenter.BookSuccessResponse(result))
	}
}

// UpdateBook is handler/controller which updates data of Books in the BookShop
func UpdateBook(service book.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Book
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.BookErrorResponse(err))
		}
		result, err := service.UpdateBook(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.BookErrorResponse(err))
		}
		return c.JSON(presenter.BookSuccessResponse(result))
	}
}

// RemoveBook is handler/controller which removes Books from the BookShop
func RemoveBook(service book.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.DeleteRequest
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.BookErrorResponse(err))
		}
		bookID := requestBody.ID
		err = service.RemoveBook(bookID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.BookErrorResponse(err))
		}
		return c.JSON(&fiber.Map{
			"status": true,
			"data":   "updated successfully",
			"err":    nil,
		})
	}
}
*/
/*
package utils

import (
"github.com/go-playground/validator/v10"
"github.com/google/uuid"
)

// NewValidator func for create a new validator for model fields.
func NewValidator() *validator.Validate {
	// Create a new validator for a Book model.
	validate := validator.New()

	// Custom validation for uuid.UUID fields.
	_ = validate.RegisterValidation("uuid", func(fl validator.FieldLevel) bool {
		field := fl.Field().String()
		if _, err := uuid.Parse(field); err != nil {
			return true
		}
		return false
	})

	return validate
}

// ValidatorErrors func for show validation errors for each invalid fields.
func ValidatorErrors(err error) map[string]string {
	// Define fields map.
	fields := map[string]string{}

	// Make error message for each invalid field.
	for _, err := range err.(validator.ValidationErrors) {
		fields[err.Field()] = err.Error()
	}

	return fields

}
*/
