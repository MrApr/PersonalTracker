package models

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
	"time"
)

//dbConnection determines and specifies database connection
type dbConnection struct {
	user   string
	pass   string
	host   string
	port   int
	dbName string
}

//DB holds database instance
var DB *gorm.DB

func ConnectToDb(maxConn int, maxOpen int) {
	conn := dbConnection{
		user:   "heroes",
		pass:   "22510070",
		host:   "localhost",
		port:   3306,
		dbName: "process_tracker",
	} //Todo make it to read from conf file
	//Todo check whether gorm creates connection pool or not
	DB = conn.mysqlConnection()

	dbConfigurator, err := DB.DB()
	if err != nil {
		panic(err)
	}

	conn.configDBConn(dbConfigurator, maxConn, maxOpen)
}

//makeMysqlConnection to mysql server
func (dbConn *dbConnection) mysqlConnection() *gorm.DB {
	dsn := dbConn.makeConnString()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

//makeConnString to connect to database
func (dbConn *dbConnection) makeConnString() string {
	port := strconv.Itoa(dbConn.port)
	dsn := dbConn.user + ":" + dbConn.pass + "@tcp(" + dbConn.host + ":" + port + ")/" + dbConn.dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	return dsn
}

//configure connected database
func (dbConn *dbConnection) configDBConn(db *sql.DB, maxIdleConn int, maxOpenConn int) {
	if maxIdleConn != 0 {
		db.SetMaxIdleConns(maxIdleConn)
	}

	if maxOpenConn != 0 {
		db.SetMaxOpenConns(maxOpenConn)
	}

	db.SetConnMaxLifetime(time.Minute * 1)
	db.SetConnMaxIdleTime(time.Hour)
}
