package migrations

import (
	"../models"
	"../seed"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)
var DataBase *gorm.DB
func addConstraintsToDB(db *gorm.DB) {
	db.Migrator().CreateConstraint(&models.Countries{}, "states")
	db.Migrator().CreateConstraint(&models.Countries{}, "fk_users_profiles")
}
func StartSeeding(db *gorm.DB){
	seed.SeedDB(db)
}
// db.Migrator().CreateConstraint(&Users{}), "Profiles")
// db.Migrator().CreateConstraint(&Users{}), "fk_users_profiles")
// which translates to the following sql code for postgres:

// ALTER TABLE `Profiles` ADD CONSTRAINT `fk_users_profiles` FORIEGN KEY (`useres_id`) REFRENCES `users`(`id`))
func Initdb() {
	dsn := "host=172.18.0.3 user=admin password=Admin dbname=inventory_app port=5432 sslmode=disable "
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicln("Database connection Error")
	}
	DataBase=db
	log.Println("Strated Migrating Database")
	db.AutoMigrate(models.Users{}, models.ItemTypes{}, models.Shapes{}, models.InventoryLocations{}, models.Countries{}, models.States{}, models.Addresses{}, models.Companies{}, models.Sizes{}, models.Items{}, models.ItemInfos{}, models.ItemImages{}, models.RelatedItems{})
	//addConstraintsToDB(db)

}
