package bootstrap

import (
	"fmt"
	"github.com/uzzalhcse/amadeus-go/app/exceptions"
	"github.com/uzzalhcse/amadeus-go/config"
	"github.com/uzzalhcse/amadeus-go/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func createConnectionPool(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func NewDatabase(dbConfig config.DatabaseConfig) (DB *gorm.DB) {
	connString := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name,
		dbConfig.Username,
		dbConfig.Password,
	)

	var err error
	DB, err = createConnectionPool(connString)
	if err != nil {
		exceptions.PanicIfNeeded(fmt.Errorf("[INIT] failed to connect to the database: %v", err))
	}

	fmt.Println("[INIT] Database connection established")

	err = database.Migrate(DB)
	if err != nil {
		fmt.Println("DB Migration Error: ", err.Error())
	}

	return DB
}
func CloseDBConnection(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		exceptions.PanicIfNeeded(err)
		return
	}

	err = sqlDB.Close()
	if err != nil {
		exceptions.PanicIfNeeded(err)
	}
}
