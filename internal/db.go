package internal

import (
	"fmt"
	"log/slog"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		os.Getenv(DatabaseUser),
		os.Getenv(DatabasePassword),
		os.Getenv(DatabaseAddress),
		os.Getenv(DatabasePort),
		os.Getenv(DatabaseName),
		os.Getenv(DatabaseCharset),
		os.Getenv(DatabaseParseTime),
		os.Getenv(DatabaseLocale),
	)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		slog.Error("Couldn't connect to database")
		os.Exit(1)
	}
}

func MigrateModels() {

}
