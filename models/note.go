package models

import (
	"log"

	"github.com/AungKyawPhyo1142/go-cli-app/config"
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	Title string `gorm:"type:varchar(30)"`
	Body  string `gorm:"type:text"`
}

func GetNotes() ([]Note, error) {
	var notes []Note

	if err := config.DB.Select("ID", "Title", "Body").Find(&notes).Error; err != nil {
		log.Fatalf("Error getting notes from database: %v", err)
		return nil, err
	}

	return notes, nil

}

func (n *Note) SaveNotes(note Note) error {
	if err := config.DB.Create(&note).Error; err != nil {
		log.Fatalf("Error saving notes to database: %v", err)
		return err
	}

	return nil
}
