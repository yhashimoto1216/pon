package db

import (
	"log"

	"github.com/yhashimoto1216/pon/api/internal/model"
)

func Migrate() {
	err := DB.AutoMigrate(
		&model.User{},
		&model.Sentence{},
		&model.Review{},
	)
	if err != nil {
		log.Fatal("Migration failed:", err)
	}
}
