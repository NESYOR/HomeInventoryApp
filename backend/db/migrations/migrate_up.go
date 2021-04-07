package migrations

import (
	"log"

	"../models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func addConstraintsToDB(db *gorm.DB) {
	db.Migrator().CreateConstraint(&models.Countries{}, "states")
	db.Migrator().CreateConstraint(&models.Countries{}, "fk_users_profiles")
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
	log.Println("Strated Migrating Database")
	db.AutoMigrate(models.Users{}, models.Item_Types{}, models.Shapes{}, models.Inventory_Locations{}, models.Countries{}, models.States{}, models.Addresses{}, models.Companies{}, models.Sizes{}, models.Items{}, models.Item_Infos{}, models.Item_Images{}, models.Related_Items{})
	addConstraintsToDB(db)
}
