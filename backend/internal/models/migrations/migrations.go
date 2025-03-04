package migrations

import (
	"github.com/TheCodeGhinux/TaskHub/taskhub/internal/models"
)

// AllMigrationModels returns all the models that need to be migrated.
func AllMigrationModels() []interface{} {
	return []interface{}{
		&models.User{},
		&models.Tenant{},
		&models.Task{},
	}
}

// func MigrateDb() {
// 	err := db.DB.Postgres.AutoMigrate(
// 		&models.Tenant{},
// 		&models.User{},
// 		&models.Task{},
// 	)

// 	if err != nil {
// 		log.Fatal("Failed to migrate database:", err)
// 	}
// }
