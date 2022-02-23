package controllers

import (
	"goapi/database"
	"goapi/models"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

var SecretKey = "secret"

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var existUser models.User

	database.DB.Where("email = ?", data["email"]).First(&existUser)

	if existUser.ID != 0 {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "User already exist",
		})
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	p := string(password)
	user := models.User{
		Email:    data["email"],
		Password: p,
		FullName: data["fullname"],
	}

	database.DB.Create(&user)

	return c.Status(200).JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "User not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"])); err != nil {

		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Invalid password",
		})
	}

	claim := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claim.SignedString([]byte(SecretKey))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Something went wrong",
		})
	}

	cookie := &fiber.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(cookie)

	return c.JSON(fiber.Map{
		"message": "Login success",
	})
}

func User(c *fiber.Ctx) error {
	var cookie = c.Cookies("token")

	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	claims := token.Claims.(*jwt.MapClaims)

	var user models.User

	database.DB.Where("id = ?", claims).First(&user)

	return c.JSON(user)
}
func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name: "token",
		Value: "",
		Expires: time.Now().Add(-1 * time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "Logout success",
	})
 
}