package migrations

import (
	"github.com/TheCodeGhinux/TaskHub/taskhub/pkg/repository/db"
	"gorm.io/gorm"
)

// MigrateDb connects to the database and runs migrations.
func MigrateDb() {
	// Connect to the database using your DB connection logic
	// db := db.Connect()

	// Use the AutoMigrate function to automatically migrate all models
	if err := MigrateModels(db.DB.Postgres, AllMigrationModels(), nil); err != nil {
		// Handle the error appropriately (e.g., log it or return it)
		panic("Failed to migrate database: " + err.Error())
	}
}

// MigrateModels handles the automatic migration and column alterations.
func MigrateModels(db *gorm.DB, models []interface{}, alterColumns []AlterColumn) error {
	// Migrate the models using AutoMigrate
	if err := db.AutoMigrate(models...); err != nil {
		return err
	}

	// If there are columns to alter, apply the changes
	for _, alterColumn := range alterColumns {
		if err := alterColumn.UpdateColumnType(db); err != nil {
			return err
		}
	}

	return nil
}
