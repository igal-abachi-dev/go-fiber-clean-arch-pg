package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-clean-arch-pg/logic"
	"net/http"
)

// http://127.0.0.1:8080/api/category/
func CategoryController(app fiber.Router, logic logic.CategoryLogic) {
	app.Get("/", GetAllCategories(logic))
}

func GetAllCategories(logic logic.CategoryLogic) fiber.Handler {
	return func(c *fiber.Ctx) error {

		fetched, err := logic.GetAllCategories()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(nil)
		}
		return c.JSON(fetched)
	}
}

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
