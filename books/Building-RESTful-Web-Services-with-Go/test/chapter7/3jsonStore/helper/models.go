package helper

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "gituser"
	password = "dq4A@#19"
	dbname   = "mydb"
)

//type Shipment struct {
//	gorm.Model
//	Packages []Package // GORM 会自动将 Shipment.ID 的值赋给每个 Package.ShipmentID
//	Data     string    `sql:"type:JSONB NOT NULL DEFAULT '{}'::JSONB" json:"-"`
//}

type Package struct {
	gorm.Model
	Data string `sql:"JSONB NOT NULL DEFAULT '{}'::JSONB"`
	//ShipmentID uint   // 明确指明外键
}

// GORM creates tables with plural names.
// Use this to suppress it
//func (Shipment) TableName() string {
//	return "Shipment"
//}

func (Package) TableName() string {
	return "Package"
}

// InitDB 数据库初始化
func InitDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
		host,
		port,
		user,
		password,
		dbname)
	sqlDB, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了可以重新使用连接的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	db, err := gorm.Open(postgres.New(postgres.Config{
		PreferSimpleProtocol: true,
		Conn:                 sqlDB,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	/*
		// The below AutoMigrate is equivalent to this
		if !db.HasTable("Shipment") {
			db.CreateTable(&Shipment{})
		}

		if !db.HasTable("Package") {
			db.CreateTable(&Package{})
		}
	*/

	//err = db.AutoMigrate(&Shipment{}, &Package{})
	err = db.AutoMigrate(&Package{})
	if err != nil {
		return nil, err
	}
	return db, err
}
