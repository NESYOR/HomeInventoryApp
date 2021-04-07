package migrations

import (
	"../models"
	"gorm.io/gorm"
)

func RevokeDB(db *gorm.DB)  {
	db.Migrator().DropTable(&models.Shapes{})
	db.Migrator().DropTable(&models.Shapes{})
	db.Migrator().DropTable(&models.Shapes{})
	db.Migrator().DropTable(&models.Shapes{})
	db.Migrator().DropTable(&models.Shapes{})
	db.Migrator().DropTable(&models.Shapes{})
	db.Migrator().DropTable(&models.Shapes{})
	db.Migrator().DropTable(&models.Shapes{})
	db.Migrator().DropTable(&models.Shapes{})
}