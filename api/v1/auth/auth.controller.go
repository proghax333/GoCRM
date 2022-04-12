package auth

import (
	"GoCRM/models"
	"GoCRM/utils"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

type LoginDTO struct {
	Login string `json:"login" xml:"login" form:"login" validate:"required"`
	Pass  string `json:"password" xml:"password" form:"password" validate:"required"`
}

type RegisterDTO struct {
	Username        string `json:"username" xml:"username" form:"username" validate:"required"`
	Email           string `json:"email" xml:"email" form:"email" validate:"required"`
	Password        string `json:"password" xml:"password" form:"password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" xml:"confirm_password" form:"confirm_password" validate:"required"`
}

func Login(c *fiber.Ctx) error {
	loginData := new(LoginDTO)

	if err := c.BodyParser(loginData); err != nil {
		c.Status(401).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid data sent",
		})

		return nil
	}

	log.Printf("Login: %v, Password: %v\n", loginData.Login, loginData.Pass)

	if validation := utils.ValidateStruct(loginData); len(validation) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(validation)
	}

	result := models.Users.FindOne(mgm.Ctx(), bson.M{
		"$or": bson.A{
			bson.M{"username": loginData.Login},
			bson.M{"email": loginData.Login},
		},
	})

	fmt.Printf("%v\n", result)

	user := models.User{}

	if err := result.Decode(&user); err != nil {
		fmt.Println(err)

		c.Status(401).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid login/password entered",
		})

		return nil
	}

	if loginData.Pass != user.Password {
		c.Status(401).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid login/password entered",
		})

		return nil
	}

	token, err := utils.SignToken(jwt.MapClaims{
		"id": user.ID,
	})

	if err != nil {
		c.Status(500).JSON(fiber.Map{
			"error":   true,
			"message": "Could not log in right now",
		})

		return nil
	}

	c.JSON(fiber.Map{
		"message": "Login was successful!",
		"data":    token,
	})

	return nil
}

func Register(c *fiber.Ctx) error {
	registerData := new(RegisterDTO)

	if err := c.BodyParser(registerData); err != nil {
		c.Status(401).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid data sent",
		})

		return nil
	}

	log.Printf("Login: %v, Password: %v\n", registerData.Username, registerData.Password)

	if validation := utils.ValidateStruct(registerData); len(validation) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(validation)
	}

	if registerData.Password != registerData.ConfirmPassword {
		return c.Status(401).JSON(fiber.Map{
			"error":   true,
			"message": "Confirm password does not match",
		})
	}

	existingUserResult := models.Users.FindOne(mgm.Ctx(), bson.M{
		"$or": bson.A{
			bson.M{"username": registerData.Username},
			bson.M{"email": registerData.Email},
		},
	})

	user := models.User{}
	err := existingUserResult.Decode(&user)

	if err == nil {
		fmt.Printf("existing user error: %v\n", err)

		return c.Status(400).JSON(fiber.Map{
			"error":   true,
			"message": "User already exists",
		})
	} else if err.Error() != "mongo: no documents in result" {
		return c.Status(500).JSON(fiber.Map{
			"error":   true,
			"message": "There was an error while registering. Try again later",
		})
	}

	fmt.Printf("User value: %v\n", user)

	if user.Username == registerData.Username || user.Email == registerData.Email {
		return c.Status(400).JSON(fiber.Map{
			"error":   true,
			"message": "User already exists",
		})
	}

	newUser := &models.User{
		Username: registerData.Username,
		Email:    registerData.Email,
		Password: registerData.Password,
		IsActive: true,
	}
	result := models.Users.Create(newUser)

	fmt.Printf("Result: %v\n", result)

	c.JSON(fiber.Map{
		"message": "Registration was successful!",
		"data":    newUser,
	})

	return nil
}
