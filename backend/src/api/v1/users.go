package v1

import (
	"../../../db/infrastructure"
	"../../../db/models"
	"fmt"
	"github.com/gofiber/fiber"

)
type MyErrorHandler struct{
	message string `json:"message"`
	status int `json:"status"`
}
func GetallUsers(c *fiber.Ctx) error{
db:=infrastructure.GetConnection()
var users []models.Users
db.Find(&users)
	return c.JSON(users)
}

func GetUser(c *fiber.Ctx)error{
	db:=infrastructure.GetConnection()
	mail:=c.Params("email")
	var user models.Users
	result:=db.First(&user,"email= ?",mail)
	if result.Error != nil {
		fmt.Println(result.Error)
//errormsg:=MyErrorHandler{message: "No Such User",status: 500}
		//return c.JSON(errormsg)
		return c.SendStatus(500)
	}
	return c.JSON(user)
}

func PostUser(c *fiber.Ctx) error{
	//var user *models.Users
user:=new(models.Users)
	if err := c.BodyParser(user); err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
		return c.SendStatus(302)
	}
	fmt.Println(user)
return c.JSON(user)
}


