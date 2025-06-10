package db

func InitDB() {
	Connect()
	Migrate()
	// Seed() もし使うならここで
}
