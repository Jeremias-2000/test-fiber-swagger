package routes

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/test-fiber/database"
	"github.com/test-fiber/models"
)

type User struct {
	// this is not the model User, see this as the serializer
	ID uint `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

func CreateResponseUser(userModel models.User) User{
	return User{ID:userModel.ID,FirstName: userModel.FirstName,LastName: userModel.LastName}
}

// CreateUser godoc
// @Summary create user
// @Description Register a new user
// @Tags user
// @Accept */*
// @Produce json
// @Success 201 {object} .User{}
// @Router /api/user/signin [post]
func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err:= c.BodyParser(&user);err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&user)
	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(responseUser) 
}

// GetUsers godoc
// @Summary Show users
// @Description Show the users registered
// @Tags user
// @Accept */*
// @Produce json
// @Success 200 {object} []User{}
// @Router /api/users [get]
func GetUsers(c *fiber.Ctx) error {
	users := []models.User{}


	database.Database.Db.Find(&users)

	responseUsers:= []User{}
	for _,user := range users {
		responseUser := CreateResponseUser(user)
		responseUsers = append(responseUsers, responseUser)
	}

	return c.Status(200).JSON(responseUsers)
}

func findUser(id int,user *models.User,c *fiber.Ctx) error {
	database.Database.Db.Find(&user,"id = ?",id)
	if user.ID == 0 {
		return errors.New("user does not exist")
	}
	return c.Status(404).JSON("Error when search user")	
}

// GetUser godoc
// @Summary fetch user
// @Description fetch user registered
// @Tags user
// @Accept */*
// @Produce json
// @Success 200 {object} User{}
// @Router /api/user/:id [get]
func GetUser(c *fiber.Ctx) error {
	id, err:= c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	findUser(id,&user,c)

	responseUser:= CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)
}

// UpdateUser godoc
// @Summary update user
// @Description Update user data
// @Tags user
// @Accept */*
// @Produce json
// @Success 200 {object} User{}
// @Router /api/user/update/:id [put]
func UpdateUser(c *fiber.Ctx)  error{
	id, err:= c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

    findUser(id,&user,c)

	type UpdateUser struct {
		FirstName string `json:"first_name"`
		LastName string `json:"last_name"`
	}

	var updateData UpdateUser

	if err:= c.BodyParser(&updateData);err !=nil {
		return c.Status(500).JSON(err.Error())
	}

	user.FirstName = updateData.FirstName
	user.LastName = updateData.LastName

	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)
}

// DeleteUser godoc
// @Summary delete user
// @Description delete user registered
// @Tags user
// @Accept */*
// @Produce json
// @Success 200 {object} string
// @Router /api/user/unsubscribe/:id [delete]
func DeleteUser(c *fiber.Ctx) error {
	id, err:= c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	findUser(id,&user,c)

    database.Database.Db.Delete(&user)

	return c.Status(200).SendString("Successfully deleted user")
}

