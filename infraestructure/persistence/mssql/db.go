package persistence

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func NewRepositories(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		DbUser,
		DbPassword,
		DbHost,
		DbPort,
		DbName)

	db, err := gorm.Open(sqlserver.Open(dsn),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{TablePrefix: "blog_", SingularTable: true},
		})

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	// db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	// db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	// db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	// sqlDB, err := db.DB()
	// sqlDB.SetMaxIdleConns(10)
	// sqlDB.SetMaxOpenConns(100)
	return db, err
}
