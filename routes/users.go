package routes

import (
	"example.com/go-fiber-crud/db"
	"example.com/go-fiber-crud/models"
	"example.com/go-fiber-crud/utils"
	"github.com/gofiber/fiber/v2"
)

func signup(c *fiber.Ctx) error {
	// Set Content-Type header to application/json
	c.Set("Content-Type", "application/json")
	// Parse JSON body to User model
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}
	// Check existing user
	var existingUser models.User
	if err := models.FindUserByEmail(db.DB, user.Email, &existingUser); err == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "User already exists",
		})
	}
	// Create new user
	if err := user.Create(db.DB); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not create user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
	})
}

func login(c *fiber.Ctx) error {
	// Parse JSON body to get email and password
	loginData := new(struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	})
	if err := c.BodyParser(loginData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	// Find user by email
	var user models.User
	if err := models.FindUserByEmail(db.DB, loginData.Email, &user); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}
	// Verify password
	if !user.Validate(loginData.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}
	// Create token and return
	token, err := utils.GetToken(loginData.Email, user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not create token",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login successful",
		"token":   token,
	})
}
