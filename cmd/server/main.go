package main

import (
	"github.com/yhashimoto1216/pon/api/db"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	db.InitDB() // ← 接続 + マイグレーション

	app.Listen(":8080")
}
