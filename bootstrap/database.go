package bootstrap

import (
	"github.com/nevzatcirak/go-gin-poc/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func NewSQLiteDatabase(env *Env) *gorm.DB {
	//dbHost := env.DBHost
	//dbPort := env.DBPort
	//dbUser := env.DBUser
	//dbPass := env.DBPass

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect videoRepository")
	}
	db.AutoMigrate(&domain.Video{}, &domain.User{})

	return db
}

func CloseSQLiteDBConnection(db *gorm.DB) {
	if db == nil {
		return
	}

	sqlDB, _ := db.DB()
	err := sqlDB.Close()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection to SQLite closed.")
}
