package database

import (
	"github.com/Galbeyte1/go-rest-api/internal/comment"
	"github.com/jinzhu/gorm"
)

// MigrateDB - migrates our database and creates our comment table
func MigrateDB(db *gorm.DB) error {
	// AutoMigrate takes comment struct defintion and defines all the columns
	if result := db.AutoMigrate(&comment.Comment{}); result.Error != nil {
		return result.Error
	}
	return nil
}