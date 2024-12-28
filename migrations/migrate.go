package migrations

import (
	"log"

	"github.com/AungKyawPhyo1142/go-cli-app/config"
	"github.com/AungKyawPhyo1142/go-cli-app/models"
)

func Migrate() {

	if err := config.DB.AutoMigrate(&models.Note{}); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Print("Migration success")

}
