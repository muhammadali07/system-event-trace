package utils

import (
	"fmt"

	"github.com/muhammadali07/system-event-trace/services/journal/models"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DBInstance *gorm.DB
var err error

func ConnectDB(dbDriver, dbUser string, dbPassword string, dbHost string, dbPort int, db string) {
	var dialector gorm.Dialector
	var dsn string
	switch dbDriver {
	case "sqlite":
		dialector = sqlite.Open(dsn)
	case "mysql":
		dialector = mysql.Open(dsn)
	case "postgres":
		dsn = fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable TimeZone=Asia/Jakarta", dbHost, dbPort, dbUser, dbPassword, db)
		dialector = postgres.Open(dsn)
	case "sqlserver":
		dsn = fmt.Sprintf("sqlserver://%s:%s@%s:%v?database=%s", dbUser, dbPassword, dbHost, dbPort, db)
		dialector = sqlserver.Open(dsn)
	default:
		panic("Unsupported database driver connection: %s" + dbDriver)
	}

	DBInstance, err = gorm.Open(
		dialector,
		&gorm.Config{
			FullSaveAssociations: true,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		},
	)
	if err != nil {
		fmt.Println(err)
		panic("Database connection attempt was unsuccessful.....")
	}

	DBInstance.Logger.LogMode(logger.Info)
}

func MigrateDB() {
	DBInstance.AutoMigrate(&models.JournalData{})
	fmt.Println("Database migration completed....")
}
