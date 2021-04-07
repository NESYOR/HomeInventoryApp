package seed

import (
	"../models"
	"bufio"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)
var itemtypes []models.ItemTypes
var shapes []models.Shapes
func check(err error){
	if err!=nil{
		panic(err)
	}
}
func batchprocess() {
	f,err:=os.Open("./seed/itemtypes.txt")
	check(err)
	defer f.Close()
	scanner:=bufio.NewScanner(f)
	for scanner.Scan() {
		itemtypes = append(itemtypes, models.ItemTypes{Name: scanner.Text()})
	}
}
func SeedDB(db *gorm.DB){
	var user = models.Users{Email: "franklin12@gmail.com",Name: "Franklin Josh",Password: "Mine@3412",LastLogin: time.Now()}
	var _ = db.Create(&user)
	log.Println(user.ID)
	batchprocess()
	var _ = db.Create(itemtypes)

	var shapeNames = []string{"cube","cuboid","cylinder","torso","cone","prism","square pyramid"}

	for _,temp:=range shapeNames{
		shapes=append(shapes,models.Shapes{Name: temp})
	}
	db.Create(shapes)
}
