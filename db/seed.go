package db

import "github.com/yhashimoto1216/pon/api/internal/model"

func Seed() {
	DB.Create(&model.Sentence{
		UserID:  1,
		Content: "I like coffee.",
		Note:    "基本構文",
	})
}
