package migrations

import (
	"fmt"
	"gorm.io/gorm"
)

// AlterColumn represents the struct for altering a column's type in a table.
type AlterColumn struct {
	Model     interface{}
	TableName string
	Column    string
	Type      string
}

// UpdateColumnType updates the type of a column in the database.
func (a *AlterColumn) UpdateColumnType(db *gorm.DB) error {
	// Use raw SQL to alter the column type in PostgreSQL (adjust SQL syntax for other DBs)
	query := fmt.Sprintf("ALTER TABLE %s ALTER COLUMN %s TYPE %s USING %s::%s", a.TableName, a.Column, a.Type, a.Column, a.Type)
	if err := db.Exec(query).Error; err != nil {
		return err
	}

	// Update the GORM model to reflect the changes
	if err := db.Migrator().AlterColumn(a.Model, a.Column); err != nil {
		return err
	}

	return nil
}
