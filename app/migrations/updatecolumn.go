package migrations

import (
	"fmt"
	"gorm.io/gorm"
)

type AlterColumn struct {
	Model     interface{}
	TableName string
	Column    string
	Type      string
}

func (a *AlterColumn) UpdateColumnType(db *gorm.DB) error {
	query := fmt.Sprintf("ALTER TABLE %s MODIFY COLUMN %s %s", a.TableName, a.Column, a.Type)
	if err := db.Exec(query).Error; err != nil {
		return err
	}

	if err := db.Migrator().AlterColumn(a.Model, a.Column); err != nil {
		return err
	}

	return nil
}
