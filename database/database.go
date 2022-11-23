package database

import (
	"fmt"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPGLocal() *gorm.DB {
	var (
		dbUser = "postgresuser"
		dbPass = "postgrespassword"
		dbHost = "postgres"
		dbName = "postgres"
		dbPort = "5432"
		// sslMode = "disable"
	)

	// dsn
	dsn := fmt.Sprintf(`postgres://%v:%v@%v:%v/%v?sslmode=disable`,
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbName,
	)
	// dsn := fmt.Sprintf(`
	// 	host=%s user=%s password=%s dbname=%s port=%s sslmode=%s`,
	// 	dbHost,
	// 	dbUser,
	// 	dbPass,
	// 	dbName,
	// 	dbPort,
	// 	sslMode,
	// )

	log.Print("dsn:", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Warn("Connected to database local Failed:", err)
	}
	log.Warn("Connected to database local")

	return db
}
